package greenhouse

import (
	"samurenkoroma/services/internal/core/domain/event"
)

type Created struct {
	event.BaseEvent
	ID string
}

func (e Created) EventName() string {
	return "farm.greenhouse.created"
}

func NewGreenhouseCreated(facilityID string) Created {
	return Created{
		ID: facilityID,
	}
}
