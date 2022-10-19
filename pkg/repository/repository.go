package repository

import (
	"database/sql"

	gofermart "github.com/AltynayK/go-musthave-diploma-tpl"
)

type Authorization interface {
	CreateUser(user gofermart.User) (int, error)
}

type Order interface {
}

type Repository struct {
	Authorization
	Order
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
