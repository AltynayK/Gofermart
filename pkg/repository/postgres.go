package repository

import (
	"fmt"
	"time"

	"github.com/AltynayK/go-musthave-diploma-tpl/configs"
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable  = "users"
	ordersTable = "orders"
)

type DataBase struct {
	config *configs.Config
	db     *sqlx.DB
}

func NewDataBase(config *configs.Config) Repository {
	return &DataBase{
		config: config,
		db:     NewPostgresDB(config),
	}
}

func NewRepository(config *configs.Config) Repository {
	return NewDataBase(config)
}
func NewPostgresDB(config *configs.Config) *sqlx.DB {
	db, err := sqlx.Open("postgres", configs.DatabaseURI)
	if err != nil {
		fmt.Print(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Print(err)
	}
	CreateTableUsers(db)
	CreateTableOrders(db)
	return db

}

func CreateTableUsers(db *sqlx.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id serial primary key, login varchar UNIQUE, password varchar, current decimal DEFAULT '0')")
	if err != nil {
		fmt.Print(err)
	}

}
func CreateTableOrders(db *sqlx.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS orders (id serial primary key, number varchar UNIQUE, user_id int, status varchar, accrual decimal DEFAULT '0', uploaded_at timestamptz NOT NULL, withdrawn decimal DEFAULT '0', processed_at timestamptz DEFAULT CURRENT_DATE)")
	if err != nil {
		fmt.Print(err)
	}

}

func (r *DataBase) CreateUser(user models.User) error {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password) values ($1, $2) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (r *DataBase) GetUser(login, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, login, password)
	return user, err
}
func (r *DataBase) CreateOrder(userID int, number string) error {
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
func (r *DataBase) GetOrderByUserAndNumber(userID int, number int) ([]models.OrdersOut, error) {
	var orders []models.OrdersOut
	query := fmt.Sprintf("SELECT number FROM %s WHERE user_id = $1 AND number=$2", ordersTable)
	err := r.db.Select(&orders, query, userID, number)
	return orders, err

}
func (r *DataBase) GetOrder(number int) ([]models.OrdersOut, error) {
	var orders []models.OrdersOut
	query := fmt.Sprintf("SELECT number FROM %s WHERE number=$1", ordersTable)
	err := r.db.Select(&orders, query, number)
	return orders, err

}
func (r *DataBase) GetAllOrders(userID int) ([]models.OrdersOut, error) {
	var orders []models.OrdersOut
	query := fmt.Sprintf("SELECT number, status, accrual, uploaded_at FROM %s WHERE user_id = $1 ORDER BY uploaded_at DESC", ordersTable)
	err := r.db.Select(&orders, query, userID)
	return orders, err

}

func (r *DataBase) PostWithdrawBalance(order models.Withdrawals) (int64, error) {
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
func (r *DataBase) PostNewWithdrawBalance(order models.Withdrawals, userID int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	createOrder := fmt.Sprintf("INSERT INTO %s (number, withdrawn, processed_at, uploaded_at, user_id) VALUES ($1, $2, $3, $4, $5)", ordersTable)
	tx.QueryRow(createOrder, order.Order, order.Sum, time.Now(), time.Now(), userID)

	return tx.Commit()

}

//получение баланса пользователя
func (r *DataBase) GetUserCurrent(userID int) (float32, error) {
	row := r.db.QueryRow("SELECT current FROM users WHERE id = $1", userID)
	data := models.UserBalance{}
	err := row.Scan(&data.Current)
	return data.Current, err
}

//получение общей списанной суммы
func (r *DataBase) GetUserWithdrawn(userID int) (float32, error) {
	row := r.db.QueryRow("SELECT SUM(withdrawn) FROM orders WHERE user_id = $1", userID)
	data := models.UserBalance{}
	err := row.Scan(&data.Withdrawn)
	return data.Withdrawn, err
}

//обновление баланса пользователя
func (r *DataBase) UpdateUserBalance(userID int, current float32) (int64, error) {
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

func (r *DataBase) GetAllWithdrawals(userID int) ([]models.Withdrawals, error) {
	var withdrawals []models.Withdrawals
	query := fmt.Sprintf("SELECT number, withdrawn, processed_at FROM %s WHERE user_id = $1 AND withdrawn!=$2 ORDER BY processed_at DESC", ordersTable)
	err := r.db.Select(&withdrawals, query, userID, 0)
	return withdrawals, err

}

func (r *DataBase) PostBalance(order models.OrderBalance) (int64, error) {
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

func (r *DataBase) GetOrderUserID(number string) (int, error) {
	row := r.db.QueryRow("SELECT user_id FROM orders WHERE number = $1", number)
	userID := models.Orders{}
	err := row.Scan(&userID.UserID)
	return userID.UserID, err
}
