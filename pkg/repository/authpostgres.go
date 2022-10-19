package repository

import (
	"database/sql"
	"fmt"

	gofermart "github.com/AltynayK/go-musthave-diploma-tpl"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {

	return &AuthPostgres{db: db}

}

func (r *AuthPostgres) CreateUser(user gofermart.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password) values ($1, $2) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
