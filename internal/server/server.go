package server

import (
	"book-storage/internal/config"
	"book-storage/pkg/logger"
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
)

const (
	_defaultShutdownTimeout = 5 * time.Second
)

type Server struct {
	httpServer      *http.Server
	shutdownTimeout time.Duration
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.ServerPort,
			Handler: handler,
		},
		shutdownTimeout: _defaultShutdownTimeout,
	}
}

func (s *Server) Run(ctx context.Context) error {
	log := logger.GetLoggerFromCtx(ctx)
	log.Info(ctx, "Server is running", zap.String("port", s.httpServer.Addr))

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
