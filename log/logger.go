package log

import (
	"fmt"
	"io"
	"runtime"
	"sync"
	"time"
)

type writer struct {
	writer io.Writer
	mutex  sync.Mutex
}

type Logger struct {
	Name                  string    // The logger name
	Level                 Level     // Minimum log level
	TimestampFormat       string    // The format of the logging timestamp
	IncludeTimeStamp      bool      // Whether to include the current timestamp
	ShowFileAndLineNumber bool      // Whether to show file and line number
	UseUTC                bool      // Whether the timestamp is logged as UTC
	writers               []*writer // The io.Writer instances to log to
}

func NewLogger(name string, level Level) *Logger {
	l := &Logger{
		Name:                  name,
		Level:                 level,
		TimestampFormat:       "2006-01-02T15:04:05-07:00",
		IncludeTimeStamp:      true,
		ShowFileAndLineNumber: false,
		UseUTC:                false,
	}

	loggers = append(loggers, l)
	return l
}

func (l *Logger) Output(writers ...io.Writer) {
	for i := range writers {
		for j := range l.writers {
			if writers[i] == l.writers[j].writer {
				writers = append(writers[:i], writers[i+1:]...)
			}
		}
	}

	var w []*writer
	for i := range writers {
		w = append(w, &writer{writer: writers[i]})
	}

	l.writers = append(l.writers, w...)
}

// Debug logs a message with the LevelDebug level.
func (l *Logger) Debug(message string) {
	l.write(message, 2)
}

// Debugf logs a formatted message with the LevelDebug level.
func (l *Logger) Debugf(message string, format ...interface{}) {
	l.write(fmt.Sprintf(message, format...), 2)
}

// Info logs a message with the LevelInformation level.
func (l *Logger) Info(message string) {
	l.write(message, 2)
}

// Infof logs a formatted message with the LevelInformation level.
func (l *Logger) Infof(message string, format ...interface{}) {
	l.write(fmt.Sprintf(message, format...), 2)
}

// Warning logs a message with the LevelWarning level.
func (l *Logger) Warning(message string) {
	l.write(message, 2)
}

// Warningf logs a formatted message with the LevelWarning level.
func (l *Logger) Warningf(message string, format ...interface{}) {
	l.write(fmt.Sprintf(message, format...), 2)
}

// Error logs a message with the LevelError level.
func (l *Logger) Error(message string) {
	l.write(message, 2)
}

// Errorf logs a formatted message with the LevelError level.
func (l *Logger) Errorf(message string, format ...interface{}) {
	l.write(fmt.Sprintf(message, format...), 2)
}

// Fatal logs a message with the LevelFatal level and panics.
func (l *Logger) Fatal(message string) {
	l.write(message, 2)
	panic(message)
}

// Fatalf logs a formatted message with the LevelFatal level and panics.
func (l *Logger) Fatalf(message string, format ...interface{}) {
	l.write(fmt.Sprintf(message, format...), 2)
	panic(fmt.Sprintf(message, format...))
}

// write loops all the writers on a Logger and writes the string to them.
func (l *Logger) write(message string, calldepth int) {
	prefix := fmt.Sprintf("[%s] [%s]", l.Level.String(), l.Name)

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
	for i := range l.writers {
		l.writers[i].mutex.Lock()
		l.writers[i].writer.Write(b)
		l.writers[i].mutex.Unlock()
	}
}
