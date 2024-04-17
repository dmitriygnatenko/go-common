package db

import (
	"time"
)

const (
	defaultDriver  = "mysql"
	defaultHost    = "localhost"
	defaultPort    = 3306
	defaultSslMode = "disabled"
)

type Config struct {
	driver string

	username string
	password string
	dbname   string
	host     string
	port     uint16
	sslMode  string

	maxOpenConns uint16
	maxIdleConns uint16

	maxConnLifetime     *time.Duration
	maxIdleConnLifetime *time.Duration
}

type ConfigOption func(*Config)

type ConfigOptions []ConfigOption

func (s *ConfigOptions) Add(option ConfigOption) {
	*s = append(*s, option)
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

func WithSSLMode(sslMode string) ConfigOption {
	return func(s *Config) {
		s.sslMode = sslMode
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
