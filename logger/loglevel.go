package logger

// LogLevel represents the severity level of a log message.
type LogLevel int

const (
	FATAL LogLevel = iota
	ERROR
	WARN
	INFO
	DEBUG
)

// String representation of LogLevel
func (l LogLevel) String() string {
	return [...]string{"FATAL", "ERROR", "WARN", "INFO", "DEBUG"}[l]
}
