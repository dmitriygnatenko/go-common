package lru_memory_cache

type Config struct {
	capacity uint
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

func WithCapacity(capacity uint) ConfigOption {
	return func(s *Config) {
		s.capacity = capacity
	}
}
