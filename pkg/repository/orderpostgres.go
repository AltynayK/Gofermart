package repository

import (
	"fmt"
	"time"

	gofermart "github.com/AltynayK/go-musthave-diploma-tpl"
	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) Create(userID int, number string) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createOrder := fmt.Sprintf("INSERT INTO %s (number, user_id, status, uploaded_at) VALUES ($1, $2, $3, $4) RETURNING id", ordersTable)
	row := tx.QueryRow(createOrder, number, userID, "NEW", time.Now())
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *OrderPostgres) GetAll(userID int) ([]gofermart.OrdersOut, error) {
	var orders []gofermart.OrdersOut
	query := fmt.Sprintf("SELECT number, uploaded_at FROM %s WHERE user_id = $1", ordersTable)
	err := r.db.Select(&orders, query, userID)
	return orders, err

}
