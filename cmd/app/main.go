package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // Assuming you're using PostgreSQL, adjust if using a different database
	"github.com/thongsoi/testc/database"
	"github.com/thongsoi/testc/internal/order"
)

func main() {
	// Initialize the database connection
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

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
