package field_block

import (
	"context"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
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
	Save(ctx context.Context, block *FieldBlock) error
	Update(ctx context.Context, block *FieldBlock) error
	FindByID(ctx context.Context, id types.FieldBlockId) (*FieldBlock, error)
	FindByFieldID(ctx context.Context, fieldID string) ([]*FieldBlock, error)
	FindAvailable(ctx context.Context) ([]*FieldBlock, error)
	FindByStatus(ctx context.Context, status valueobject.AreaStatus) ([]*FieldBlock, error)
	FindByCropCycleID(ctx context.Context, cropCycleID string) (*FieldBlock, error)
}
