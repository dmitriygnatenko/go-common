package ttl_memory_cache

import "time"

type Config struct {
	expiration *time.Duration
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

func WithExpiration(expiration time.Duration) ConfigOption {
	return func(s *Config) {
		s.expiration = &expiration
	}
}
