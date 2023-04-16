# Hermes
This repository provides a logger service that implements the ZeroLog repository and supports three types of logger writers:

InfluxDB writer: logs messages to an InfluxDB database
*CLI writer: logs messages to the console (stdout)
*File writer: logs messages to a file

## Installation
To use the logger service, you'll need to install the package using Go modules:

```
go get github.com/your-username/logger-service
```

## Usage
To use the logger service, you'll first need to create a logger instance using one of the available writer types:

```go
import "github.com/your-username/logger-service"

// Create a new logger instance with an InfluxDB writer
logger := logger.NewLogger(logger.InfluxDBWriter, "my-influxdb-url", "my-influxdb-token", "my-influxdb-org", "my-influxdb-bucket")

// Create a new logger instance with a CLI writer
logger := logger.NewLogger(logger.CLIWriter)

// Create a new logger instance with a file writer
logger := logger.NewLogger(logger.FileWriter, "/path/to/logfile.log")
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

Configuration
You can configure the logger service by setting environment variables or using a configuration file. See the configuration documentation for more details.

Contributing
Contributions to this repository are welcome! See the contributing guidelines for more information.
