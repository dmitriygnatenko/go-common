package main

import "git.dmitriygnatenko.ru/dima/go-common/logger"

func main() {

	l := logger.NewLogger(logger.NewConfig(
		logger.WithEmailLogEnabled(true),
		logger.WithStdoutLogEnabled(true),
		logger.WithFileLogEnabled(true),
	))

	l.Info("dfgdgdsg", "dsfsd", 232423, "aerew2", "val333")

}
