package queries

import (
	"context"
	"time"
)

type ListCropPlansQuery struct {
	CropTypeID string `json:"crop_type_id,omitempty"`
	Status     string `json:"status,omitempty"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
}

type CropPlanListItem struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	CropTypeID   string     `json:"crop_type_id"`
	CropTypeName string     `json:"crop_type_name"`
	Duration     int        `json:"duration"`
	Version      int        `json:"version"`
	Status       string     `json:"status"`
	PublishedAt  *time.Time `json:"published_at"`
	CreatedAt    time.Time  `json:"created_at"`
}

type ListCropPlansHandler struct {
	readRepo CropPlanReadRepository
}

func (h *ListCropPlansHandler) Handle(ctx context.Context, q ListCropPlansQuery) ([]CropPlanListItem, error) {
	return h.readRepo.GetList(ctx, q)
}
