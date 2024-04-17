package logger

import (
	"log/slog"
)

func NewLogger(c Config) *slog.Logger {
	handler := NewHandler(c)
	return slog.New(&handler)
}
