// order/order.go
package order

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/thongsoi/testc/database"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	markets, err := FetchMarkets(database.GetDB())
	if err != nil {
		http.Error(w, "Unable to fetch markets", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/order.html")
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Markets []Market
	}{
		Markets: markets,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func GetMarketsHandler(w http.ResponseWriter, r *http.Request) {
	markets, err := FetchMarkets(database.GetDB())
	if err != nil {
		http.Error(w, "Unable to fetch markets", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(markets)
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	marketID := r.URL.Query().Get("market_id")
	marketIDInt, err := strconv.Atoi(marketID)
	if err != nil {
		http.Error(w, "Invalid market ID", http.StatusBadRequest)
		return
	}

	products, err := FetchProductsByMarket(database.GetDB(), marketIDInt)
	if err != nil {
		http.Error(w, "Unable to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func SubmitOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Validate form inputs
	marketID := r.FormValue("market_id")
	productID := r.FormValue("product_id")
	quantity := r.FormValue("quantity")

	if marketID == "" || productID == "" || quantity == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Convert form values to appropriate types
	marketIDInt, err := strconv.Atoi(marketID)
	if err != nil {
		http.Error(w, "Invalid market ID", http.StatusBadRequest)
		return
	}

	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	quantityInt, err := strconv.Atoi(quantity)
	if err != nil {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	// Process the order
	// For now, we'll just return a confirmation message
	fmt.Println(marketIDInt)
	fmt.Println(productIDInt)
	fmt.Println(quantityInt)

	//test uses of 3 variables

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<div>Order submitted successfully!</div>"))
}

func FetchMarkets(db *sql.DB) ([]Market, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT id, en_name FROM markets")
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

func FetchProductsByMarket(db *sql.DB, marketID int) ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT id, name, price FROM products WHERE market_id = $1", marketID)
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
