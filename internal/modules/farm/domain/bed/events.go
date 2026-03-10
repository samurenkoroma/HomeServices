package bed

import (
	"samurenkoroma/services/internal/core/domain/event"
)

type Created struct {
	event.BaseEvent

	BedID string
}

func (e Created) EventName() string {
	return "farm.bed.created"
}

func NewBedCreated(id string) Created {
	return Created{
		BedID: id,
	}
}
