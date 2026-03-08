package event

import "time"

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
