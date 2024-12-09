package main

import (
	"book-storage/internal/config"
	"book-storage/internal/pkg/app"
	"book-storage/pkg/logger"
	"book-storage/pkg/migrator"
	"context"

	"go.uber.org/zap"
)

func main() {
	mainLogger := logger.New()
	ctx := context.WithValue(context.Background(), logger.LoggerKey, mainLogger)

	cfg, err := config.New()
	if err != nil {
		mainLogger.Fatal(ctx, zap.String("err", err.Error()))
	}

	migrator.Start(ctx, cfg)

	app.Run(ctx, cfg)
}
