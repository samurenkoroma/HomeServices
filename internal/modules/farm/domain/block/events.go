package block

import (
	"samurenkoroma/services/internal/core/domain/event"
)

type Created struct {
	event.BaseEvent

	BlockID string
}

func (e Created) EventName() string {
	return "farm.block.created"
}

func NewBlockCreated(id string) Created {
	return Created{
		BlockID: id,
	}
}
