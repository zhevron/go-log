package log

import (
	"fmt"
	"runtime"
	"time"
)

// Logger defines a named logging instance. It is used to indicate the origin
// of a logged line in the output.
type Logger struct {
	Name                  string // The logger name
	MinimumLevel          level  // Minimum log level
	TimestampFormat       string // The format of the logging timestamp
	IncludeTimeStamp      bool   // Whether to include the current timestamp
	ShowFileAndLineNumber bool   // Whether to show file and line number
	UseUTC                bool   // Whether the timestamp is logged as UTC
}

// GetLogger will return the Logger instance with the given name.
// If the Logger does not exist, one will be created a minimum level of Info.
func GetLogger(name string) *Logger {
	for i := range loggers {
		if loggers[i].Name == name {
			return loggers[i]
		}
	}

	return NewLogger(name, Info)
}

// NewLogger creates a new Logger instance with the given name and minimum
// log level. If there's already an instance with the given name,
func NewLogger(name string, minimumLevel level) *Logger {
	l := &Logger{
		Name:                  name,
		MinimumLevel:          minimumLevel,
		TimestampFormat:       DefaultTimestampFormat,
		IncludeTimeStamp:      true,
		ShowFileAndLineNumber: DefaultShowFileAndLineNumber,
		UseUTC:                DefaultUseUTC,
	}

	loggers = append(loggers, l)
	return l
}

// Debug logs a message with the LevelDebug level.
func (l *Logger) Debug(message string) {
	l.write(message, Debug, 2)
}

// Debugf logs a formatted message with the LevelDebug level.
func (l *Logger) Debugf(message string, format ...interface{}) {
	l.write(fmt.Sprintf(message, format...), Debug, 2)
}

// Info logs a message with the LevelInformation level.
func (l *Logger) Info(message string) {
	l.write(message, Info, 2)
}

// Infof logs a formatted message with the LevelInformation level.
func (l *Logger) Infof(message string, format ...interface{}) {
	l.write(fmt.Sprintf(message, format...), Info, 2)
}

// Warning logs a message with the LevelWarning level.
func (l *Logger) Warning(message string) {
	l.write(message, Warning, 2)
}

// Warningf logs a formatted message with the LevelWarning level.
func (l *Logger) Warningf(message string, format ...interface{}) {
	l.write(fmt.Sprintf(message, format...), Warning, 2)
}

// Error logs a message with the LevelError level.
func (l *Logger) Error(message string) {
	l.write(message, Error, 2)
}

// Errorf logs a formatted message with the LevelError level.
func (l *Logger) Errorf(message string, format ...interface{}) {
	l.write(fmt.Sprintf(message, format...), Error, 2)
}

// Fatal logs a message with the LevelFatal level and panics.
func (l *Logger) Fatal(message string) {
	l.write(message, Fatal, 2)
	panic(message)
}

// Fatalf logs a formatted message with the LevelFatal level and panics.
func (l *Logger) Fatalf(message string, format ...interface{}) {
	l.write(fmt.Sprintf(message, format...), Fatal, 2)
	panic(fmt.Sprintf(message, format...))
}

func (l *Logger) write(message string, logLevel level, calldepth int) {
	if logLevel < l.MinimumLevel {
		return
	}

	prefix := fmt.Sprintf("[%s] [%s]", logLevel.String(), l.Name)

	if l.IncludeTimeStamp {
		t := time.Now()

		if l.UseUTC {
			t = t.UTC()
		}

		prefix = fmt.Sprintf("%s [%s]", prefix, t.Format(l.TimestampFormat))
	}

	if l.ShowFileAndLineNumber {
		_, file, line, ok := runtime.Caller(calldepth)
		if !ok {
			file = "<unknown>"
			line = 0
		}

		prefix = fmt.Sprintf("%s [%s:%d]", prefix, file, line)
	}

	message = fmt.Sprintf("%s %s", prefix, message)
	if message[len(message)-1:] != "\n" {
		message = message + "\n"
	}

	b := []byte(message)
	for i := range writers {
		writers[i].mutex.Lock()
		writers[i].writer.Write(b)
		writers[i].mutex.Unlock()
	}
}
