package models

type Market struct {
	MarketID   int    `json:"market_id"`
	MarketName string `json:"market_name"`
}

type Submarket struct {
	SubmarketID   int    `json:"submarket_id"`
	SubmarketName string `json:"submarket_name"`
	MarketID      int    `json:"market_id"`
}

type Order struct {
	OrderID     int `json:"order_id"`
	MarketID    int `json:"market_id"`
	SubmarketID int `json:"submarket_id"`
}
