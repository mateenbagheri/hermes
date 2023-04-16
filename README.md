# Hermes
>Hermes was the messenger of the gods in Greek mythology, known for his cunning, speed, and versatility. He was the son of Zeus and the nymph Maia, and he was born in a cave on Mount Cyllene in Arcadia.

This repository provides a logger service that implements the ZeroLog repository and supports three types of logger writers:

* DB writer: logs messages to a database (there is an influxDB implementation integrated too)
* CLI writer: logs messages to the console (stdout)
* File writer: logs messages to a file

## Installation
To use the logger service, you'll need to install the package using Go modules:

```
go get github.com/your-username/logger-service
```

## Usage
To use the logger service, you'll first need to create a logger instance using one of the available writer types:

```go
import logger "github.com/doki-programs/hermes"

logger := logger.New(logger.ZeroLoggerType).
	WithLevel(logger.DebugLevel).
	WithServiceName("test").
	WithWriters(
      	// Create a new logger instance with a CLI writer
		logger.ConsoleWriterType,
      	// Create a new logger instance with a file writer
		logger.FileWriterType,
      	// Create a new logger instance with an InfluxDB writer
		logger.DatabaseWriterType,
	).
	Build().
	WithScope("main")

logger.Debug("this is a test")
```
Once you have a logger instance, you can use it to log messages at different levels:

```go
logger.Trace("This is a trace message")
logger.Debug("This is a debug message")
logger.Info("This is an info message")
logger.Warn("This is a warning message")
logger.Error("This is an error message")
logger.Fatal("This is a fatal message")
logger.Panic("This is a panic message")
```

You can also log messages with additional key-value pairs:

```go
logger.Tracev("This is a trace message with additional data", "key1", "value1", "key2", "value2")
```

## Configuration
You can configure the logger service by setting environment variables or using a configuration file named `.env` in main directory. See the configuration documentation at `.env.example` for more details.

## Contributing
Thank you for considering contributing to Hermes! Please follow these guidelines to ensure your contribution is properly considered:

1. Fork the repository and create your branch from main.
2. Make your changes, and add new tests as appropriate.
3. Run go test to make sure all tests pass.
4. Format your code with gofmt -s.
5. Ensure your code passes golint.
6. Commit your changes and push your branch to your forked repository.
Create a pull request to the main Hermes repository.

If you would like to report a bug or suggest a new feature, please feel free to open a GitHub issue in the Hermes repository. We welcome all feedback and suggestions!




