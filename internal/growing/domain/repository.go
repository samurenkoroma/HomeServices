package domain

import (
	"context"
	"samurenkoroma/services/internal/growing/domain/facility"
	"time"
)

type FacilityOverviewDTO struct {
	ID     string
	Name   string
	Type   string
	Length float64
	Width  float64
}

type FacilitiesListItemDTO struct {
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
type FacilitiesListDTO struct {
	Items      []FacilitiesListItemDTO `json:"items"`
	Total      int                     `json:"total"`
	Page       int                     `json:"page"`
	Limit      int                     `json:"limit"`
	TotalPages int                     `json:"total_pages"`
}

type FacilitiesListParams struct {
	Limit int
	Page  int
}

type FacilityReadRepository interface {
	GetOverview(ctx context.Context, id string) (*FacilityOverviewDTO, error)
	GetList(ctx context.Context, params FacilitiesListParams) (*FacilitiesListDTO, error)
}

type GrowingFacilitiesRepository interface {
	Get(id facility.FacilityID) (*facility.GrowingFacility, error)
	Save(unit *facility.GrowingFacility) error
}
