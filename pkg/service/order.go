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

func (s *OrderService) Create(userID int, number string) error {
	return s.repo.Create(userID, number)
}

func (s *OrderService) GetOrderByUserAndNumber(userID int, number int) ([]gofermart.OrdersOut, error) {
	return s.repo.GetOrderByUserAndNumber(userID, number)
}

func (s *OrderService) GetOrder(number int) ([]gofermart.OrdersOut, error) {
	return s.repo.GetOrder(number)
}
func (s *OrderService) GetAll(userID int) ([]gofermart.OrdersOut, error) {
	return s.repo.GetAll(userID)
}

// func (s *OrderService) GetUserBalance(userID int) ([]gofermart.UserBalance, error) {
// 	return s.repo.GetUserBalance(userID)
// }

func (s *OrderService) PostWithdrawBalance(order gofermart.Withdrawals) (int64, error) {
	return s.repo.PostWithdrawBalance(order)
}

func (s *OrderService) GetUserCurrent(userID int) (int, error) {
	return s.repo.GetUserCurrent(userID)
}

func (s *OrderService) GetUserWithdrawn(userID int) (int, error) {
	return s.repo.GetUserWithdrawn(userID)
}

func (s *OrderService) UpdateUserBalance(userID int, current int) (int64, error) {
	return s.repo.UpdateUserBalance(userID, current)
}

func (s *OrderService) GetAllWithdrawals(userID int) ([]gofermart.Withdrawals, error) {
	return s.repo.GetAllWithdrawals(userID)
}
