package main

import (
	"log/slog"
	"url-shortener/internal/config"
	"url-shortener/logger"
)

func main() {
	// config - cleanenv
	cfg := *config.New()
	// fmt.Printf("%+v\n", *cfg)

	// logger - slog - import "log/slog"
	log := logger.SetupLogger(cfg.Config.ENV)

	log.Info("Starting url shortener", slog.String("env", cfg.Config.ENV))
	log.Debug("debug messages are enabled")

	// storage - sqlite

	// router - chi, chi render

	// server -
}
