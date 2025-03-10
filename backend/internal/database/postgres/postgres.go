package postgres

import (
	"book-storage/pkg/logger"
	"context"
	"database/sql"

	"fmt"

	_ "github.com/lib/pq"

	"go.uber.org/zap"
)

type Config struct {
	UserName string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	DBName   string `env:"POSTGRES_DB"`
}

type DB struct {
	*sql.DB
}

func New(ctx context.Context, config *Config) DB {
	logs := logger.GetLoggerFromCtx(ctx)

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", config.UserName, config.Password, config.DBName, config.Host, config.Port)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logs.Fatal(ctx, "can`t connect to database", zap.String("error:", err.Error()))
	}

	if err := db.Ping(); err != nil {
		logs.Fatal(ctx, "failed connecting to database", zap.String("error:", err.Error()))
	}

	logs.Debug(ctx, "database connected", zap.String("dsn", dsn))

	return DB{db}
}
