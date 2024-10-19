package order

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

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
	/*
		marketID := r.FormValue("market_id")
		productID := r.FormValue("product_id")
		quantity := r.FormValue("quantity")
	*/
	// Here you would typically process the order
	// For now, we'll just return a confirmation message

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<div>Order submitted successfully!</div>"))
}

func FetchMarkets(db *sql.DB) ([]Market, error) {
	// Implementation depends on your database schema
	// This is a placeholder
	return []Market{
		{ID: 1, EnName: "Market A"},
		{ID: 2, EnName: "Market B"},
	}, nil
}

func FetchProductsByMarket(db *sql.DB, marketID int) ([]Product, error) {
	// Implementation depends on your database schema
	// This is a placeholder
	return []Product{
		{ID: 1, Name: "Product 1", Price: 10.99},
		{ID: 2, Name: "Product 2", Price: 15.99},
	}, nil
}
