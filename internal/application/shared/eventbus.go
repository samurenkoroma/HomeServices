package shared

import "samurenkoroma/services/internal/domain/shared"

type EventHandler interface {
	Handle(event shared.DomainEvent) error
}

type EventBus interface {
	Publish(events []shared.DomainEvent) error
	Register(eventName string, handler EventHandler)
}
