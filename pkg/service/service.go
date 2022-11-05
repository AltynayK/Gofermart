package service

import (
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/models"
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/repository"
)

type Auth interface {
	CreateUser(user models.User) error
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Orders interface {
	Create(userID int, number string) error
	GetAll(userID int) ([]models.OrdersOut, error)
	GetOrderByUserAndNumber(userID int, number int) ([]models.OrdersOut, error)
	GetOrder(number int) ([]models.OrdersOut, error)
	PostWithdrawBalance(order models.Withdrawals) (int64, error)
	PostNewWithdrawBalance(order models.Withdrawals, userID int) error
	GetUserCurrent(userID int) (float32, error)
	GetUserWithdrawn(userID int) (float32, error)
	UpdateUserBalance(userID int, current float32) (int64, error)
	GetAllWithdrawals(userID int) ([]models.Withdrawals, error)
	PostBalance(order models.OrderBalance) (int64, error)
	GetOrderUserID(number string) (int, error)
}

type Service struct {
	Auth
	Orders
}

func NewService(repos *repository.MyStruct) *Service {
	return &Service{
		Auth:   NewAuthService(repos.Repository),
		Orders: NewOrderService(repos.Repository),
	}
}
