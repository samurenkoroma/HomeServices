package cropplan

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/common/application/uow"
	"samurenkoroma/services/internal/growing/domain/cropplan/cropplan"

	"github.com/google/uuid"
)

type CreateCropPlanHandler struct {
	UowFactory uow.Factory
}

type CreateCropPlanCmd struct {
	AreaID   string `json:"area_id"`
	Name     string `json:"name"`
	CropName string `json:"crop_name"`
}

func DecodeCreateCropPlan(data []byte) (any, error) {
	var cmd CreateCropPlanCmd
	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, fmt.Errorf("failed to decode CreateCropPlan command: %w", err)
	}

	if cmd.AreaID == "" {
		return nil, errors.New("bed_id is required")
	}
	if cmd.Name == "" {
		return nil, errors.New("name is required")
	}
	if cmd.CropName == "" {
		return nil, errors.New("crop_name is required")
	}

	return cmd, nil
}

func (h *CreateCropPlanHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(CreateCropPlanCmd)
	if !ok {
		return errors.New("invalid command type")
	}

	uowObj, err := h.UowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin unit of work: %w", err)
	}
	defer uowObj.Rollback()

	// Создаем план
	plan := cropplan.NewCropPlan(
		cropplan.PlanID(uuid.New().String()),
		cropplan.GrowingAreaID(c.AreaID),
		c.Name,
		c.CropName,
	)

	// Сохраняем
	repo := uowObj.CropPlans()
	if err := repo.Save(plan); err != nil {
		return fmt.Errorf("failed to save crop plan: %w", err)
	}

	uowObj.RegisterAggregate(plan)

	if err := uowObj.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
