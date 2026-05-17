package events

type DomainEvent interface {
	EventName() string
	OccurredAt() time.Time
}

// BaseEvent provides common functionality for domain events
type BaseEvent struct {
	Occurred time.Time `json:"occurred_at"`
}

func (e BaseEvent) OccurredAt() time.Time {
	return e.Occurred
}
