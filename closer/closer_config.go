package closer

import "time"

type Config struct {
	timeout *time.Duration
	logger  Logger
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
