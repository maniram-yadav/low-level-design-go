package logger

import "time"

type Message struct {
	Content   string
	Level     LogLevel
	Namespace string
	Timestamp time.Time
}
