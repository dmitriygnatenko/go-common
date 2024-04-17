package logger

import (
	"log/slog"
	"time"
)

const (
	defaultTimeFormat = time.DateTime
)

type Config struct {
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
	addSource  bool
	timeFormat string
}

type ConfigOption func(*Config)

type ConfigOptions []ConfigOption

func (s *ConfigOptions) Add(option ConfigOption) {
	*s = append(*s, option)
}

func NewConfig(opts ...ConfigOption) Config {
	c := &Config{}
	for _, opt := range opts {
		opt(c)
	}

	return *c
}

func WithStdoutLogEnabled(enabled bool) ConfigOption {
	return func(s *Config) {
		s.stdoutLogEnabled = enabled
	}
}

func WithFileLogEnabled(enabled bool) ConfigOption {
	return func(s *Config) {
		s.fileLogEnabled = enabled
	}
}

func WithEmailLogEnabled(enabled bool) ConfigOption {
	return func(s *Config) {
		s.emailLogEnabled = enabled
	}
}

func WithStdoutLogLevel(level slog.Level) ConfigOption {
	return func(s *Config) {
		s.stdoutLogLevel = level
	}
}

func WithFileLogLevel(level slog.Level) ConfigOption {
	return func(s *Config) {
		s.fileLogLevel = level
	}
}

func WithEmailLogLevel(level slog.Level) ConfigOption {
	return func(s *Config) {
		s.emailLogLevel = level
	}
}

func WithAddSource(add bool) ConfigOption {
	return func(s *Config) {
		s.addSource = add
	}
}

func WithTimeFormat(format string) ConfigOption {
	return func(s *Config) {
		s.timeFormat = format
	}
}
