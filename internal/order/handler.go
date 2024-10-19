// order/order.go
package order

import (
	"context"
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/thongsoi/testc/database"
)

var orderTemplate *template.Template

// Initialize templates once at server startup
func init() {
	var err error
	orderTemplate, err = template.ParseFiles("templates/order.html")
	if err != nil {
		log.Fatal("Unable to parse template:", err)
	}
}

// FormHandler renders the order form with markets data
func FormHandler(w http.ResponseWriter, r *http.Request) {
	markets, err := FetchMarkets(database.GetDB())
	if err != nil {
		http.Error(w, "Unable to fetch markets", http.StatusInternalServerError)
		return
	}

	data := struct {
		Markets []Market
	}{
		Markets: markets,
	}

	err = orderTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

// GetMarketsHandler returns a list of markets in JSON format
func GetMarketsHandler(w http.ResponseWriter, r *http.Request) {
	markets, err := FetchMarkets(database.GetDB())
	if err != nil {
		http.Error(w, "Unable to fetch markets", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, markets)
}

// GetProductsHandler returns products based on market ID in JSON format
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	marketID, err := getQueryParamInt(r, "market_id")
	if err != nil {
		http.Error(w, "Invalid market ID", http.StatusBadRequest)
		return
	}

	products, err := FetchProductsByMarket(database.GetDB(), marketID)
	if err != nil {
		http.Error(w, "Unable to fetch products", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, products)
}

// SubmitOrderHandler processes the order form submission
func SubmitOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	marketID, err := getFormValueInt(r, "market_id")
	if err != nil {
		http.Error(w, "Invalid market ID", http.StatusBadRequest)
		return
	}

	productID, err := getFormValueInt(r, "product_id")
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	quantity, err := getFormValueInt(r, "quantity")
	if err != nil {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	// Process order: Save the order in the database (to be implemented)
	if err := processOrder(marketID, productID, quantity); err != nil {
		http.Error(w, "Unable to process order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<div>Order submitted successfully!</div>"))
}

// FetchMarkets retrieves all markets from the database
func FetchMarkets(db *sql.DB) ([]Market, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "SELECT id, en_name FROM markets"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var markets []Market
	for rows.Next() {
		var m Market
		if err := rows.Scan(&m.ID, &m.EnName); err != nil {
			return nil, err
		}
		markets = append(markets, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return markets, nil
}

// FetchProductsByMarket retrieves products by market ID
func FetchProductsByMarket(db *sql.DB, marketID int) ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "SELECT id, name, price FROM products WHERE market_id = $1"
	rows, err := db.QueryContext(ctx, query, marketID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// Helper functions

// processOrder saves the order to the database (to be implemented)
func processOrder(marketID, productID, quantity int) error {
	db := database.GetDB()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO orders (market_id, product_id, quantity) VALUES ($1, $2, $3)"
	_, err := db.ExecContext(ctx, query, marketID, productID, quantity)
	return err
}

// respondWithJSON is a utility to send a JSON response
func respondWithJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// getFormValueInt retrieves and converts a form value to an int
func getFormValueInt(r *http.Request, key string) (int, error) {
	value := r.FormValue(key)
	return strconv.Atoi(value)
}

// getQueryParamInt retrieves and converts a query param to an int
func getQueryParamInt(r *http.Request, key string) (int, error) {
	value := r.URL.Query().Get(key)
	return strconv.Atoi(value)
}
