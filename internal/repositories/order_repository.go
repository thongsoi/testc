package repositories

import (
	"database/sql"

	"github.com/thongsoi/testc/internal/models"
)

type OrderRepository struct {
	DB *sql.DB
}

func (r *OrderRepository) GetMarkets() ([]models.Market, error) {
	query := `SELECT market_id, market_name FROM markets`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var markets []models.Market
	for rows.Next() {
		var market models.Market
		if err := rows.Scan(&market.MarketID, &market.MarketName); err != nil {
			return nil, err
		}
		markets = append(markets, market)
	}
	return markets, nil
}

func (r *OrderRepository) GetSubmarkets(marketID int) ([]models.Submarket, error) {
	query := `SELECT submarket_id, submarket_name, market_id FROM submarkets WHERE market_id=$1`
	rows, err := r.DB.Query(query, marketID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var submarkets []models.Submarket
	for rows.Next() {
		var submarket models.Submarket
		if err := rows.Scan(&submarket.SubmarketID, &submarket.SubmarketName, &submarket.MarketID); err != nil {
			return nil, err
		}
		submarkets = append(submarkets, submarket)
	}
	return submarkets, nil
}

func (r *OrderRepository) CreateOrder(order models.Order) error {
	query := `INSERT INTO orders (market_id, submarket_id) VALUES ($1, $2)`
	_, err := r.DB.Exec(query, order.MarketID, order.SubmarketID)
	return err
}
