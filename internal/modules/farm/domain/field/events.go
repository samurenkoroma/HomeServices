package field

import (
	"samurenkoroma/services/internal/core/domain/event"
)

type Created struct {
	event.BaseEvent

	FieldID string
}

func (e Created) EventName() string {
	return "farm.field.created"
}

func NewFieldCreated(id string) Created {
	return Created{
		FieldID: id,
	}
}
