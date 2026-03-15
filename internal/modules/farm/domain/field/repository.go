package field

import (
	"context"
	"samurenkoroma/services/internal/core/domain/types"
	"time"
)

type OverviewDTO struct {
	ID     string
	Name   string
	Type   string
	Length float64
	Width  float64
}

type ListItemDTO struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Area        float64   `json:"area"`
	ActiveCrops int       `json:"active_crops"`
	Status      string    `json:"status"`
	YieldTrend  float64   `json:"yield_trend"`
	Location    string    `json:"location"`
	Thumbnail   *string   `json:"thumbnail"`
	UpdatedAt   time.Time `json:"updated_at"`
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
	Save(unit *Field) error
	Update(ctx context.Context, field *Field) error
	FindByID(ctx context.Context, id types.FieldId) (*Field, error)
	FindAll(ctx context.Context) ([]*Field, error)
}
