// Package log provides a simple logging interface.
package log

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// Level identifies a log level.
type Level int

// Log levels.
const (
	LevelNone        Level = 0 // Do not send log messages to sinks.
	LevelDebug       Level = 1 // Use for messages related to debugging.
	LevelInformation Level = 2 // Use for informational messages.
	LevelWarning     Level = 3 // Use for warnings.
	LevelError       Level = 4 // Use for non-fatal errors.
	LevelFatal       Level = 5 // Use for fatal errors.
)

// TimestampFormat decided the format of the timestamps printed.
var TimestampFormat = "2006-01-02T15:04:05-07:00"

// IncludeTimestamp decides whether the current timestamp should be included in
// each log message.
var IncludeTimeStamp = true

// ShowFileAndLineNumber decides whether the logger should look up the call
// stack and try to guess which file and line invoked the call to Log.
var ShowFileAndLineNumber = false

// UseUTC decides whether the timestamp is logged in the local timezone or UTC.
var UseUTC = false

// sinks contain all the defined log sinks.
var sinks []*Sink

// AddSink adds a log sink to the pool of log message receivers.
func AddSink(sink *Sink) {
	for _, s := range sinks {
		if s == sink {
			return
		}
	}

	sinks = append(sinks, sink)
}

// Log logs a message to the defined sinks.
func Log(level Level, message string) {
	prefix := fmt.Sprintf("[%s]", level.String())

	if IncludeTimeStamp {
		t := time.Now()

		if UseUTC {
			t = t.UTC()
		}

		prefix = fmt.Sprintf("%s [%s]", prefix, t.Format(TimestampFormat))
	}

	if ShowFileAndLineNumber {
		i := 1
		_, file, line, ok := runtime.Caller(0)
		for ok && strings.Contains(file, "log.go") {
			_, file, line, ok = runtime.Caller(i)
			i++
		}

		prefix = fmt.Sprintf("%s [%s:%d]", prefix, file, line)
	}

	message = fmt.Sprintf("%s %s", prefix, message)
	if message[len(message)-1:] != "\n" {
		message = message + "\n"
	}

	for _, sink := range sinks {
		sink.Write(level, message)
	}
}

// Debug logs a message with the LevelDebug level.
func Debug(message string) {
	Log(LevelDebug, message)
}

// Debugf logs a formatted message with the LevelDebug level.
func Debugf(message string, format ...interface{}) {
	Log(LevelDebug, fmt.Sprintf(message, format...))
}

// Info logs a message with the LevelInformation level.
func Info(message string) {
	Log(LevelInformation, message)
}

// Infof logs a formatted message with the LevelInformation level.
func Infof(message string, format ...interface{}) {
	Log(LevelInformation, fmt.Sprintf(message, format...))
}

// Warning logs a message with the LevelWarning level.
func Warning(message string) {
	Log(LevelWarning, message)
}

// Warningf logs a formatted message with the LevelWarning level.
func Warningf(message string, format ...interface{}) {
	Log(LevelWarning, fmt.Sprintf(message, format...))
}

// Error logs a message with the LevelError level.
func Error(message string) {
	Log(LevelError, message)
}

// Errorf logs a formatted message with the LevelError level.
func Errorf(message string, format ...interface{}) {
	Log(LevelError, fmt.Sprintf(message, format...))
}

// Fatal logs a message with the LevelFatal level and panics.
func Fatal(message string) {
	Log(LevelFatal, message)
	panic(message)
}

// Fatalf logs a formatted message with the LevelFatal level and panics.
func Fatalf(message string, format ...interface{}) {
	Log(LevelFatal, fmt.Sprintf(message, format...))
	panic(fmt.Sprintf(message, format...))
}

// String returns the string version of a log level.
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"

	case LevelInformation:
		return "INFO"

	case LevelWarning:
		return "WARNING"

	case LevelError:
		return "ERROR"

	case LevelFatal:
		return "FATAL"
	}

	return ""
}
