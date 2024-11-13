package http

import (
	"book-storage/internal/config"
	"context"
	"net/http"
	"time"
)

const (
	_defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	httpServer      *http.Server
	shutdownTimeout time.Duration
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.Server.HttpServerPort,
			Handler: handler,
		},
		shutdownTimeout: _defaultShutdownTimeout,
	}
}

func (s *Server) Run() error {
	err := s.httpServer.ListenAndServe()
	if err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	ctx, shutdown := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer shutdown()

	return s.httpServer.Shutdown(ctx)
}
