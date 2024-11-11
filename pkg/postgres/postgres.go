package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
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
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", config.User, config.Password, config.DBName, config.Host, config.Port)
	db, err := sql.Open("postgres", dsn)
	if err != nil {

	}
	defer db.Close()

	if err := db.Ping(); err != nil {

	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &DB{db, psql}, nil
}
