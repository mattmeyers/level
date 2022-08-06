package level

import (
	"fmt"
	"io"
	"os"
	"time"
)

var _ Logger = (*BasicLogger)(nil)

// BasicLogger implements the Logger interface using the defined Level constants. The provided
// level is treated as the minimum. Any messages passed to a level that is at least the defined
// level will be printed.
//
// Every log message is treated as a single line. If there is no newline at the end of the
// message, then one will be added.
type BasicLogger struct {
	w     io.Writer
	level Level
}

// NewBasicLogger constructs a new logger. An error will be returned if an invalid level is
// provided. If no output writer is provided, then os.Stdout will be used.
func NewBasicLogger(level Level, out io.Writer) (*BasicLogger, error) {
	if err := level.validate(); err != nil {
		return nil, err
	}

	if out == nil {
		out = os.Stdout
	}

	return &BasicLogger{
		w:     out,
		level: level,
	}, nil
}

// Logs at the Debug level.
func (l *BasicLogger) Debug(format string, args ...interface{}) {
	if l.level <= Debug {
		l.printPrefixTag(Debug)
		l.printMessage([]byte(fmt.Sprintf(format, args...)))
	}
}

// Logs at the Info level.
func (l *BasicLogger) Info(format string, args ...interface{}) {
	if l.level <= Info {
		l.printPrefixTag(Info)
		l.printMessage([]byte(fmt.Sprintf(format, args...)))
	}
}

// Logs at the Warn level.
func (l *BasicLogger) Warn(format string, args ...interface{}) {
	if l.level <= Warn {
		l.printPrefixTag(Warn)
		l.printMessage([]byte(fmt.Sprintf(format, args...)))
	}
}

// Logs at the Error level.
func (l *BasicLogger) Error(format string, args ...interface{}) {
	if l.level <= Error {
		l.printPrefixTag(Error)
		l.printMessage([]byte(fmt.Sprintf(format, args...)))
	}
}

// Logs at the Fatal level then calls os.Exit(1).
func (l *BasicLogger) Fatal(format string, args ...interface{}) {
	if l.level <= Fatal {
		l.printPrefixTag(Fatal)
		l.printMessage([]byte(fmt.Sprintf(format, args...)))
		os.Exit(1)
	}
}

func (l *BasicLogger) printPrefixTag(level Level) {
	l.w.Write([]byte(fmt.Sprintf("%s [%s]: ", time.Now().Format(time.RFC3339), level)))
}

var newline = []byte{'\n'}

func (l *BasicLogger) printMessage(message []byte) {
	l.w.Write(message)
	if len(message) == 0 || message[len(message)-1] != '\n' {
		l.w.Write(newline)
	}
}
