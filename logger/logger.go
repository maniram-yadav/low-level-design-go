package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	config LoggerConfig
}

// NewLogger creates a new logger with the given configuration.
func NewLogger(config LoggerConfig) *Logger {
	return &Logger{config: config}
}
func (l *Logger) Log(content string, level LogLevel, namespace string) error {
	// Check if the message level is above the configured logging level.
	if level < l.config.LoggingLevel {
		return nil
	}

	// Enrich the message with a timestamp.
	message := Message{
		Content:   content,
		Level:     level,
		Namespace: namespace,
		Timestamp: time.Now(),
	}

	// Route the message to the appropriate sink.
	sink, exists := l.config.Sinks[level]
	if !exists {
		return fmt.Errorf("no sink configured for level %s", level)
	}

	return sink.Write(message)
}
