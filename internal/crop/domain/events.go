package domain

import "time"

type StageAddedEvent struct {
	PlanID string
	Stage  string
	Order  int
	Time   time.Time
}

func (e StageAddedEvent) EventName() string {
	return "cropplan.stage.added"
}
func (e StageAddedEvent) OccurredAt() time.Time {
	return e.Time
}

type StageStartedEvent struct {
	PlanID    string
	StageID   string
	StartedAt time.Time
	Time      time.Time
}

func (e StageStartedEvent) EventName() string {
	return "cropplan.stage.started	"
}
func (e StageStartedEvent) OccurredAt() time.Time {
	return e.Time
}

type StageCompletedEvent struct {
	PlanID      string
	StageID     string
	CompletedAt time.Time
	Time        time.Time
}

func (e StageCompletedEvent) EventName() string     { return "cropplan.stage.completed" }
func (e StageCompletedEvent) OccurredAt() time.Time { return e.Time }

type StageSkippedEvent struct {
	PlanID    string
	StageID   string
	SkippedAt time.Time
	Time      time.Time
}

func (e StageSkippedEvent) EventName() string     { return "cropplan.stage.skipped" }
func (e StageSkippedEvent) OccurredAt() time.Time { return e.Time }

type CropPlanCreatedEvent struct {
	PlanID   string
	AreaID   string
	CropName string
	Time     time.Time
}

func (e CropPlanCreatedEvent) EventName() string     { return "cropplan.stage.created" }
func (e CropPlanCreatedEvent) OccurredAt() time.Time { return e.Time }

type CropPlanActivatedEvent struct {
	PlanID      string
	ActivatedAt time.Time
}

func (e CropPlanActivatedEvent) EventName() string     { return "cropplan.stage.activated" }
func (e CropPlanActivatedEvent) OccurredAt() time.Time { return e.ActivatedAt }

type CropPlanCompletedEvent struct {
	PlanID      string
	HarvestKg   float64
	CompletedAt time.Time
}

func (e CropPlanCompletedEvent) EventName() string     { return "cropplan.stage.completed" }
func (e CropPlanCompletedEvent) OccurredAt() time.Time { return e.CompletedAt }

type CropPlanCancelledEvent struct {
	PlanID      string
	CancelledAt time.Time
}

func (e CropPlanCancelledEvent) EventName() string     { return "cropplan.stage.cancelled" }
func (e CropPlanCancelledEvent) OccurredAt() time.Time { return e.CancelledAt }
