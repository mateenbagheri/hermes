package main

import (
	"github.com/doki-programs/hermes/logger"
	"github.com/joho/godotenv"
)

// this main function is not necc
func main() {
	err := godotenv.Load()
	if err != nil {
		panic("could not load environments")
	}

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
