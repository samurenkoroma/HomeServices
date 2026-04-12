package command

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/common/application/uow"
)

// CompleteCropPlanHandler команда завершения плана (сбор урожая)
type CompleteCropPlanHandler struct {
	UowFactory uow.Factory
}

type CompleteCropPlanCmd struct {
	PlanID    string  `json:"plan_id"`
	HarvestKg float64 `json:"harvest_kg"`
}

func DecodeCompleteCropPlan(data []byte) (any, error) {
	var cmd CompleteCropPlanCmd
	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, err
	}
	if cmd.PlanID == "" {
		return nil, errors.New("plan_id is required")
	}
	if cmd.HarvestKg < 0 {
		return nil, errors.New("harvest_kg must be non-negative")
	}
	return cmd, nil
}

func (h *CompleteCropPlanHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(CompleteCropPlanCmd)
	if !ok {
		return errors.New("invalid command type")
	}

	uowObj, err := h.UowFactory.Begin(ctx)
	if err != nil {
		return err
	}
	defer uowObj.Rollback()

	planRepo := uowObj.CropPlans()
	plan, err := planRepo.FindByID(ctx, c.PlanID)
	if err != nil {
		return err
	}

	if err := plan.Complete(c.HarvestKg); err != nil {
		return err
	}

	if err := planRepo.Update(ctx, plan); err != nil {
		return err
	}

	uowObj.RegisterAggregate(plan)
	return uowObj.Commit()
}
