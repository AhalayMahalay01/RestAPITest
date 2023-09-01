package main

import (
	"RestAPITest/internal/config"
	"RestAPITest/internal/http-server/handlers/url/save"
	"RestAPITest/internal/lib/logger/handlers/slogpretty"
	"RestAPITest/internal/storage/sqllite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
	"net/http"
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

	log.Info("starting url-shortener", slog.String("env", sfg.Env), slog.String("version", "13"))
	log.Debug("debug messages are enabled")
	//log.Error("error messages are enabled")

	storage, err := sqllite.New(sfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", slog.Error)
		return
	}

	/*id, err := storage.SaveURL("https://google.com", "google")
	if err != nil {
		log.Error("failed to save url", slog.Error)
		os.Exit(1)
	}

	log.Info("saved url", slog.Int64("id", id))

	//id, err = storage.SaveURL("https://google.com", "google")
	//if err != nil {
	//	log.Error("failed to save url", slog.Error)
	//	os.Exit(1)
	//}*/

	_ = storage

	router := chi.NewRouter()

	router.Use(middleware.RequestID) // RequestID дає можливість віднайти по даному ID для подальшої роботи з помилками
	// router.Use(middleware.RealIP) ще один можливий варіант
	//middleware Повертає користувача при неправильному вводу логіна або пароля

	router.Use(middleware.Logger) // логірує всі вхідні запити
	//router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer) // якщо виникає паніка всередині Handler
	router.Use(middleware.URLFormat) // для зручного написання url при підключенні їх до router

	router.Post("/url", save.New(log, storage))

	log.Info("starting server", slog.String("address", sfg.Address)) // виводимо повідомлення про запуск серверу, та на якій адресі він запускається

	srv := &http.Server{ //створюємо сам сервер
		Addr:         sfg.Address,
		Handler:      router,
		ReadTimeout:  sfg.HTTPServer.Timeout,
		WriteTimeout: sfg.HTTPServer.Timeout,
		IdleTimeout:  sfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil { // функція для блокування виклику під час виникнення помилки
		log.Error("failed to start server")
	}

	log.Error("server STOP!")

	// TODO: run  server:
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = setupPrettySlog()
		/*log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)*/
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

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
