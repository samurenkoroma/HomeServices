package facility

import (
	"samurenkoroma/services/internal/core/domain/event"
)

type FacilityCreated struct {
	event.BaseEvent

	FacilityID string
}

func (e FacilityCreated) EventName() string {
	return "growing.facility.created"
}

func NewFacilityCreated(facilityID string) FacilityCreated {
	return FacilityCreated{
		FacilityID: facilityID,
	}
}
