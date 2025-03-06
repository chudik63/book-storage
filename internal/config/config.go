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

	SMTPMail     string `env:"SMTP_MAIL"`
	SMTPHost     string `env:"SMTP_HOST"`
	SMTPPort     int    `env:"SMTP_PORT"`
	SMTPPassword string `env:"SMTP_PASSWORD"`

	VerificationSubject  string `env:"VERIFICATION_SUBJECT"`
	VerificationTemplate string `env:"VERIFICATION_TEMPLATE" env-default:"templates/verification_email.html"`

	Domain string `env:"DOMAIN" env-default:"localhost"`
}

func New() (*Config, error) {
	cfg := Config{}

	//err := cleanenv.ReadEnv(&cfg)

	err := cleanenv.ReadConfig(".env", &cfg)

	if cfg == (Config{}) {
		return nil, errors.New("config is empty")
	}

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
