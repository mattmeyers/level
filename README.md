# level

![Build](https://github.com/mattmeyers/level/actions/workflows/build.yml/badge.svg)

`level` is an opinionated, no-frills level logger library. It is intended to be simple and quick to use, but not guarenteed to be enough for large projects.

## Installing

This library can be installed with:

```sh
go get -u github.com/mattmeyers/level
```

## Usage

This library exposes the simple `Logger` interface that describes five basic logging levels. The interface is defined as

```go
type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
}
```

The minimum level at which a logger logs is defined by the following values. The lower the value, the more that gets logged.

```go
const (
	Debug Level = iota
	Info
	Warn
	Error
	Fatal
)
```

This library provides one implementation of the interface: the `BasicLogger`. This implementation logs messages to the provided writer. For example:

```go
// In main.go
package main

import (
	"os"
	
	"github.com/mattmeyers/level"
)

func main() {
	logger, _ := level.NewBasicLogger(level.Info, os.Stdout)
	
	// This will not be printed.
	logger.Debug("Starting...")

	value, err := doThing()
	if err != nil {
		// The message will be printed, then the process will exit.
		logger.Fatal("Something bad happened: %v", err)
	}

	// This message will get printed.
	logger.Info("Got some value: %v", value)
}

func doThing() (int, error) {
	return 1, nil
}
```

Running this will produce

```bash
$ go run main.go
2022-03-29T23:17:18-04:00 [INFO]: Got some value: 1
```