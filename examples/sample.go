package main

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mateenbagheri/hermes"
)

// this main function is not necc
func main() {
	err := godotenv.Load()
	if err != nil {
		panic("could not load environments")
	}

	logger := hermes.New(hermes.ZeroLoggerType).
		WithInfluxConfig(
			os.Getenv("INFLUX_ADDRESS"),
			os.Getenv("INFLUX_TOKEN"),
			os.Getenv("INFLUX_ORGANIZATION"),
			os.Getenv("INFLUX_BUCKET"),
		).
		WithLevel(hermes.DebugLevel).
		WithServiceName("test").
		WithWriters(
			hermes.ConsoleWriterType,
			hermes.FileWriterType,
			hermes.DatabaseWriterType,
		).
		Build().
		WithScope("main")

	logger.Debug("this is a test")
	// because we can not connect directly to influx directory.
	// in production there is no need for sleep function. because
	// the server stays up all the time.
	time.Sleep(2 * time.Second)
}
