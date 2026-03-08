package greenhouse

import (
	"samurenkoroma/services/internal/core/domain/event"
)

type GreenhouseCreated struct {
	event.BaseEvent
	ID string
}

func (e GreenhouseCreated) EventName() string {
	return "farm.greenhouse.created"
}

func NewGreenhouseCreated(facilityID string) GreenhouseCreated {
	return GreenhouseCreated{
		ID: facilityID,
	}
}
