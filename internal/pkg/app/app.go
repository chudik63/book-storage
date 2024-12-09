package app

import (
	"book-storage/internal/config"
	"book-storage/internal/transport/http"
	"book-storage/pkg/logger"
	"book-storage/pkg/postgres"
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func Run(ctx context.Context, cfg *config.Config) {
	logs := logger.GetLoggerFromCtx(ctx)

	db, err := postgres.New(ctx, &cfg.Config)
	if err != nil {
		logs.Fatal(ctx, zap.String("err", err.Error()))
	}

	_ = db

	// create repos
	// create services

	mux := http.NewBookStorageMux(ctx)

	httpServer := http.NewServer(cfg, mux)

	go func() {
		if err := httpServer.Run(); err != nil {
			logs.Error(ctx, "failed starting the server", zap.String("err: ", err.Error()))
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c

	if err := httpServer.Stop(); err != nil {
		logs.Error(ctx, "failed shutting down the server", zap.String("err: ", err.Error()))
	}
}
