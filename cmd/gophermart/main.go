package main

import (
	"fmt"

	"github.com/AltynayK/go-musthave-diploma-tpl/configs"
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/handler"
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/repository"
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/service"
)

func main() {
	config := configs.NewConfig()
	db := repository.NewPostgresDB(config)
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(handler.Server)
	if err := srv.Run(config.RunAddress, handlers.InitRoutes()); err != nil {
		fmt.Print(err)
	}
}
