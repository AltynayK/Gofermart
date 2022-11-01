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
	chanVal         = 5
)

type Server struct {
	httpServer      *http.Server
	config          *configs.Config
	db              *sqlx.DB
	repos           *repository.Repository
	queueForAccrual chan string
}

func NewServer() *Server {

	config := configs.NewConfig()
	db := repository.NewPostgresDB(config)
	repos := repository.NewRepository(db)
	return &Server{
		config:          config,
		db:              db,
		repos:           repos,
		queueForAccrual: make(chan string, chanVal),
	}
}
func (s *Server) Run(ctx context.Context, handlers *Handler, handler http.Handler) error {

	s.httpServer = &http.Server{

		Addr:           NewServer().config.RunAddress,
		Handler:        handler,
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
	close(s.queueForAccrual)
	NewServer().db.Close()
	if err := s.Shutdown(shutdownCtx); err != nil {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
