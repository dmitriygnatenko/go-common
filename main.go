package main

import (
	"context"
	"log/slog"

	"git.dmitriygnatenko.ru/dima/go-common/logger"
)

func main() {

	err := logger.Init(logger.NewConfig(
		logger.WithStdoutLogAddSource(true),
		logger.WithStdoutLogEnabled(true),
		logger.WithFileLogLevel(slog.LevelError),
		logger.WithEmailLogEnabled(true),
		logger.WithFileLogEnabled(true),
		logger.WithFileLogFilepath("test_log.txt"),
	))

	_ = err

	ctx := context.Background()

	ctx = logger.With(ctx, "test111", 74658743)

	logger.InfoKV(ctx, "dfgdgdfgdssg", "dsfsd", "val333", "dfgdf", 11)

	// logger.Error(ctx, "dsfdsf dsf dsfs")

}
