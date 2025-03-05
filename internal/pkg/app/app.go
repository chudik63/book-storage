package app

import (
	"book-storage/internal/config"
	"book-storage/internal/database/postgres"
	"book-storage/internal/repository"
	"book-storage/internal/server"
	"book-storage/internal/service"
	transport "book-storage/internal/transport/http"
	"book-storage/pkg/hasher"
	"book-storage/pkg/logger"
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func Run(ctx context.Context, cfg *config.Config) {
	// Infrastructure
	logs := logger.GetLoggerFromCtx(ctx)

	db := postgres.New(ctx, &cfg.Config)

	hasher := hasher.New(cfg.LocalParameter)

	// Repos, serivices and API handlers
	userRepository := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepository, hasher)

	handler := transport.NewHandler(userService, logs)

	// HTTP server
	srv := server.NewServer(cfg, handler.Init(cfg))

	go func() {
		if err := srv.Run(ctx); err != nil {
			logs.Error(ctx, "failed running http server", zap.String("err: ", err.Error()))
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c

	if err := srv.Stop(); err != nil {
		logs.Error(ctx, "failed shutting down the server", zap.String("err: ", err.Error()))
	}

	db.Close()
}
