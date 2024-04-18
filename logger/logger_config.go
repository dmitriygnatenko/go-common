package logger

import "log/slog"

type LoggerConfig struct {
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
	smtpPort          uint
	smtpUser          string
	smtpPassword      string
	email             string
	subject           string
}

type LoggerConfigOption func(*LoggerConfig)

type LoggerConfigOptions []LoggerConfigOption

func (s *LoggerConfigOptions) Add(option LoggerConfigOption) {
	*s = append(*s, option)
}

func NewConfig(opts ...LoggerConfigOption) LoggerConfig {
	c := &LoggerConfig{}
	for _, opt := range opts {
		opt(c)
	}

	return *c
}

// stdout log

func WithStdoutLogEnabled(enabled bool) LoggerConfigOption {
	return func(s *LoggerConfig) {
		s.stdoutLogEnabled = enabled
	}
}

func WithStdoutLogLevel(level slog.Level) LoggerConfigOption {
	return func(s *LoggerConfig) {
		s.stdoutLogLevel = level
	}
}
func WithStdoutLogAddSource(add bool) LoggerConfigOption {
	return func(s *LoggerConfig) {
		s.stdoutLogAddSource = add
	}
}

// file log

func WithFileLogEnabled(enabled bool) LoggerConfigOption {
	return func(s *LoggerConfig) {
		s.fileLogEnabled = enabled
	}
}

func WithFileLogLevel(level slog.Level) LoggerConfigOption {
	return func(s *LoggerConfig) {
		s.fileLogLevel = level
	}
}

func WithFileLogAddSource(add bool) LoggerConfigOption {
	return func(s *LoggerConfig) {
		s.fileLogAddSource = add
	}
}

func WithFilepath(path string) LoggerConfigOption {
	return func(s *LoggerConfig) {
		s.filepath = path
	}
}

// email log

func WithEmailLogEnabled(enabled bool) LoggerConfigOption {
	return func(s *LoggerConfig) {
		s.emailLogEnabled = enabled
	}
}

func WithEmailLogLevel(level slog.Level) LoggerConfigOption {
	return func(s *LoggerConfig) {
		s.emailLogLevel = level
	}
}

func WithEmailLogAddSource(add bool) LoggerConfigOption {
	return func(s *LoggerConfig) {
		s.emailLogAddSource = add
	}
}

func WithEmailRecipient(email string) LoggerConfigOption {
	return func(s *LoggerConfig) {
		s.email = email
	}
}

func WithEmailSubject(subject string) LoggerConfigOption {
	return func(s *LoggerConfig) {
		s.subject = subject
	}
}

func WithSMTPHost(host string) LoggerConfigOption {
	return func(s *LoggerConfig) {
		s.smtpHost = host
	}
}

func WithSMTPPort(port uint) LoggerConfigOption {
	return func(s *LoggerConfig) {
		s.smtpPort = port
	}
}

func WithSMTPUser(user string) LoggerConfigOption {
	return func(s *LoggerConfig) {
		s.smtpUser = user
	}
}

func WithSMTPPassword(password string) LoggerConfigOption {
	return func(s *LoggerConfig) {
		s.smtpPassword = password
	}
}
