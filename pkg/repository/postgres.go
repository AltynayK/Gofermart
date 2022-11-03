package repository

import (
	"log"

	"github.com/AltynayK/go-musthave-diploma-tpl/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable  = "users"
	ordersTable = "orders"
)

func NewPostgresDB(config *configs.Config) *sqlx.DB {
	db, err := sqlx.Open("postgres", configs.DatabaseURI)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	CreateTableUsers(db)
	CreateTableOrders(db)
	return db

}

func CreateTableUsers(db *sqlx.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id serial primary key, login varchar UNIQUE, password varchar, current decimal DEFAULT '0')")
	if err != nil {
		log.Fatal(err)
	}

}
func CreateTableOrders(db *sqlx.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS orders (id serial primary key, number varchar UNIQUE, user_id int, status varchar, accrual decimal DEFAULT '0', uploaded_at timestamptz NOT NULL, withdrawn decimal DEFAULT '0', processed_at timestamptz DEFAULT CURRENT_DATE)")
	if err != nil {
		log.Fatal(err)
	}

}
