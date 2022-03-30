package level

import (
	"errors"
	"fmt"
	"strings"
)

// Logger represents a standard level logging interface. Every method logs the provided
// message using fmt.Printf with a timestamp and level prefix. Fatal logs the message
// like the other methods, but calls os.Exit(1) afterwards.
type Logger interface {
	// Logs at the Debug level.
	Debug(format string, args ...interface{})
	// Logs at the Info level.
	Info(format string, args ...interface{})
	// Logs at the Warn level.
	Warn(format string, args ...interface{})
	// Logs at the Error level.
	Error(format string, args ...interface{})
	// Logs at the Fatal level then calls os.Exit(1).
	Fatal(format string, args ...interface{})
}

// Level represents a logging level. This restricts the logger to print only messages with
// at least this level.
type Level int

// The available logging levels.
const (
	Debug Level = iota
	Info
	Warn
	Error
	Fatal
)

func (l Level) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	}

	return fmt.Sprintf("%%!(Level=%d)", l)
}

// ParseLevel converts a string to the corresponding Level. Comparisons are case insensitive.
// If an unknown level is provided, then an error will be returned.
func ParseLevel(l string) (Level, error) {
	switch strings.ToLower(l) {
	case "debug":
		return Debug, nil
	case "info":
		return Info, nil
	case "warn":
		return Warn, nil
	case "error":
		return Error, nil
	case "fatal":
		return Fatal, nil
	}

	return Level(-1), errors.New("invalid log level")
}

func (l Level) validate() error {
	if strings.HasPrefix(l.String(), "%!") {
		return errors.New("invalid Level")
	}

	return nil
}
