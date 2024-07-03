package closer

import (
	"context"
	"time"
)

type Config struct {
	timeout *time.Duration
	logger  Logger
}

type ConfigOption func(*Config)

type ConfigOptions []ConfigOption

type Logger interface {
	Info(ctx context.Context, msg string)
	Errorf(ctx context.Context, format string, args ...any)
}

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

func WithTimeout(timeout time.Duration) ConfigOption {
	return func(s *Config) {
		s.timeout = &timeout
	}
}

func WithLogger(logger Logger) ConfigOption {
	return func(s *Config) {
		s.logger = logger
	}
}
