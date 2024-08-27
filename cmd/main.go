package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"tz-kode/internal/apiserver"
	"tz-kode/internal/config"
	"tz-kode/internal/handlers"
	"tz-kode/internal/repository"
	"tz-kode/internal/service"
	"tz-kode/pkg/logger"
	"tz-kode/pkg/logger/sl"
	"tz-kode/pkg/postgres"

	"github.com/gorilla/sessions"
)

func main() {
	cfg := config.MustLoad()
	logger := logger.SetupLogger(cfg.Env)
	logger.Info("Starting app", slog.String("port", cfg.Address), slog.String("env", cfg.Env))
	logger.Debug("Debug messages are enabled")

	db, err := postgres.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		logger.Error("failed to initialize db: %s", sl.Err(err))
	}
	logger.Info("initialize database")
	repository := repository.NewRepository(db)
	logger.Info("initialize repository")
	services := service.NewService(repository)
	logger.Info("initialize services")
	sessionStore := sessions.NewCookieStore([]byte(cfg.SessionKey))
	handler := handlers.NewHandler(services, sessionStore)
	logger.Info("initialize handlers")

	srv := new(apiserver.APIServer)

	go func() {
		if err := srv.Start(cfg, handler.InitRouters()); err != nil {
			logger.Error("error occured while running http server: %s", sl.Err(err))
		}
	}()
	logger.Info("starting application")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logger.Info("application shutdown")
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Error("error occured on server shutdown: %s", sl.Err(err))
	}

	logger.Info("application shutdown")
	if err := db.Close(); err != nil {
		logger.Error("error occured on db connection close: %s", sl.Err(err))
	}
}
