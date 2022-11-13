package service

import (
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/models"
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/repository"
)

type OrderService struct {
	repo repository.Repository
}

func NewOrderService(repo repository.Repository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(userID int, number string) error {
	return s.repo.CreateOrder(userID, number)
}

func (s *OrderService) GetOrderByUserAndNumber(userID int, number int) ([]models.OrdersOut, error) {
	return s.repo.GetOrderByUserAndNumber(userID, number)
}

func (s *OrderService) GetOrder(number int) ([]models.OrdersOut, error) {
	return s.repo.GetOrder(number)
}
func (s *OrderService) GetAllOrders(userID int) ([]models.OrdersOut, error) {
	return s.repo.GetAllOrders(userID)
}

func (s *OrderService) GetOrderUserID(number string) (int, error) {
	return s.repo.GetOrderUserID(number)

}
func (s *OrderService) PostWithdrawBalance(order models.Withdrawals) (int64, error) {
	return s.repo.PostWithdrawBalance(order)
}

func (s *OrderService) PostNewWithdrawBalance(order models.Withdrawals, userID int) error {
	return s.repo.PostNewWithdrawBalance(order, userID)
}

func (s *OrderService) GetUserCurrent(userID int) (float32, error) {
	return s.repo.GetUserCurrent(userID)
}

func (s *OrderService) GetUserWithdrawn(userID int) (float32, error) {
	return s.repo.GetUserWithdrawn(userID)
}

func (s *OrderService) UpdateUserBalance(userID int, current float32) (int64, error) {
	return s.repo.UpdateUserBalance(userID, current)
}

func (s *OrderService) GetAllWithdrawals(userID int) ([]models.Withdrawals, error) {
	return s.repo.GetAllWithdrawals(userID)
}

func (s *OrderService) PostBalance(order models.OrderBalance) (int64, error) {
	return s.repo.PostBalance(order)
}
