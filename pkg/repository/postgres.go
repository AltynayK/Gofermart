package repository

import (
	"database/sql"
	"fmt"

	"github.com/AltynayK/go-musthave-diploma-tpl/configs"
	_ "github.com/lib/pq"
)

const (
	usersTable  = "users"
	ordersTable = "orders"
)

func NewPostgresDB(config *configs.Config) *sql.DB {
	db, err := sql.Open("postgres", configs.DatabaseURI)

	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	CreateTableUsers(db)
	CreateTableOrders(db)
	return db

}

func CreateTableUsers(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id serial primary key, login varchar UNIQUE, password varchar, current int)")
	if err != nil {
		fmt.Print(err)
	}

}
func CreateTableOrders(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS orders (id serial primary key, number varchar UNIQUE, user_id int, status varchar, accrual int, uploaded_at varchar, withdrawn int)")
	if err != nil {
		fmt.Print(err)
	}

}
