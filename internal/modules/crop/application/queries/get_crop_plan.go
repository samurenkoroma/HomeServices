package queries

import (
	"context"
)

type GetCropPlanQuery struct {
	PlanID string `json:"plan_id" validate:"required"`
}

type CropPlanDetailDTO struct {
	ID            string            `json:"id"`
	CropTypeID    string            `json:"crop_type_id"`
	CropTypeName  string            `json:"crop_type_name"`
	VarietyID     *string           `json:"variety_id"`
	VarietyName   *string           `json:"variety_name"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	Duration      int               `json:"duration"`
	Version       int               `json:"version"`
	Status        string            `json:"status"`
	Stages        []StageDTO        `json:"stages"`
	RotationRules []RotationRuleDTO `json:"rotation_rules"`
	Environment   EnvironmentDTO    `json:"environment"`
	Nutrients     NutrientsDTO      `json:"nutrients"`
	CreatedAt     time.Time         `json:"created_at"`
	PublishedAt   *time.Time        `json:"published_at"`
}

type StageDTO struct {
	Order       int     `json:"order"`
	Name        string  `json:"name"`
	Duration    int     `json:"duration"`
	MinTemp     float64 `json:"min_temp"`
	MaxTemp     float64 `json:"max_temp"`
	OptimalTemp float64 `json:"optimal_temp"`
	WaterPerDay float64 `json:"water_per_day"`
}

type RotationRuleDTO struct {
	PredecessorCropTypeID string `json:"predecessor_crop_type_id"`
	PredecessorName       string `json:"predecessor_name"`
	MinYears              int    `json:"min_years"`
	Recommended           bool   `json:"recommended"`
}

type GetCropPlanHandler struct {
	readRepo CropPlanReadRepository
}

func (h *GetCropPlanHandler) Handle(ctx context.Context, q GetCropPlanQuery) (*CropPlanDetailDTO, error) {
	return h.readRepo.GetByID(ctx, q.PlanID)
}
