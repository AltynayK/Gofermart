package repository

import (
	gofermart "github.com/AltynayK/go-musthave-diploma-tpl"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user gofermart.User) (int, error)
	GetUser(login, password string) (gofermart.User, error)
}

type Order interface {
	Create(userId int, number string) (int, error)
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
