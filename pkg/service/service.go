package service

import (
	gofermart "github.com/AltynayK/go-musthave-diploma-tpl"
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/repository"
)

type Authorization interface {
	CreateUser(user gofermart.User) error
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Order interface {
	Create(userID int, number string) error
	GetAll(userID int) ([]gofermart.OrdersOut, error)
	GetOrderByUserAndNumber(userID int, number int) ([]gofermart.OrdersOut, error)
	GetOrder(number int) ([]gofermart.OrdersOut, error)
	GetUserBalance(userID int) ([]gofermart.UserBalance, error)
}

type Service struct {
	Authorization
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Order:         NewOrderService(repos.Order),
	}
}
