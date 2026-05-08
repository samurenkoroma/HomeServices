package cropplan

import (
	"samurenkoroma/services/internal/core/domain/event"
)

type CropPlanActivated struct {
	event.BaseEvent
	PlanID string
}

func NewCropPlanActivated(planID string) *CropPlanActivated {
	return &CropPlanActivated{
		BaseEvent: event.NewBaseEvent(),
		PlanID:    planID,
	}
}
func (e CropPlanActivated) EventName() string { return "crop-plan.plan.activated" }
