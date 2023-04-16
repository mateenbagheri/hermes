package main

import (
	"time"

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
	// because we can not connect directly to influx directory.
	// in production there is no need for sleep function. because
	// the server stays up all the time.
	time.Sleep(2 * time.Second)
}
