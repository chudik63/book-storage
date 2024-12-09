package main

import (
	"book-storage/internal/config"
	"book-storage/internal/pkg/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
