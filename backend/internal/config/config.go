package config

import (
	"book-storage/internal/database/postgres"
	"errors"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	postgres.Config
	MigrationsPath string `env:"MIGRATIONS_PATH" env-default:"migrations"`
	ServerPort     string `env:"SERVER_PORT" env-default:"8080"`
	LocalParameter string `env:"LOCAL_PARAMETER" env-default:""`
}

func New() (*Config, error) {
	cfg := Config{}

	err := cleanenv.ReadEnv(&cfg)

	if cfg == (Config{}) {
		return nil, errors.New("config is empty")
	}

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
