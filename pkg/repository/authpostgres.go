package repository

import (
	"fmt"

	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {

	return &AuthPostgres{db: db}

}

func (r *AuthPostgres) CreateUser(user models.User) error {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password) values ($1, $2) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) GetUser(login, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, login, password)
	return user, err
}
