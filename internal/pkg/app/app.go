package app

import (
	"book-storage/internal/config"
	"book-storage/pkg/logger"
	"book-storage/pkg/postgres"
	"context"
)

func Run(cfg *config.Config) {
	ctx := context.Background()

	mainLogger := logger.New()
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)

	db, err := postgres.New(ctx, cfg)
	if err != nil {
		panic(err)
	}

	_ = db

	// create repos
	// create services
}
