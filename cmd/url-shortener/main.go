package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	// config - cleanenv

	cfg := config.New()

	fmt.Printf("%+v\n", *cfg)

	// logger - slog - import "log/slog"

	// storage - sqlite

	// router - chi, chi render

	// server -
}
