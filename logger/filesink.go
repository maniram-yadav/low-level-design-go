package logger

import (
	"fmt"
	"os"
	"time"
)

type FileSink struct {
	file *os.File
}

func NewFileSink(filePath string) (*FileSink, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FileSink{file: file}, nil
}

func (fs *FileSink) Write(message Message) error {
	_, err := fmt.Fprintf(fs.file, "[%s] %s %s: %s\n", message.Timestamp.Format(time.RFC3339), message.Level, message.Namespace, message.Content)
	return err
}
