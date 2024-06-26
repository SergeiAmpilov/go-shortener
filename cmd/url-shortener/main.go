package main

import (
	"log/slog"
	"os"
	"url-shortener/internal/config"
	mwLogger "url-shortener/internal/http-server/middleware/logger"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage/sqlite"
	"url-shortener/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	/*

		err = stor.DeleteURL("bit.ly/c/1234567891")

		if err != nil {
			log.Error("failed to delete url", sl.Err(err))
		}
		//////


		resUrl, err := stor.GetURL("bit.ly/c/1234567891")
		if err != nil {
			fmt.Errorf("%w", err)
			os.Exit(1)
		}

		fmt.Println(resUrl)
		/////


			resAdd, err := stor.SaveURL(
				"https://www.youtube.com/watch?v=rCJvW2xgnk0&t=2633s",
				"bit.ly/c/1234567891",
			)
			if err != nil {
				log.Error("failed to add link", sl.Err(err))
				os.Exit(1)
			}
			fmt.Println(resAdd)
	*/

	// server -
}
