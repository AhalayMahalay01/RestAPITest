package main

import (
	"RestAPITest/internal/config"
	"RestAPITest/internal/storage/sqllite"
	"golang.org/x/exp/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	sfg := config.MustLoad()

	log := setupLogger(sfg.Env)

	log.Info("starting url-shortener", slog.String("env", sfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := sqllite.New(sfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", slog.Error)
		return
	}

	id, err := storage.SaveURL("https://google.com", "google")
	if err != nil {
		log.Error("failed to save url", slog.Error)
		os.Exit(1)
	}

	log.Info("saved url", slog.Int64("id", id))

	//id, err = storage.SaveURL("https://google.com", "google")
	//if err != nil {
	//	log.Error("failed to save url", slog.Error)
	//	os.Exit(1)
	//}

	_ = storage

	// TODO: init router: chi, "chi render"

	// TODO: run  server:
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log

}
