package eventbus

import (
	"context"
	"samurenkoroma/services/internal/core/domain/event"
	"samurenkoroma/services/internal/core/port/messaging"
	"sync"
)

type InMemoryEventBus struct {
	handlers map[string][]messaging.EventHandler
	mu       sync.RWMutex
}

func NewInMemoryEventBus() *InMemoryEventBus {
	return &InMemoryEventBus{
		handlers: make(map[string][]messaging.EventHandler),
	}
}

func (b *InMemoryEventBus) Register(eventName string, handler messaging.EventHandler) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.handlers[eventName] = append(b.handlers[eventName], handler)
}

func (b *InMemoryEventBus) Publish(
	ctx context.Context,
	events []event.DomainEvent,
) error {

	for _, e := range events {

		b.mu.RLock()
		handlers := b.handlers[e.EventName()]
		b.mu.RUnlock()

		for _, h := range handlers {
			if err := h(ctx, e); err != nil {
				return err
			}
		}
	}

	return nil
}
