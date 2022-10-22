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

func (r *OrderPostgres) Create(userID int, number string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	var id int
	createOrder := fmt.Sprintf("INSERT INTO %s (number, user_id, status, uploaded_at) VALUES ($1, $2, $3, $4) RETURNING id", ordersTable)
	row := tx.QueryRow(createOrder, number, userID, "NEW", time.Now())
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func (r *OrderPostgres) GetOrderByUserAndNumber(userID int, number int) ([]gofermart.OrdersOut, error) {
	var orders []gofermart.OrdersOut
	query := fmt.Sprintf("SELECT number FROM %s WHERE user_id = $1 AND number=$2", ordersTable)
	err := r.db.Select(&orders, query, userID, number)
	return orders, err

}
func (r *OrderPostgres) GetOrder(number int) ([]gofermart.OrdersOut, error) {
	var orders []gofermart.OrdersOut
	query := fmt.Sprintf("SELECT number FROM %s WHERE number=$1", ordersTable)
	err := r.db.Select(&orders, query, number)
	return orders, err

}
func (r *OrderPostgres) GetAll(userID int) ([]gofermart.OrdersOut, error) {
	var orders []gofermart.OrdersOut
	query := fmt.Sprintf("SELECT number, uploaded_at FROM %s WHERE user_id = $1 ORDER BY uploaded_at DESC", ordersTable)
	err := r.db.Select(&orders, query, userID)
	return orders, err

}

func (r *OrderPostgres) GetUserBalance(userID int) ([]gofermart.UserBalance, error) {
	var balance []gofermart.UserBalance
	query := fmt.Sprintf("SELECT current, withdrawn FROM %s WHERE id = $1", usersTable)
	err := r.db.Select(&balance, query, userID)
	return balance, err
}
