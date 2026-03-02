package application

import (
	"context"
	"samurenkoroma/services/internal/growing/domain/cropplan"
	"samurenkoroma/services/internal/growing/domain/facility"
)

type FacilityOverviewDTO struct {
	ID     string
	Name   string
	Type   string
	Length float64
	Width  float64
}

type FacilityReadRepository interface {
	GetOverview(ctx context.Context, id string) (*FacilityOverviewDTO, error)
}

type GrowingFacilitiesRepository interface {
	Get(id facility.FacilityID) (*facility.GrowingFacility, error)
	Save(unit *facility.GrowingFacility) error
}

type CropPlanRepository interface {
	Get(id cropplan.CropPlanID) (*cropplan.CropPlan, error)
	Save(plan *cropplan.CropPlan) error
}
