package handler

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

const (
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
	maxHeaderBytes = 1 << 20
)

func (s *Server) Run(RunAddress string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           RunAddress,
		Handler:        handler,
		MaxHeaderBytes: maxHeaderBytes,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
