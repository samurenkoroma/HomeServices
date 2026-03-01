package eventbus

import (
	"samurenkoroma/services/internal/application/event"
	"sync"
)

type InMemoryEventBus struct {
	handlers map[string][]event.EventHandler
	mu       sync.RWMutex
}

func NewInMemoryEventBus() *InMemoryEventBus {
	return &InMemoryEventBus{
		handlers: make(map[string][]event.EventHandler),
	}
}

func (b *InMemoryEventBus) Register(eventName string, handler event.EventHandler) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.handlers[eventName] = append(b.handlers[eventName], handler)
}

func (b *InMemoryEventBus) Publish(events []event.DomainEvent) error {

	for _, e := range events {

		b.mu.RLock()
		handlers := b.handlers[e.EventName()]
		b.mu.RUnlock()

		for _, h := range handlers {
			if err := h.Handle(e); err != nil {
				return err
			}
		}
	}

	return nil
}
