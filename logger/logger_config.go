package logger

import (
	"log/slog"
)

const (
	defaultStdoutLogEnabled = true
)

type Config struct {
	stdoutLogEnabled bool
	fileLogEnabled   bool
	emailLogEnabled  bool

	// INFO by default
	stdoutLogLevel slog.Level
	fileLogLevel   slog.Level
	emailLogLevel  slog.Level
}

type ConfigOption func(*Config)

type ConfigOptions []ConfigOption

func (s *ConfigOptions) Add(option ConfigOption) {
	*s = append(*s, option)
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
