package main

import (
	"log/slog"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage/sqlite"
	"url-shortener/logger"
)

func main() {
	// config - cleanenv
	cfg := *config.New()

	// logger - slog - import "log/slog"
	log := logger.SetupLogger(cfg.Config.ENV)

	log.Info("Starting url shortener", slog.String("env", cfg.Config.ENV))
	log.Debug("debug messages are enabled")

	// storage - sqlite

	stor, err := sqlite.New(cfg.Config.STORAGE_PATH)

	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = stor

	// router - chi, chi render

	// server -
}
