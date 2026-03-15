package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/core/domain/repository"
)

type CreateCropPlanHandler struct {
	UowFactory repository.Factory
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

	return nil
}
