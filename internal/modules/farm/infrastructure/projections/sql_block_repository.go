package projections

import "context"

type OverviewFieldBlockDTO struct {
	ID     string
	Name   string
	Length float64
	Width  float64
}

type FieldBlockItemDTO struct {
	Id   string  `json:"id"`
	Name string  `json:"name"`
	Area float64 `json:"area"`
}
type ListFieldBlockDTO struct {
	Items      []FieldBlockItemDTO `json:"items"`
	Total      int                 `json:"total"`
	Page       int                 `json:"page"`
	Limit      int                 `json:"limit"`
	TotalPages int                 `json:"total_pages"`
}

type FieldBlockQueryParams struct {
	Limit int
	Page  int
}

type ReadRepository interface {
	GetOverview(ctx context.Context, id string) (*OverviewFieldBlockDTO, error)
	GetList(ctx context.Context, params FieldBlockQueryParams) (*ListFieldBlockDTO, error)
}
