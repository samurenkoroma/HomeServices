package domain

import "time"

type DomainEvent interface {
	EventName() string
	OccurredAt() time.Time
}

type BaseEvent struct {
	occurredAt time.Time
}

func NewBaseEvent() BaseEvent {
	return BaseEvent{
		occurredAt: time.Now(),
	}
}

func (e BaseEvent) OccurredAt() time.Time {
	return e.occurredAt
}
