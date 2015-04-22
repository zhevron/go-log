// Package log provides a simple logging interface.
package log

// Level identifies a log level.
type Level int

// Log levels.
const (
	None    Level = 0 // Do not send log messages to sinks.
	Debug   Level = 1 // Use for messages related to debugging.
	Info    Level = 2 // Use for informational messages.
	Warning Level = 3 // Use for warnings.
	Error   Level = 4 // Use for non-fatal errors.
	Fatal   Level = 5 // Use for fatal errors.
)

// loggers contain all the created Logger objects.
var loggers []*Logger

// GetLogger will return the logger with the given name.
// If the logger does not exist, one will be created a default level of Info.
func GetLogger(name string) *Logger {
	for i := range loggers {
		if loggers[i].Name == name {
			return loggers[i]
		}
	}

	return NewLogger(name, Info)
}

// String returns the string version of a log level.
func (l Level) String() string {
	switch l {
	case Debug:
		return "DEBUG"

	case Info:
		return "INFO"

	case Warning:
		return "WARNING"

	case Error:
		return "ERROR"

	case Fatal:
		return "FATAL"
	}

	return ""
}
