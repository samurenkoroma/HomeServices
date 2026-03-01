package domain

import "context"

type EventHandler func(ctx context.Context, event DomainEvent) error

type EventBus interface {
	Register(eventName string, handler EventHandler)
	Publish(context.Context, []DomainEvent) error
}
