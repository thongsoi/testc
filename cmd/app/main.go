package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // Assuming you're using PostgreSQL, adjust if using a different database
	"github.com/thongsoi/testc/database"
	"github.com/thongsoi/testc/internal/order"
)

func main() {
	// Initialize database connection
	db, err := initDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	// Set up routes
	http.HandleFunc("/", order.FormHandler)
	http.HandleFunc("/get-markets", order.GetMarketsHandler)
	http.HandleFunc("/get-products", order.GetProductsHandler)
	http.HandleFunc("/submit-order", order.SubmitOrderHandler)

	// Serve static files (if any)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initDB() (*sql.DB, error) {
	// Replace these with your actual database credentials
	connStr := "user=your_username dbname=your_dbname password=your_password host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Set the database connection in your package
	database.GetDB()

	return db, nil
}
