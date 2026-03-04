package events

import (
	common "samurenkoroma/services/internal/common/domain"
)

type FacilityCreated struct {
	common.BaseEvent
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
