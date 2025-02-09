package cqrs

import (
	"fmt"
	"sync"
)

type EventBus struct {
	handlers map[string][]func(interface{})
	mu       sync.Mutex
}

func NewEventBus() *EventBus {
	return &EventBus{handlers: make(map[string][]func(interface{})), mu: sync.Mutex{}}
}

func (eb *EventBus) Subscribe(eventName string, handler func(interface{})) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	fmt.Printf("\nInside Subscribe method  Event name :%s", eventName)
	if _, ok := eb.handlers[eventName]; !ok {
		eb.handlers[eventName] = make([]func(interface{}), 0)
	}
	eb.handlers[eventName] = append(eb.handlers[eventName], handler)
}

func (eb *EventBus) Publish(event interface{}) {

	eb.mu.Lock()
	defer eb.mu.Unlock()

	eventName := fmt.Sprintf("%T", event)
	fmt.Printf("\nInside Publish method  Event name :%s", eventName)
	if handlers, ok := eb.handlers[eventName]; ok {
		for _, handler := range handlers {
			go handler(event)
		}
	}

}
