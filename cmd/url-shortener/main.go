package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	// config - cleanenv

	cfg := config.MustLoad()

	fmt.Println(*cfg)

	// logger - slog - import "log/slog"

	// storage - sqlite

	// router - chi, chi render

	// server -
}
