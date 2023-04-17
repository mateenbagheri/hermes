package hermes

// Logger interface defines the logging methods that are common across different loggers
type Logger interface {
	TraceLogger
	DebugLogger
	InfoLogger
	WarnLogger
	ErrorLogger
	FatalLogger
	PanicLogger

	// WithScope returns a new Logger instance with a specified scope added to the context
	WithScope(scope string) Logger
}

// Type is a custom type to represent different logger types
type Type int8

const (
	// ZeroLoggerType represents the Zero logger implementation
	ZeroLoggerType Type = iota
)

// Level represents different logging levels that can be used for filtering logs
type Level int8

const (
	// DebugLevel is the logging level for debug logs
	DebugLevel Level = iota
	// InfoLevel is the logging level for informational logs
	InfoLevel
	// WarnLevel is the logging level for warning logs
	WarnLevel
	// ErrorLevel is the logging level for error logs
	ErrorLevel
	// FatalLevel is the logging level for fatal logs
	FatalLevel
	// PanicLevel is the logging level for panic logs
	PanicLevel

	// TraceLevel is a special logging level to log trace level logs
	TraceLevel Level = -1
)

// TraceLogger defines logging methods for trace level logs
type TraceLogger interface {
	Trace(message string)
	Tracef(format string, args ...any)

	// Tracev logs trace level logs with additional context data in the form of key-value pairs
	Tracev(message string, keyValue ...any)
}

// DebugLogger defines logging methods for debug level logs
type DebugLogger interface {
	Debug(message string)
	Debugf(format string, args ...any)

	// Debugv logs debug level logs with additional context data in the form of key-value pairs
	Debugv(message string, KeyValue ...any)
}

// InfoLogger defines logging methods for info level logs
type InfoLogger interface {
	Info(message string)
	Infof(format string, args ...any)

	// Infov logs info level logs with additional context data in the form of key-value pairs
	Infov(message string, keyValue ...any)
}

// WarnLogger defines logging methods for warn level logs
type WarnLogger interface {
	Warn(message string)
	Warnf(format string, args ...any)

	// Warnv logs warn level logs with additional context data in the form of key-value pairs
	Warnv(message string, keyValue ...any)
}

// ErrorLogger defines logging methods for error level logs
type ErrorLogger interface {
	Err(err error)
	Error(message string)
	Errorf(format string, args ...any)

	// Errorv logs error level logs with additional context data in the form of key-value pairs
	Errorv(message string, keyValue ...any)
}

// FatalLogger defines logging methods for fatal level logs
type FatalLogger interface {
	Fatal(message string)
	Fatalf(format string, args ...any)

	// Fatalv logs fatal level logs with additional context data in the form of key-value pairs
	Fatalv(message string, keyValue ...any)
}

// PanicLogger defines logging methods for panic level logs
type PanicLogger interface {
	Panic(message string)
	Panicf(format string, args ...any)

	// Panicv logs panic level logs with additional context data in the form of key-value pairs
	Panicv(message string, keyValue ...any)
}
