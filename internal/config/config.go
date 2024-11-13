package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		HttpServerPort string `yaml:"httpServerPort"`
	}
	Database struct {
		Host     string `yaml:"host"`
		DBPort   string `yaml:"dbport"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	}
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
