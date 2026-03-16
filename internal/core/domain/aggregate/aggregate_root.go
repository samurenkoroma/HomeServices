package aggregate

import (
	"samurenkoroma/services/internal/core/domain/event"
	"time"
)

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

type Entity[T any] struct {
	BaseAggregate
	Id        T
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewEntity[T any](id T) Entity[T] {
	now := time.Now()
	return Entity[T]{
		BaseAggregate: BaseAggregate{
			events: make([]event.DomainEvent, 0),
		},
		Id:        id,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	}
}
