package services

import (
	"github.com/thongsoi/testc/internal/models"
	"github.com/thongsoi/testc/internal/repositories"
)

type OrderService struct {
	Repo *repositories.OrderRepository
}

func (s *OrderService) GetMarkets() ([]models.Market, error) {
	return s.Repo.GetMarkets()
}

func (s *OrderService) GetSubmarkets(marketID int) ([]models.Submarket, error) {
	return s.Repo.GetSubmarkets(marketID)
}

func (s *OrderService) CreateOrder(order models.Order) error {
	return s.Repo.CreateOrder(order)
}
