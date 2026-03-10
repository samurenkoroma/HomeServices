package block

import (
	"context"
	"samurenkoroma/services/internal/modules/farm/domain"
)

type OverviewDTO struct {
	ID     string
	Name   string
	Length float64
	Width  float64
}

type ListItemDTO struct {
	Id   string  `json:"id"`
	Name string  `json:"name"`
	Area float64 `json:"area"`
}
type ListDTO struct {
	Items      []ListItemDTO `json:"items"`
	Total      int           `json:"total"`
	Page       int           `json:"page"`
	Limit      int           `json:"limit"`
	TotalPages int           `json:"total_pages"`
}

type QueryParams struct {
	Limit int
	Page  int
}

type ReadRepository interface {
	GetOverview(ctx context.Context, id string) (*OverviewDTO, error)
	GetList(ctx context.Context, params QueryParams) (*ListDTO, error)
}

type Repository interface {
	Get(id domain.GrowingAreaID) (*FieldBlock, error)
	Save(unit *FieldBlock) error
}
