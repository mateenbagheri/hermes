package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/doki-programs/hermes/config/db"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/rs/zerolog"
)

// WriterType represents the type of writer to use for the logger.
type WriterType uint8

const (
	// ConsoleWriterType represents a console writer.
	ConsoleWriterType WriterType = iota
	// FileWriterType represents a file writer.
	FileWriterType
	// DatabaseWriterType represents a database writer.
	DatabaseWriterType
)

// Factory is a struct that holds the type, service name, and zerolog.Logger.
type Factory struct {
	typ        Type
	service    string
	zerologger *zerolog.Logger
}

// New is a function that creates and returns a new Factory.
func New(typ Type) *Factory {
	return &Factory{
		typ:        typ,
		zerologger: new(zerolog.Logger),
	}
}

// WithWriters is a function that sets the writers for the Factory.
func (f *Factory) WithWriters(writers ...WriterType) *Factory {
	switch f.typ {
	case ZeroLoggerType:
		logger := zerolog.
			New(zerolog.MultiLevelWriter(f.zerologWriters(writers...)...)).
			With().
			Timestamp().
			Logger()

		f.zerologger = &logger
	default:
		panic("typ not acceptable")
	}

	return f
}

// WithServiceName is a function that sets the service name for the Factory.
func (f *Factory) WithServiceName(name string) *Factory {
	f.service = name

	return f
}

// WithLevel is a function that sets the logging level for the Factory.
func (f *Factory) WithLevel(level Level) *Factory {
	switch f.typ {
	case ZeroLoggerType:
		*f.zerologger = f.zerologger.Level(zerolog.Level(level))
	default:
		panic("level not acceptable")
	}

	return f
}

// WithStackError is a function that sets the stack error for the Factory.
func (f *Factory) WithStackError() *Factory {
	if f.typ == ZeroLoggerType {
		*f.zerologger = f.zerologger.With().Stack().Logger()

		return f
	}

	return f
}

// Build is a function that builds and returns the Logger.
func (f *Factory) Build() Logger { //nolint
	switch f.typ {
	case ZeroLoggerType:
		*f.zerologger = f.zerologger.With().Timestamp().Logger()

		return &ZeroLogger{Logger: f.zerologger, service: f.service}
	default:
		return nil
	}
}

// zerologWriters is a function that sets the writers for the zerolog.
func (f *Factory) zerologWriters(writerTypes ...WriterType) []io.Writer {
	writers := make([]io.Writer, 0)

	for i := range writerTypes {
		switch writerTypes[i] {
		case ConsoleWriterType:
			f.appendZerologConsoleWriter(&writers)
		case FileWriterType:
			f.appendZerologFileWriter(f.service, &writers)
		case DatabaseWriterType:
			f.appendZerologInfluxWriter(&writers)
		default:
			panic("writer type not acceptable")
		}
	}

	return writers
}

// appendZerologConsoleWriter is a function that appends the console writer.
func (f *Factory) appendZerologConsoleWriter(writers *[]io.Writer) {
	writer := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	*writers = append(*writers, writer)
}

// appendZerologFileWriter is a method of Factory type that appends a file writer to the slice of io.Writer pointers.
// It creates a log file with the given service name under the "logs" directory, and opens it in append mode for writing.
// If the file does not exist, it creates a new one. The file has read and write permissions for owner and read permission for others.
// The created *os.File pointer is appended to the writers slice.
// Parameters:
// - serviceName: string, the name of the service to include in the log file name.
// - writers: *[]io.Writer, a pointer to a slice of io.Writer pointers that file writer will be appended to.
// Returns: None
// TODO: handle log file rotator by app or by os
func (f *Factory) appendZerologFileWriter(serviceName string, writers *[]io.Writer) {
	file, err := os.OpenFile(
		fmt.Sprintf("%s/%s.log", "logs", serviceName),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, //nolint
		0644,
	)
	if err != nil {
		panic(err)
	}

	*writers = append(*writers, file)
}

func (f *Factory) appendZerologInfluxWriter(writers *[]io.Writer) {
	// TODO: rewrite to add new config rules here

	config := db.NewInfluxConfig()

	client := influxdb2.NewClient(config.Address(), config.AuthToken())
	api := client.WriteAPI(config.Organization(), config.Bucket())
	writer := &influxWriter{client, api}

	*writers = append(*writers, writer)
}
