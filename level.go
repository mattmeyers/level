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
	// Logs at the LevelDebug level.
	Debug(format string, args ...interface{})
	// Logs at the LevelInfo level.
	Info(format string, args ...interface{})
	// Logs at the LevelWarn level.
	Warn(format string, args ...interface{})
	// Logs at the LevelError level.
	Error(format string, args ...interface{})
	// Logs at the LevelFatal level then calls os.Exit(1).
	Fatal(format string, args ...interface{})
}

// Level represents a logging level. This restricts the logger to print only messages with
// at least this level.
type Level int

// The available logging levels.
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	}

	return fmt.Sprintf("%%!(Level=%d)", l)
}

// ParseLevel converts a string to the corresponding Level. Comparisons are case insensitive.
// If an unknown level is provided, then an error will be returned.
func ParseLevel(l string) (Level, error) {
	switch strings.ToLower(l) {
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "warn":
		return LevelWarn, nil
	case "error":
		return LevelError, nil
	case "fatal":
		return LevelFatal, nil
	}

	return Level(-1), errors.New("invalid log level")
}

func (l Level) validate() error {
	if strings.HasPrefix(l.String(), "%!") {
		return errors.New("invalid Level")
	}

	return nil
}
