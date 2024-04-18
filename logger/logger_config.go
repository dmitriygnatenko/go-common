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
	filepath         string

	// email config
	emailLogEnabled   bool
	emailLogLevel     slog.Level // INFO by default
	emailLogAddSource bool
	smtpHost          string
	smtpPort          uint16
	smtpUsername      string
	smtpPassword      string
	emailRecipient    string
	emailSubject      string
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

// file

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

func WithSMTPHost(host string) ConfigOption {
	return func(s *Config) {
		s.smtpHost = host
	}
}

func WithSMTPPort(port uint16) ConfigOption {
	return func(s *Config) {
		s.smtpPort = port
	}
}

func WithSMTPUsername(user string) ConfigOption {
	return func(s *Config) {
		s.smtpUsername = user
	}
}

func WithSMTPPassword(password string) ConfigOption {
	return func(s *Config) {
		s.smtpPassword = password
	}
}
