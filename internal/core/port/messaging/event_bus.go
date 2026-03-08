package messaging

import (
	"context"
	"samurenkoroma/services/internal/core/domain/event"
)

type EventHandler func(ctx context.Context, e event.DomainEvent) error

type EventBus interface {
	Register(eventName string, handler EventHandler)
	Publish(context.Context, []event.DomainEvent) error
}
