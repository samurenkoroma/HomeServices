package field

import (
	"samurenkoroma/services/internal/core/domain/event"
	"samurenkoroma/services/internal/core/domain/types"
)

type Created struct {
	event.BaseEvent

	ID string
}

func (e Created) EventName() string {
	return "farm.field.created"
}

func NewFieldCreated(id string) Created {
	return Created{
		BaseEvent: event.NewBaseEvent(),
		ID:        id,
	}
}

type BlockAddedToField struct {
	event.BaseEvent
	FieldID     types.FieldId
	BlockID     types.FieldBlockId
	BlockNumber int
	Area        types.Area
}

func (b BlockAddedToField) EventName() string {
	return "farm.field.block_added"
}
