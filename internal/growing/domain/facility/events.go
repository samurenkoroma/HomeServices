package facility

import "time"

type FacilityCreatedEvent struct {
	FacilityID string
	Time       time.Time
}

func (e FacilityCreatedEvent) EventName() string {
	return "FacilityCreated"
}

func (e FacilityCreatedEvent) OccurredAt() time.Time {
	return e.Time
}
