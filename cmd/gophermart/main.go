package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/handler"
)

func main() {
	handlers := handler.NewHandler()
	srv := new(handler.Server)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	if err := srv.Run(ctx, handlers, handlers.InitRoutes()); err != nil {
		fmt.Print(err)
	}
}
