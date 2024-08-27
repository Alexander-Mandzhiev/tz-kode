package logger

import (
	"log/slog"
	"os"
)

const (
	envDevelopment = "development"
	envProduction  = "production"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envDevelopment:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProduction:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
