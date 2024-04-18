package logger

import "log/slog"

type Config struct {
	// stdout config
	stdoutLogEnabled   bool
	stdoutLogLevel     slog.Level // INFO by default
	stdoutLogAddSource bool

	// file config
	fileLogEnabled   bool
	fileLogLevel     slog.Level // INFO by default
	fileLogAddSource bool
	fileLogFilepath  string

	// email config
	emailLogEnabled   bool
	emailLogLevel     slog.Level // INFO by default
	emailLogAddSource bool
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

// stdout log

func WithStdoutLogEnabled(enabled bool) ConfigOption {
	return func(s *Config) {
		s.stdoutLogEnabled = enabled
	}
}

func WithStdoutLogLevel(level slog.Level) ConfigOption {
	return func(s *Config) {
		s.stdoutLogLevel = level
	}
}
func WithStdoutLogAddSource(add bool) ConfigOption {
	return func(s *Config) {
		s.stdoutLogAddSource = add
	}
}

// file log

func WithFileLogEnabled(enabled bool) ConfigOption {
	return func(s *Config) {
		s.fileLogEnabled = enabled
	}
}

func WithFileLogLevel(level slog.Level) ConfigOption {
	return func(s *Config) {
		s.fileLogLevel = level
	}
}

func WithFileLogAddSource(add bool) ConfigOption {
	return func(s *Config) {
		s.fileLogAddSource = add
	}
}

func WithFileLogFilepath(path string) ConfigOption {
	return func(s *Config) {
		s.fileLogFilepath = path
	}
}

// email log

func WithEmailLogEnabled(enabled bool) ConfigOption {
	return func(s *Config) {
		s.emailLogEnabled = enabled
	}
}

func WithEmailLogLevel(level slog.Level) ConfigOption {
	return func(s *Config) {
		s.emailLogLevel = level
	}
}

func WithEmailLogAddSource(add bool) ConfigOption {
	return func(s *Config) {
		s.emailLogAddSource = add
	}
}
