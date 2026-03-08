package domain

import (
	"samurenkoroma/services/internal/core/domain/event"
)

type CropPlanCreated struct {
	event.BaseEvent
	PlanID     string
	CropTypeID string
}
type CropPlanPublished struct {
	event.BaseEvent
	PlanID  string
	Version int
}

type CropPlanDeprecated struct {
	event.BaseEvent
	PlanID string
}

type GrowthStageAdded struct {
	event.BaseEvent
	PlanID string
	Order  int
	Name   string
}

type RotationRuleAdded struct {
	event.BaseEvent
	PlanID string
	Order  int
	Name   string
}

func (e RotationRuleAdded) EventName() string  { return "crop.plan.rotation_added" }
func (e GrowthStageAdded) EventName() string   { return "crop.plan.stage_added" }
func (e CropPlanCreated) EventName() string    { return "crop.plan.created" }
func (e CropPlanPublished) EventName() string  { return "crop.plan.published" }
func (e CropPlanDeprecated) EventName() string { return "crop.plan.deprecated" }

func NewCropPlanCreated(planID string, cropTypeID string) CropPlanCreated {
	return CropPlanCreated{
		BaseEvent:  event.NewBaseEvent(),
		PlanID:     planID,
		CropTypeID: cropTypeID,
	}
}

func NewCropPlanPublished(planID string, version int) CropPlanPublished {
	return CropPlanPublished{
		BaseEvent: event.NewBaseEvent(),
		PlanID:    planID,
		Version:   version,
	}
}

func NewGrowthStageAdded(planID string, order int, name string) GrowthStageAdded {
	return GrowthStageAdded{
		BaseEvent: event.NewBaseEvent(),
		PlanID:    planID,
		Order:     order,
		Name:      name,
	}
}

func NewRotationRuleAdded(planID string, order int, name string) RotationRuleAdded {
	return RotationRuleAdded{
		BaseEvent: event.NewBaseEvent(),
		PlanID:    planID,
		Order:     order,
		Name:      name,
	}
}
