package logger

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"sync"
)

type CtxAttrKey struct{}

var (
	once      sync.Once
	ctxAttrMu sync.RWMutex
	logger    *Logger
)

type Logger struct {
	config  Config
	logFile *os.File

	stdoutLogger *slog.Logger
	fileLogger   *slog.Logger
	emailLogger  *slog.Logger
}

func Init(c Config) error {
	var err error

	once.Do(func() {
		logger = &Logger{config: c}

		if c.stdoutLogEnabled {
			logger.stdoutLogger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: c.stdoutLogLevel,
			}))
		}

		if c.fileLogEnabled {
			logger.logFile, err = os.OpenFile(c.filepath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				return
			}

			logger.fileLogger = slog.New(slog.NewJSONHandler(logger.logFile, &slog.HandlerOptions{
				Level: c.fileLogLevel,
			}))
		}

		if c.emailLogEnabled {
			if c.smtpClient == nil {
				err = errors.New("empty SMTP client")
				return
			}

			ew, ewErr := NewEmailWriter(c.smtpClient, c.emailRecipient, c.emailSubject)
			if ewErr != nil {
				err = ewErr
				return
			}

			logger.emailLogger = slog.New(slog.NewJSONHandler(ew, &slog.HandlerOptions{
				Level: c.emailLogLevel,
			}))
		}
	})

	return err
}

func Default() *Logger {
	if logger == nil {
		logger = &Logger{}
	}

	return logger
}

func ErrorKV(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelError, msg, args...)
}

func WarnKV(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelWarn, msg, args...)
}

func InfoKV(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelInfo, msg, args...)
}

func DebugKV(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelDebug, msg, args...)
}

func Errorf(ctx context.Context, format string, args ...any) {
	log(ctx, slog.LevelError, fmt.Sprintf(format, args...))
}

func Warnf(ctx context.Context, format string, args ...any) {
	log(ctx, slog.LevelWarn, fmt.Sprintf(format, args...))
}

func Infof(ctx context.Context, format string, args ...any) {
	log(ctx, slog.LevelInfo, fmt.Sprintf(format, args...))
}

func Debugf(ctx context.Context, format string, args ...any) {
	log(ctx, slog.LevelDebug, fmt.Sprintf(format, args...))
}

func Error(ctx context.Context, msg string) {
	log(ctx, slog.LevelError, msg)
}

func Warn(ctx context.Context, msg string) {
	log(ctx, slog.LevelWarn, msg)
}

func Info(ctx context.Context, msg string) {
	log(ctx, slog.LevelInfo, msg)
}

func Debug(ctx context.Context, msg string) {
	log(ctx, slog.LevelDebug, msg)
}

func log(ctx context.Context, level slog.Level, msg string, args ...any) {
	if Default().stdoutLogger != nil {
		Default().stdoutLogger.Log(ctx, level, msg, append(args, AttrFromCtx(ctx)...)...)
	}

	if Default().fileLogger != nil {
		Default().fileLogger.Log(ctx, level, msg, append(args, AttrFromCtx(ctx)...)...)
	}

	if Default().emailLogger != nil {
		Default().emailLogger.Log(ctx, level, msg, append(args, AttrFromCtx(ctx)...)...)
	}
}

func Close() error {
	if Default().logFile != nil {
		if err := Default().logFile.Close(); err != nil {
			return err
		}
	}

	return nil
}
