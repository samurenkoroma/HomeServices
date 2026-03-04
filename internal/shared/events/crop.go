package events

import (
	common "samurenkoroma/services/internal/common/domain"
)

type CropPlanCreated struct {
	common.BaseEvent
	PlanID     string
	CropTypeID string
}
type CropPlanPublished struct {
	common.BaseEvent
	PlanID  string
	Version int
}

type CropPlanDeprecated struct {
	common.BaseEvent
	PlanID string
}

type GrowthStageAdded struct {
	common.BaseEvent
	PlanID string
	Order  int
	Name   string
}

type RotationRuleAdded struct {
	common.BaseEvent
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
		BaseEvent:  common.NewBaseEvent(),
		PlanID:     planID,
		CropTypeID: cropTypeID,
	}
}

func NewCropPlanPublished(planID string, version int) CropPlanPublished {
	return CropPlanPublished{
		BaseEvent: common.NewBaseEvent(),
		PlanID:    planID,
		Version:   version,
	}
}

func NewGrowthStageAdded(planID string, order int, name string) GrowthStageAdded {
	return GrowthStageAdded{
		BaseEvent: common.NewBaseEvent(),
		PlanID:    planID,
		Order:     order,
		Name:      name,
	}
}

func NewRotationRuleAdded(planID string, order int, name string) RotationRuleAdded {
	return RotationRuleAdded{
		BaseEvent: common.NewBaseEvent(),
		PlanID:    planID,
		Order:     order,
		Name:      name,
	}
}
