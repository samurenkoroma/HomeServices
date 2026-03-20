package projections

import (
	"context"
	"time"
)

type OverviewFieldDTO struct {
	ID     string
	Name   string
	Type   string
	Length float64
	Width  float64
}

type FieldItemDTO struct {
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
type ListFieldDTO struct {
	Items      []FieldItemDTO `json:"items"`
	Total      int            `json:"total"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	TotalPages int            `json:"total_pages"`
}

type FieldQueryParams struct {
	Limit int
	Page  int
}

type FieldReadRepository interface {
	GetOverview(ctx context.Context, id string) (*OverviewFieldDTO, error)
	GetList(ctx context.Context, params FieldQueryParams) (*ListFieldDTO, error)
}
