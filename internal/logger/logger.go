package logger

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	logLevel := slog.LevelInfo

	if os.Getenv("DEBUG") == "1" {
		logLevel = slog.LevelDebug
	}

	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
}
