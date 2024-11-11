package app

import (
	"book-storage/internal/config"
	"book-storage/pkg/logger"
	"book-storage/pkg/postgres"
	"context"
)

const (
	serviceName = "bookStorage"
)

func Run(cfg *config.Config) {
	ctx := context.Background()

	mainLogger := logger.New(serviceName)
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)

	db, err := postgres.New(ctx, cfg.Config)
	if err != nil {
		panic(err)
	}

	_ = db

	// create repos
	// create services
}
