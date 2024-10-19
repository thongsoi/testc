// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/thongsoi/testc/database"
	"github.com/thongsoi/testc/internal/order"
)

func main() {
	// Initialize the database
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer func() {
		if err := database.CloseDB(); err != nil {
			log.Println("Failed to close the database:", err)
		}
	}()

	// Register handlers for your application
	http.HandleFunc("/order-form", order.FormHandler)
	http.HandleFunc("/get-markets", order.GetMarketsHandler)
	http.HandleFunc("/get-products", order.GetProductsHandler)
	http.HandleFunc("/submit-order", order.SubmitOrderHandler)

	// Start the server in a separate goroutine to support graceful shutdown
	go func() {
		fmt.Println("Server starting on port 9000...")
		if err := http.ListenAndServe(":9000", nil); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %s", err)
		}
	}()

	// Create a channel to listen for interrupt or termination signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	<-quit

	// Perform graceful shutdown
	fmt.Println("\nShutting down server...")
	time.Sleep(2 * time.Second) // Allow some time for in-flight requests to complete
	fmt.Println("Server stopped.")
}
