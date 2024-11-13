package postgres

import (
	"book-storage/pkg/logger"
	"context"
	"database/sql"

	"fmt"

	_ "github.com/lib/pq"

	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type DB struct {
	db   *sql.DB
	psql squirrel.StatementBuilderType
}

func New(ctx context.Context, config Config) (*DB, error) {
	logs := logger.GetLoggerFromCtx(ctx)

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", config.User, config.Password, config.DBName, config.Host, config.Port)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logs.Error(ctx, "can`t connecting to database", zap.String("error:", err.Error()))
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logs.Error(ctx, "failed connecting to database", zap.String("error:", err.Error()))
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &DB{db, psql}, nil
}
