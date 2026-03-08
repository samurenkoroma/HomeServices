package aggregate

import "samurenkoroma/services/internal/core/domain/event"

type Aggregate interface {
	AddEvent(e event.DomainEvent)
	PullEvents() []event.DomainEvent
}

type BaseAggregate struct {
	events []event.DomainEvent
}

func (a *BaseAggregate) AddEvent(e event.DomainEvent) {
	a.events = append(a.events, e)
}

func (a *BaseAggregate) PullEvents() []event.DomainEvent {
	ev := a.events
	a.events = nil
	return ev
}
