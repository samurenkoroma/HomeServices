package cropplan

import (
	"samurenkoroma/services/internal/core/domain/event"
	"time"
)

type CropPlanCreatedEvent struct {
	event.BaseEvent
	PlanID       string
	AreaId       string
	Name         string
	VarietyID    string
	VarietyName  string
	SpeciesName  string
	PlantingDate time.Time
}

func (e CropPlanCreatedEvent) EventName() string { return "cropplan.created" }

// CropPlanActivatedEvent событие активации плана
type CropPlanActivatedEvent struct {
	event.BaseEvent
	PlanID      string
	ActivatedAt time.Time
}

func (e CropPlanActivatedEvent) EventName() string { return "cropplan.activated" }

// CropPlanCompletedEvent событие завершения плана
type CropPlanCompletedEvent struct {
	event.BaseEvent
	PlanID      string
	CompletedAt time.Time
	HarvestKg   float64
}

func (e CropPlanCompletedEvent) EventName() string { return "cropplan.completed" }

// CropPlanCancelledEvent событие отмены плана
type CropPlanCancelledEvent struct {
	event.BaseEvent
	PlanID      string
	Reason      string
	CancelledAt time.Time
}

func (e CropPlanCancelledEvent) EventName() string { return "cropplan.cancelled" }

// StageAddedEvent событие добавления этапа
type StageAddedEvent struct {
	event.BaseEvent
	PlanID    string
	StageID   string
	StageName string
	BBCHStart int
	BBCHEnd   int
	Order     int
}

func (e StageAddedEvent) EventName() string { return "cropplan.stage.added" }

// StageStartedEvent событие начала этапа
type StageStartedEvent struct {
	event.BaseEvent
	PlanID    string
	StageID   string
	StageName string
	BBCHCode  int
	StartedAt time.Time
}

func (e StageStartedEvent) EventName() string { return "cropplan.stage.started" }

// StageCompletedEvent событие завершения этапа
type StageCompletedEvent struct {
	event.BaseEvent
	PlanID      string
	StageID     string
	StageName   string
	CompletedAt time.Time
}

func (e StageCompletedEvent) EventName() string { return "cropplan.stage.completed" }

// StageSkippedEvent событие пропуска этапа
type StageSkippedEvent struct {
	event.BaseEvent
	PlanID    string
	StageID   string
	StageName string
	SkippedAt time.Time
}

func (e StageSkippedEvent) EventName() string { return "cropplan.stage.skipped" }
