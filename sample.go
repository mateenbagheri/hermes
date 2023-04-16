package main

import "github.com/doki-programs/hermes/logger"

// this main function is not necc
func main() {
	logger := logger.New(logger.ZeroLoggerType).
		WithLevel(logger.DebugLevel).
		WithServiceName("test").
		WithWriters(
			logger.ConsoleWriterType,
			logger.FileWriterType,
			logger.DatabaseWriterType,
		).
		Build().
		WithScope("main")

	logger.Debug("this is a test")
}
