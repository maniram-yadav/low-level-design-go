package logger

import (
	"fmt"
	"time"
)

type ConsoleSink struct{}

func (cs *ConsoleSink) Write(message Message) error {
	_, err := fmt.Printf("[%s] %s %s: %s\n", message.Timestamp.Format(time.RFC3339), message.Level, message.Namespace, message.Content)
	return err
}
