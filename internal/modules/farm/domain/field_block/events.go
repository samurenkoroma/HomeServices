package field_block

import (
	"samurenkoroma/services/internal/core/domain/event"
	"samurenkoroma/services/internal/core/domain/types"
	"time"
)

type Created struct {
	event.BaseEvent
	FieldID string
	BlockID string
}

func (e Created) EventName() string {
	return "farm.block.created"
}

type BlockAssignedToCrop struct {
	event.BaseEvent
	BlockID     types.FieldBlockId
	CropCycleID string
	CropID      string
	VarietyID   string
	PlantedAt   time.Time
}

func (e BlockAssignedToCrop) EventName() string { return "farm.block.assigned_to_crop" }
