package service

import (
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/models"
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) error
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Order interface {
	Create(userID int, number string) error
	GetAll(userID int) ([]models.OrdersOut, error)
	GetOrderByUserAndNumber(userID int, number int) ([]models.OrdersOut, error)
	GetOrder(number int) ([]models.OrdersOut, error)
	//GetUserBalance(userID int) ([]gofermart.UserBalance, error)
	PostWithdrawBalance(order models.Withdrawals) (int64, error)
	GetUserCurrent(userID int) (float32, error)
	GetUserWithdrawn(userID int) (float32, error)
	UpdateUserBalance(userID int, current float32) (int64, error)
	GetAllWithdrawals(userID int) ([]models.Withdrawals, error)
	PostBalance(order models.OrderBalance) (int64, error)
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
