package cropplan

import (
	"samurenkoroma/services/internal/core/domain/event"
	"time"
)

// CropPlanCreatedEvent событие создания плана
type CropPlanCreatedEvent struct {
	PlanID       string
	BedID        string
	Name         string
	VarietyID    string
	VarietyName  string
	CropName     string
	PlantingDate time.Time
	OccurredAt   time.Time
}

func (e CropPlanCreatedEvent) EventName() string     { return "cropplan.created" }
func (e CropPlanCreatedEvent) OccurredAt() time.Time { return e.OccurredAt }

// CropPlanActivatedEvent событие активации плана
type CropPlanActivatedEvent struct {
	PlanID      string
	ActivatedAt time.Time
	OccurredAt  time.Time
}

func (e CropPlanActivatedEvent) EventName() string     { return "cropplan.activated" }
func (e CropPlanActivatedEvent) OccurredAt() time.Time { return e.OccurredAt }

// CropPlanCompletedEvent событие завершения плана
type CropPlanCompletedEvent struct {
	PlanID      string
	CompletedAt time.Time
	HarvestKg   float64
	OccurredAt  time.Time
}

func (e CropPlanCompletedEvent) EventName() string     { return "cropplan.completed" }
func (e CropPlanCompletedEvent) OccurredAt() time.Time { return e.OccurredAt }

// CropPlanCancelledEvent событие отмены плана
type CropPlanCancelledEvent struct {
	PlanID      string
	Reason      string
	CancelledAt time.Time
	OccurredAt  time.Time
}

func (e CropPlanCancelledEvent) EventName() string     { return "cropplan.cancelled" }
func (e CropPlanCancelledEvent) OccurredAt() time.Time { return e.OccurredAt }

// StageAddedEvent событие добавления этапа
type StageAddedEvent struct {
	PlanID     string
	StageID    string
	StageName  string
	BBCHStart  int
	BBCHEnd    int
	Order      int
	OccurredAt time.Time
}

func (e StageAddedEvent) EventName() string     { return "cropplan.stage.added" }
func (e StageAddedEvent) OccurredAt() time.Time { return e.OccurredAt }

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
