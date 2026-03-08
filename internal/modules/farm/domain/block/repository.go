package bed

import (
	"context"
	"samurenkoroma/services/internal/modules/farm/domain"
)

type BedOverviewDTO struct {
	ID     string
	Name   string
	Length float64
	Width  float64
}

type BedListItemDTO struct {
	Id   string  `json:"id"`
	Name string  `json:"name"`
	Area float64 `json:"area"`
}
type FacilitiesListDTO struct {
	Items      []BedListItemDTO `json:"items"`
	Total      int              `json:"total"`
	Page       int              `json:"page"`
	Limit      int              `json:"limit"`
	TotalPages int              `json:"total_pages"`
}

type BedListParams struct {
	Limit int
	Page  int
}

type BedReadRepository interface {
	GetOverview(ctx context.Context, id string) (*BedOverviewDTO, error)
	GetList(ctx context.Context, params BedListParams) (*FacilitiesListDTO, error)
}

type BedRepository interface {
	Get(id domain.GrowingAreaID) (*Bed, error)
	Save(unit *Bed) error
}
