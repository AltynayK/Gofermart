package configs

import (
	"flag"
	"os"
)

type Config struct {
	RunAddress           string
	AccrualSystemAddress string
	DatabaseURI          string
}

var (
	RunAddress           string
	AccrualSystemAddress string
	DatabaseURI          string
)

func init() {
	//increment#5
	flag.StringVar(&RunAddress, "a", "127.0.0.1:8080", "RunAddress - адрес запуска HTTP-сервера")
	flag.StringVar(&AccrualSystemAddress, "r", "", "AccrualSystemAddress")
	//flag.StringVar(&DatabaseURI, "d", "host=localhost port=5432 user=altynay password=password dbname=somedb sslmode=disable", "DatabaseURI")
	flag.StringVar(&DatabaseURI, "d", "", "DatabaseURI")
	//flag.StringVar(&DatabaseURI, "d", "", "DatabaseURI")
}

func NewConfig() *Config {
	flag.Parse()
	if u, f := os.LookupEnv("RUN_ADDRESS"); f {
		RunAddress = u
	}
	if u, flg := os.LookupEnv("ACCRUAL_SYSTEM_ADDRESS"); flg {
		AccrualSystemAddress = u
	}
	if u, f := os.LookupEnv("DATABASE_URI"); f {
		DatabaseURI = u
	}
	return &Config{
		RunAddress:           RunAddress,
		AccrualSystemAddress: AccrualSystemAddress,
		DatabaseURI:          DatabaseURI,
	}
}
