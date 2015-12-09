package log

type level uint8

const (
	// Debug is used for messages related to debugging.
	Debug level = iota
	// Info is used for informational messages.
	Info
	// Warning is used for warnings.
	Warning
	// Error is used for non-fatal errors.
	Error
	// Fatal used for fatal errors.
	Fatal
)

func (l level) String() string {
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

	default:
		return "UNKNOWN"
	}
}
