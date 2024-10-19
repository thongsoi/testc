package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/thongsoi/testc/internal/handlers"
	"github.com/thongsoi/testc/internal/repositories"
	"github.com/thongsoi/testc/internal/services"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://dev1:dev1pg@localhost:5432/biomassx?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//r := mux.NewRouter()
	// Static files
	//http.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	orderRepo := &repositories.OrderRepository{DB: db}
	orderService := &services.OrderService{Repo: orderRepo}
	orderHandler := &handlers.OrderHandler{Service: orderService}

	http.HandleFunc("/api/markets", orderHandler.GetMarkets)
	http.HandleFunc("/api/submarkets", orderHandler.GetSubmarkets)
	http.HandleFunc("/submit-order", orderHandler.SubmitOrder)

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
