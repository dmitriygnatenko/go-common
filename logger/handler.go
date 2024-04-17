package logger

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"sync"
)

type Handler struct {
	config  HandlerConfig
	handler slog.Handler
	mu      *sync.Mutex
	buf     *bytes.Buffer
}

type HandlerConfig struct {
	// stdout config
	stdoutLogEnabled bool
	stdoutLogLevel   slog.Level // INFO by default

	// file config
	fileLogEnabled bool
	fileLogLevel   slog.Level // INFO by default

	// email config
	emailLogEnabled bool
	emailLogLevel   slog.Level // INFO by default

	// common config
	minLogLevel slog.Level
	timeFormat  string
}

func NewHandler(c Config) Handler {
	buf := &bytes.Buffer{}
	minLogLevel := getMinLogLevel(c)

	if len(c.timeFormat) == 0 {
		c.timeFormat = defaultTimeFormat
	}

	return Handler{
		buf: buf,
		mu:  &sync.Mutex{},

		handler: slog.NewJSONHandler(buf, &slog.HandlerOptions{
			AddSource: c.addSource,
			Level:     minLogLevel,
		}),

		config: HandlerConfig{
			stdoutLogEnabled: c.stdoutLogEnabled,
			fileLogEnabled:   c.fileLogEnabled,
			emailLogEnabled:  c.emailLogEnabled,
			stdoutLogLevel:   c.stdoutLogLevel,
			fileLogLevel:     c.fileLogLevel,
			emailLogLevel:    c.emailLogLevel,
			minLogLevel:      minLogLevel,
			timeFormat:       c.timeFormat,
		},
	}
}

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	message := r.Time.Format(h.config.timeFormat) + " " + r.Level.String() + " " + r.Message

	var attrs map[string]any

	r.Attrs(func(attr slog.Attr) bool {

		return true
	})

	if h.config.stdoutLogEnabled {
		fmt.Println(message)
	}

	return nil
}

func (h *Handler) Enabled(_ context.Context, l slog.Level) bool {
	return l >= h.config.minLogLevel
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{handler: h.handler.WithAttrs(attrs), buf: h.buf, mu: h.mu, config: h.config}
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{handler: h.handler.WithGroup(name), buf: h.buf, mu: h.mu, config: h.config}
}

func getMinLogLevel(c Config) slog.Level {
	minLogLevel := slog.LevelError

	if c.stdoutLogEnabled && c.stdoutLogLevel < minLogLevel {
		minLogLevel = c.stdoutLogLevel
	}

	if c.fileLogEnabled && c.fileLogLevel < minLogLevel {
		minLogLevel = c.fileLogLevel
	}

	if c.emailLogEnabled && c.emailLogLevel < minLogLevel {
		minLogLevel = c.fileLogLevel
	}

	return minLogLevel
}
