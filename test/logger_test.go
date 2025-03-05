package test

import (
	"fmt"
	"lld/logger"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {

	fileSink, err := logger.NewFileSink("E:\\golang\\application.log")
	if err != nil {
		fmt.Println("Failed to create file sink:", err)
		return
	}

	// Create a console sink.
	consoleSink := &logger.ConsoleSink{}

	// Configure the logger.
	config := logger.LoggerConfig{
		TimeFormat:   time.RFC3339,
		LoggingLevel: logger.INFO,
		Sinks: map[logger.LogLevel]logger.Sink{
			logger.INFO:  consoleSink,
			logger.ERROR: fileSink,
			logger.FATAL: fileSink,
		},
	}

	log := logger.NewLogger(config)

	log.Log("This is an info message", logger.INFO, "main")
	log.Log("This is an error message", logger.ERROR, "main")
	log.Log("This is a debug message", logger.DEBUG, "main")
}
