package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/handler"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	srv := new(handler.Server)
	if err := srv.Run(ctx); err != nil {
		fmt.Println(err)
	}
}
