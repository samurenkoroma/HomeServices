package bed

import (
	"samurenkoroma/services/internal/core/domain/event"
)

type BedCreated struct {
	event.BaseEvent

	BedID string
}

func (e BedCreated) EventName() string {
	return "farm.bed.created"
}

func NewBedCreated(id string) BedCreated {
	return BedCreated{
		BedID: id,
	}
}
