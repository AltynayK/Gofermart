package repository

import (
	gofermart "github.com/AltynayK/go-musthave-diploma-tpl"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user gofermart.User) error
	GetUser(login, password string) (gofermart.User, error)
}

type Order interface {
	Create(userID int, number string) error
	GetAll(userID int) ([]gofermart.OrdersOut, error)
	GetOrderByUserAndNumber(userID int, number int) ([]gofermart.OrdersOut, error)
	GetOrder(number int) ([]gofermart.OrdersOut, error)
	GetUserBalance(userID int) ([]gofermart.UserBalance, error)
}

type Repository struct {
	Authorization
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Order:         NewOrderPostgres(db),
	}
}
