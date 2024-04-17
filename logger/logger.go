package logger

import (
	"log/slog"
)

func NewLogger(c Config) *slog.Logger {
	return slog.New(NewHandler(c))
}
