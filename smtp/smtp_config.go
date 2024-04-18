package smtp

const (
	defaultHost = "localhost"
	defaultPort = 587
)

type Config struct {
	host     string
	port     uint16
	username string
	password string
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

func WithUsername(username string) ConfigOption {
	return func(s *Config) {
		s.username = username
	}
}

func WithPassword(password string) ConfigOption {
	return func(s *Config) {
		s.password = password
	}
}

func WithHost(host string) ConfigOption {
	return func(s *Config) {
		s.host = host
	}
}

func WithPort(port uint16) ConfigOption {
	return func(s *Config) {
		s.port = port
	}
}
