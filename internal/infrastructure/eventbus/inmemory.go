package eventbus

import (
	"context"
	"samurenkoroma/services/internal/common/domain"
	"sync"
)

type InMemoryEventBus struct {
	handlers map[string][]domain.EventHandler
	mu       sync.RWMutex
}

func NewInMemoryEventBus() *InMemoryEventBus {
	return &InMemoryEventBus{
		handlers: make(map[string][]domain.EventHandler),
	}
}

func (b *InMemoryEventBus) Register(eventName string, handler domain.EventHandler) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.handlers[eventName] = append(b.handlers[eventName], handler)
}

func (b *InMemoryEventBus) Publish(
	ctx context.Context,
	events []domain.DomainEvent,
) error {

	for _, event := range events {

		b.mu.RLock()
		handlers := b.handlers[event.EventName()]
		b.mu.RUnlock()

		for _, h := range handlers {
			if err := h(ctx, event); err != nil {
				return err
			}
		}
	}

	return nil
}
