package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/AltynayK/go-musthave-diploma-tpl/configs"
	"github.com/AltynayK/go-musthave-diploma-tpl/pkg/repository"
	"github.com/jmoiron/sqlx"
)

const (
	readTimeout     = 10 * time.Second
	writeTimeout    = 10 * time.Second
	maxHeaderBytes  = 1 << 20
	shutdownTimeout = 5 * time.Second
)

type Server struct {
	httpServer *http.Server
	config     *configs.Config
	db         *sqlx.DB
	repos      repository.Repository
	addr       string
}

func NewServer() *Server {
	config := configs.NewConfig()
	db := repository.NewPostgresDB(config)
	repos := repository.NewRepository(config)
	addr := config.RunAddress
	return &Server{
		config: config,
		db:     db,
		repos:  repos,
		addr:   addr,
	}
}
func (s *Server) Run(ctx context.Context) error {

	s.httpServer = &http.Server{
		Addr:           NewServer().addr,
		Handler:        NewHandler().InitRoutes(),
		MaxHeaderBytes: maxHeaderBytes,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
	}
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Print("listen and serve:")
		}
	}()
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	close(NewHandler().queueForAccrual)
	NewServer().db.Close()
	if err := s.Shutdown(shutdownCtx); err != nil {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
