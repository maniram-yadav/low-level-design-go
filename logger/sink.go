package logger

type Sink interface {
	Write(message Message) error
}
