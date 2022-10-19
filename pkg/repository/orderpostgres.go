package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) Create(userId int, number string) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createOrder := fmt.Sprintf("INSERT INTO %s (number, user_id) VALUES ($1, $2) RETURNING id", ordersTable)
	row := tx.QueryRow(createOrder, number, userId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
