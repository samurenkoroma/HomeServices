package field

import (
	"samurenkoroma/services/internal/core/domain/event"
)

type FieldCreated struct {
	event.BaseEvent

	FieldID string
}

func (e FieldCreated) EventName() string {
	return "farm.field.created"
}

func NewFieldCreated(id string) FieldCreated {
	return FieldCreated{
		FieldID: id,
	}
}
