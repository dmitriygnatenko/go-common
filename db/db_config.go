package db

import (
	"time"
)

const (
	defaultDriver = "mysql"
	defaultHost   = "localhost"
	defaultPort   = 3306
)

type Config struct {
	driver string

	username string
	password string
	dbname   string
	host     string
	port     uint16

	maxOpenConns uint16
	maxIdleConns uint16

	maxOpenConnLifetime *time.Duration
	maxIdleConnLifetime *time.Duration
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

func WithDriver(driver string) ConfigOption {
	return func(s *Config) {
		s.driver = driver
	}
}

func WithUsername(username string) ConfigOption {
	return func(s *Config) {
		s.username = username
	}
}

func WithDatabase(dbname string) ConfigOption {
	return func(s *Config) {
		s.dbname = dbname
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

func WithMaxOpenConns(maxOpenConns uint16) ConfigOption {
	return func(s *Config) {
		s.maxOpenConns = maxOpenConns
	}
}

func WithMaxIdleConns(maxIdleConns uint16) ConfigOption {
	return func(s *Config) {
		s.maxIdleConns = maxIdleConns
	}
}

func WithMaxOpenConnLifetime(lifetime time.Duration) ConfigOption {
	return func(s *Config) {
		s.maxOpenConnLifetime = &lifetime
	}
}

func WithMaxIdleConnLifetime(lifetime time.Duration) ConfigOption {
	return func(s *Config) {
		s.maxIdleConnLifetime = &lifetime
	}
}
