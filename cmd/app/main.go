// main.go
package main

import (
	"log"
	"net/http"

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

	// Set up HTTP handlers
	http.HandleFunc("/form", order.FormHandler)
	http.HandleFunc("/markets", order.GetMarketsHandler)
	http.HandleFunc("/products", order.GetProductsHandler)
	http.HandleFunc("/submit", order.SubmitOrderHandler)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
