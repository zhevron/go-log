package log

import (
	"io"
	"sync"
)

// DefaultTimestampFormat defines the timestamp format all new loggers will be
// created with. You can override this per logger instance by setting
// Logger.TimestampFormat.
var DefaultTimestampFormat = "2006-01-02T15:04:05-07:00"

// DefaultShowFileAndLineNumber defines whether the ShowFileAndLineNumber
// property on new Logger instances is enabled by default.
var DefaultShowFileAndLineNumber = false

// DefaultUseUTC defines whether the UseUTC property on new Logger instances is
// enabled by default.
var DefaultUseUTC = false

type writer struct {
	writer io.Writer
	mutex  sync.Mutex
}

var loggers []*Logger

var writers []*writer

// AddOutput adds one or more io.Writer instances as output sinks for the
// library. If a writer that has previously been added is passed, it will be
// ignored.
func AddOutput(ioWriters ...io.Writer) {
	for i := range ioWriters {
		for j := range writers {
			if ioWriters[i] == writers[j].writer {
				ioWriters = append(ioWriters[:i], ioWriters[i+1:]...)
			}
		}
	}

	var newWriters []*writer
	for i := range ioWriters {
		newWriters = append(newWriters, &writer{writer: ioWriters[i]})
	}

	writers = append(writers, newWriters...)
}
