package main

import (
	"book-storage/internal/config"
	"book-storage/internal/pkg/app"
	"book-storage/pkg/logger"
	"book-storage/pkg/migrator"
	"context"

	"go.uber.org/zap"
)

const (
	service = "bookStorage"
)

func main() {
	mainLogger, err := logger.New(service)
	if err != nil {
		panic(err)
	}

	ctx := context.WithValue(context.Background(), logger.LoggerKey, mainLogger)

	cfg, err := config.New()
	if err != nil {
		mainLogger.Fatal(ctx, "can`t load config", zap.String("err", err.Error()))
	}

	migrator.Start(ctx, cfg)

	app.Run(ctx, cfg)
}
