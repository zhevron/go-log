package log

import (
	"io"
	"os"
	"sync"
)

// Sink defines a log sink with a target io.Writer.
type Sink struct {
	mutex  sync.Mutex // Ensure atomic writes
	Level  Level      // The minimum log level before using the sink
	Writer io.Writer  // The io.Writer to use for writing messages
}

// Default sinks.
var (
	Stdout = &Sink{Level: LevelWarning, Writer: os.Stdout}
)

// NewSink returns a new Sink.
func NewSink(level Level, writer io.Writer) *Sink {
	return &Sink{Level: level, Writer: writer}
}

// Write writes a message to the log sink io.Writer if the level is equal to
// or higher than the minumum level set on the sink.
func (s *Sink) Write(level Level, message string) error {
	if level < s.Level {
		return nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, err := s.Writer.Write([]byte(message))
	return err
}
