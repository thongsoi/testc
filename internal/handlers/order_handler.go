package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/thongsoi/testc/internal/models"
	"github.com/thongsoi/testc/internal/services"
)

type OrderHandler struct {
	Service *services.OrderService
}

func (h *OrderHandler) GetMarkets(w http.ResponseWriter, r *http.Request) {
	markets, err := h.Service.GetMarkets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(markets)
}

func (h *OrderHandler) GetSubmarkets(w http.ResponseWriter, r *http.Request) {
	marketIDStr := r.URL.Query().Get("marketID")
	marketID, _ := strconv.Atoi(marketIDStr)

	submarkets, err := h.Service.GetSubmarkets(marketID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(submarkets)
}

func (h *OrderHandler) SubmitOrder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	marketID, _ := strconv.Atoi(r.FormValue("marketID"))
	submarketID, _ := strconv.Atoi(r.FormValue("submarketID"))

	order := models.Order{
		MarketID:    marketID,
		SubmarketID: submarketID,
	}

	if err := h.Service.CreateOrder(order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/orders", http.StatusSeeOther)
}
