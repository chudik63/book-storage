package main

import (
	"book-storage/internal/config"
	"book-storage/internal/pkg/app"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
