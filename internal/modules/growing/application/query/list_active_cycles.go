package query

import (
	"context"
	"time"
)

type ListActiveCyclesQuery struct {
	SeasonID string `json:"season_id,omitempty"`
	AreaID   string `json:"area_id,omitempty"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

type ActiveCycleDTO struct {
	ID         string    `json:"id"`
	AreaID     string    `json:"area_id"`
	AreaName   string    `json:"area_name"`
	CropPlanID string    `json:"crop_plan_id"`
	CropName   string    `json:"crop_name"`
	Status     string    `json:"status"`
	StartedAt  time.Time `json:"started_at"`
	Progress   float64   `json:"progress"` // % выполнения
}

type ListActiveCyclesHandler struct {
	readRepo CropCycleReadRepository
}

func (h *ListActiveCyclesHandler) Handle(ctx context.Context, q ListActiveCyclesQuery) ([]ActiveCycleDTO, error) {
	return h.readRepo.GetActiveCycles(ctx, q)
}
