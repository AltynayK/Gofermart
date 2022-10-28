package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

const (
	readTimeout     = 10 * time.Second
	writeTimeout    = 10 * time.Second
	maxHeaderBytes  = 1 << 20
	shutdownTimeout = 5 * time.Second
)

func (s *Server) Run(ctx context.Context, handlers *Handler, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           handlers.config.RunAddress,
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
	close(handlers.queueForAccrual)
	handlers.db.Close()
	if err := s.Shutdown(shutdownCtx); err != nil {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
