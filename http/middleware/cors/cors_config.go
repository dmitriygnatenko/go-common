package cors

type Config struct {
	origin  string
	methods string
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

func WithOrigin(origin string) ConfigOption {
	return func(s *Config) {
		s.origin = origin
	}
}

func WithMethods(methods string) ConfigOption {
	return func(s *Config) {
		s.methods = methods
	}
}
