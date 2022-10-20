package service

import (
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(userID int, number string) error {
	return s.repo.Create(userID, number)
}
