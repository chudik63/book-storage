package postgres

import (
	"book-storage/internal/config"
	"book-storage/pkg/logger"
	"database/sql"

	"fmt"

	_ "github.com/lib/pq"

	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

type DB struct {
	Db      *sql.DB
	Builder squirrel.StatementBuilderType
}

func New(logs logger.Logger, config *config.Config) (*DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", config.Database.User, config.Database.Password, config.Database.DBName, config.Database.Host, config.Database.DBPort)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logs.Error("can`t connecting to database", zap.String("error:", err.Error()))
		return nil, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logs.Error("failed connecting to database", zap.String("error:", err.Error()))
		return nil, err
	}

	logs.Info("database connected")

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &DB{db, psql}, nil
}
