package service

import (
	gofermart "github.com/AltynayK/go-musthave-diploma-tpl"
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(userID int, number string) (int, error) {
	return s.repo.Create(userID, number)
}

func (s *OrderService) GetAll(userID int) ([]gofermart.OrdersOut, error) {
	return s.repo.GetAll(userID)
}
