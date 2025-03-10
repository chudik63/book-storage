package config

import (
	"book-storage/internal/database/postgres"
	"errors"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	postgres.Config
	MigrationsPath string `env:"MIGRATIONS_PATH"`
	ServerPort     string `env:"SERVER_PORT"`
}

func New() (*Config, error) {
	cfg := Config{}

	// err := cleanenv.ReadEnv(&cfg)

	err := cleanenv.ReadConfig("configs/local.env", &cfg)

	if cfg == (Config{}) {
		return nil, errors.New("config is empty")
	}

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
