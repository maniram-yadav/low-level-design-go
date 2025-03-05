package logger

type LoggerConfig struct {
	TimeFormat   string
	LoggingLevel LogLevel
	Sinks        map[LogLevel]Sink
}
