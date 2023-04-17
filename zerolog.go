package hermes

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func init() {
	// Set the error stack marshaler to pkgerrors.MarshalStack from zerolog/pkgerrors
	// This will add stack trace information to any error logs
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack //nolint
}

const (
	// Constant defining the service name field for logging
	serviceName = "service"
	// Constant defining the scope name field for logging
	scopeName = "scope"
)

// ZeroLogger is a wrapper around a zerolog Logger instance with additional service and scope fields
type ZeroLogger struct {
	*zerolog.Logger
	service string
	scope   string
}

// WithScope returns a new logger instance with the provided scope added
func (z *ZeroLogger) WithScope(scope string) Logger { //nolint
	// Copy the current logger instance and update the scope
	newLogger := *z
	newLogger.scope = scope

	return &newLogger
}

// Trace logs a trace-level message with the provided message string
func (z *ZeroLogger) Trace(message string) {
	z.appendServiceDetails(z.Logger.Trace()).Msg(message)
}

// Tracef logs a trace-level message with the provided format string and arguments
func (z *ZeroLogger) Tracef(format string, args ...any) {
	z.appendServiceDetails(z.Logger.Trace()).Msgf(format, args...)
}

// Tracev logs a trace-level message with the provided message string and key-value pairs
func (z *ZeroLogger) Tracev(message string, keyValue ...any) {
	z.logv(z.Logger.Trace(), message, keyValue...)
}

// Debug logs a debug-level message with the provided message string
func (z *ZeroLogger) Debug(message string) {
	z.appendServiceDetails(z.Logger.Debug()).Msg(message)
}

// Debugf logs a debug-level message with the provided format string and arguments
func (z *ZeroLogger) Debugf(format string, args ...any) {
	z.appendServiceDetails(z.Logger.Debug()).Msgf(format, args...)
}

// Debugv logs a debug-level message with the provided message string and key-value pairs
func (z *ZeroLogger) Debugv(message string, keyValue ...any) {
	z.logv(z.Logger.Debug(), message, keyValue...)
}

// Info logs an info-level message with the provided message string
func (z *ZeroLogger) Info(message string) {
	z.appendServiceDetails(z.Logger.Info()).Msg(message)
}

// Infof logs an info-level message with the provided format string and arguments
func (z *ZeroLogger) Infof(format string, args ...any) {
	z.appendServiceDetails(z.Logger.Info()).Msgf(format, args...)
}

// Infov logs an info-level message with the provided message string and key-value pairs
func (z *ZeroLogger) Infov(message string, keyValue ...any) {
	z.logv(z.Logger.Debug(), message, keyValue...)
}

// Warn logs a warning-level message with the provided message string
func (z *ZeroLogger) Warn(message string) {
	z.appendServiceDetails(z.Logger.Warn()).Msg(message)
}

// Warnf logs a warning-level message with the provided format string and arguments
func (z *ZeroLogger) Warnf(format string, args ...any) {
	z.appendServiceDetails(z.Logger.Warn()).Msgf(format, args...)
}

// Warnv logs a warning-level message with the provided message string and key-value pairs
func (z *ZeroLogger) Warnv(message string, keyValue ...any) {
	z.logv(z.Logger.Debug(), message, keyValue...)
}

// Err logs an error and sends the log message.
func (z *ZeroLogger) Err(err error) {
	z.appendServiceDetails(z.Logger.Err(err)).Send()
}

// Error logs an error message.
func (z *ZeroLogger) Error(message string) {
	z.appendServiceDetails(z.Logger.Error()).Msg(message)
}

// Errorf logs a formatted error message.
func (z *ZeroLogger) Errorf(format string, args ...any) {
	z.appendServiceDetails(z.Logger.Error()).Msgf(format, args...)
}

// Errorv logs an error message with additional key-value pairs.
func (z *ZeroLogger) Errorv(message string, keyValue ...any) {
	z.logv(z.Logger.Error(), message, keyValue...)
}

// Fatal logs a fatal error message.
func (z *ZeroLogger) Fatal(message string) {
	z.appendServiceDetails(z.Logger.Fatal()).Msg(message)
}

// Fatalf logs a formatted fatal error message.
func (z *ZeroLogger) Fatalf(format string, args ...any) {
	z.appendServiceDetails(z.Logger.Fatal()).Msgf(format, args...)
}

// Fatalv logs a fatal error message with additional key-value pairs.
func (z *ZeroLogger) Fatalv(message string, keyValue ...any) {
	z.logv(z.Logger.Fatal(), message, keyValue...)
}

// Panic logs a panic message.
func (z *ZeroLogger) Panic(message string) {
	z.appendServiceDetails(z.Logger.Panic()).Msg(message)
}

// Panicf logs a formatted panic message.
func (z *ZeroLogger) Panicf(format string, args ...any) {
	z.appendServiceDetails(z.Logger.Panic()).Msgf(format, args...)
}

// Panicv logs a panic message with additional key-value pairs.
func (z *ZeroLogger) Panicv(message string, keyValue ...any) {
	z.logv(z.Logger.Panic(), message, keyValue...)
}

// appendServiceDetails adds service and scope details to a given zerolog.Event instance.
func (z *ZeroLogger) appendServiceDetails(event *zerolog.Event) *zerolog.Event {
	return event.
		Str(serviceName, z.service).
		Str(scopeName, z.scope)
}

// logv logs a message with additional key-value pairs.
func (z *ZeroLogger) logv(event *zerolog.Event, message string, keyValue ...any) {
	if len(keyValue)%2 != 0 {
		panic("keyValue len must be even")
	}

	event = z.appendServiceDetails(event)

	for i := 0; i < len(keyValue); i += 2 {
		event = event.Any(keyValue[i].(string), keyValue[i+1])
	}

	event.Msg(message)
}
