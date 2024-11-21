package logger

import "log/slog"

type SMTPClient interface {
	Send(recipient string, subject string, content string, html bool) error
}

type Config struct {
	// stdout config
	stdoutLogEnabled bool
	stdoutLogLevel   slog.Level // INFO by default

	// file config
	fileLogEnabled bool
	fileLogLevel   slog.Level // INFO by default
	filepath       string

	// email config
	emailLogEnabled bool
	emailLogLevel   slog.Level // INFO by default
	smtpClient      SMTPClient
	emailRecipient  string
	emailSubject    string
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

// stdout

func WithStdoutLogEnabled(enabled bool) ConfigOption {
	return func(s *Config) {
		s.stdoutLogEnabled = enabled
	}
}

func WithStdoutLogLevel(level string) ConfigOption {
	return func(s *Config) {
		var l slog.Level
		if err := l.UnmarshalText([]byte(level)); err == nil {
			s.stdoutLogLevel = l
		}
	}
}

// file

func WithFileLogEnabled(enabled bool) ConfigOption {
	return func(s *Config) {
		s.fileLogEnabled = enabled
	}
}

func WithFileLogLevel(level string) ConfigOption {
	return func(s *Config) {
		var l slog.Level
		if err := l.UnmarshalText([]byte(level)); err == nil {
			s.fileLogLevel = l
		}
	}
}

func WithFilepath(path string) ConfigOption {
	return func(s *Config) {
		s.filepath = path
	}
}

// email

func WithEmailLogEnabled(enabled bool) ConfigOption {
	return func(s *Config) {
		s.emailLogEnabled = enabled
	}
}

func WithEmailLogLevel(level string) ConfigOption {
	return func(s *Config) {
		var l slog.Level
		if err := l.UnmarshalText([]byte(level)); err == nil {
			s.emailLogLevel = l
		}
	}
}

func WithEmailRecipient(email string) ConfigOption {
	return func(s *Config) {
		s.emailRecipient = email
	}
}

func WithEmailSubject(subject string) ConfigOption {
	return func(s *Config) {
		s.emailSubject = subject
	}
}

func WithSMTPClient(c SMTPClient) ConfigOption {
	return func(s *Config) {
		s.smtpClient = c
	}
}
