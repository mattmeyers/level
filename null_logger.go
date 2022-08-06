package level

import "os"

var _ Logger = (*NullLogger)(nil)

// NullLogger is a logger that does nothing. All log statements, regardless of level, will
// be ignored. Note that calling Fatal will still cause the process to exit.
type NullLogger struct{}

func NewNullLogger() NullLogger {
	return NullLogger{}
}

// Logs at the Debug level.
func (NullLogger) Debug(format string, args ...interface{}) {}

// Logs at the Info level.
func (NullLogger) Info(format string, args ...interface{}) {}

// Logs at the Warn level.
func (NullLogger) Warn(format string, args ...interface{}) {}

// Logs at the Error level.
func (NullLogger) Error(format string, args ...interface{}) {}

// Logs at the Fatal level then calls os.Exit(1).
func (NullLogger) Fatal(format string, args ...interface{}) {
	os.Exit(1)
}
