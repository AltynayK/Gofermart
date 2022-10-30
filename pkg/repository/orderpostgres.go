package repository

import (
	"fmt"
	"time"

	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/models"
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
func (r *OrderPostgres) GetOrderByUserAndNumber(userID int, number int) ([]models.OrdersOut, error) {
	var orders []models.OrdersOut
	query := fmt.Sprintf("SELECT number FROM %s WHERE user_id = $1 AND number=$2", ordersTable)
	err := r.db.Select(&orders, query, userID, number)
	return orders, err

}
func (r *OrderPostgres) GetOrder(number int) ([]models.OrdersOut, error) {
	var orders []models.OrdersOut
	query := fmt.Sprintf("SELECT number FROM %s WHERE number=$1", ordersTable)
	err := r.db.Select(&orders, query, number)
	return orders, err

}
func (r *OrderPostgres) GetAll(userID int) ([]models.OrdersOut, error) {
	var orders []models.OrdersOut
	query := fmt.Sprintf("SELECT number, status, accrual, uploaded_at FROM %s WHERE user_id = $1 ORDER BY uploaded_at DESC", ordersTable)
	err := r.db.Select(&orders, query, userID)
	return orders, err

}

func (r *OrderPostgres) PostWithdrawBalance(order models.Withdrawals) (int64, error) {
	query := fmt.Sprintf("UPDATE %s SET withdrawn=$2, processed_at=$3 WHERE number=$1", ordersTable)
	res, err := r.db.Exec(query, order.Order, order.Sum, time.Now())
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, err

}

//получение баланса пользователя
func (r *OrderPostgres) GetUserCurrent(userID int) (float32, error) {
	row := r.db.QueryRow("SELECT current FROM users WHERE id = $1", userID)
	data := models.UserBalance{}
	err := row.Scan(&data.Current)
	return data.Current, err
}

//получение общей списанной суммы
func (r *OrderPostgres) GetUserWithdrawn(userID int) (float32, error) {
	row := r.db.QueryRow("SELECT SUM(withdrawn) FROM orders WHERE user_id = $1", userID)
	data := models.UserBalance{}
	err := row.Scan(&data.Withdrawn)
	return data.Withdrawn, err
}

//обновление баланса пользователя
func (r *OrderPostgres) UpdateUserBalance(userID int, current float32) (int64, error) {
	query := `UPDATE users SET current=$2 WHERE id=$1; `
	res, err := r.db.Exec(query, userID, current)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *OrderPostgres) GetAllWithdrawals(userID int) ([]models.Withdrawals, error) {
	var withdrawals []models.Withdrawals
	query := fmt.Sprintf("SELECT number, withdrawn, processed_at FROM %s WHERE user_id = $1 AND withdrawn!=$2 ORDER BY processed_at DESC", ordersTable)
	err := r.db.Select(&withdrawals, query, userID, 0)
	return withdrawals, err

}

func (r *OrderPostgres) PostBalance(order models.OrderBalance) (int64, error) {
	query := fmt.Sprintf("UPDATE %s SET status=$2, accrual=$3 WHERE number=$1", ordersTable)
	res, err := r.db.Exec(query, order.Order, order.Status, order.Accrual)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, err

}

func (r *OrderPostgres) GetOrderUserID(number string) (int, error) {
	row := r.db.QueryRow("SELECT user_id FROM orders WHERE number = $1", number)
	userID := models.Orders{}
	err := row.Scan(&userID.UserID)
	return userID.UserID, err
}
