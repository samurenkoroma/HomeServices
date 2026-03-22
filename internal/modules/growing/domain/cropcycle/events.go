package cropcycle

import (
	"samurenkoroma/services/internal/core/domain/event"
	"time"
)

// CropCycleStarted — цикл начат
type CropCycleStarted struct {
	event.BaseEvent
	CycleID   string    `json:"cycle_id"`
	AreaID    string    `json:"area_id"`
	SeasonID  string    `json:"season_id"`
	StartedAt time.Time `json:"started_at"`
}

func (e CropCycleStarted) EventName() string {
	return "growing.cycle.started"
}

// OperationRecorded — операция записана
type OperationRecorded struct {
	event.BaseEvent
	CycleID   string    `json:"cycle_id"`
	Operation Operation `json:"operation"`
}

func (e OperationRecorded) EventName() string {
	return "growing.cycle.operation_recorded"
}

// CropCycleHarvested — урожай собран
type CropCycleHarvested struct {
	event.BaseEvent
	CycleID string `json:"cycle_id"`
	Yield   Yield  `json:"yield"`
}

func (e CropCycleHarvested) EventName() string {
	return "growing.cycle.harvested"
}

// CropCycleCompleted — цикл завершён
type CropCycleCompleted struct {
	event.BaseEvent
	CycleID string `json:"cycle_id"`
}

func (e CropCycleCompleted) EventName() string {
	return "growing.cycle.completed"
}

// CropCycleFailed — цикл не удался
type CropCycleFailed struct {
	event.BaseEvent
	CycleID string `json:"cycle_id"`
	Reason  string `json:"reason"`
}

func (e CropCycleFailed) EventName() string {
	return "growing.cycle.failed"
}

// CropCycleCancelled — цикл отменён
type CropCycleCancelled struct {
	event.BaseEvent
	CycleID string `json:"cycle_id"`
	Reason  string `json:"reason"`
}

func (e CropCycleCancelled) EventName() string {
	return "growing.cycle.cancelled"
}
