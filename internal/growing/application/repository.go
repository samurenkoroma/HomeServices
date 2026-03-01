package application

import (
	"samurenkoroma/services/internal/growing/domain/cropplan"
	"samurenkoroma/services/internal/growing/domain/facility"
)

type GrowingFacilitiesRepository interface {
	Get(id facility.FacilityID) (*facility.GrowingFacility, error)
	Save(unit *facility.GrowingFacility) error
}

type CropPlanRepository interface {
	Get(id cropplan.CropPlanID) (*cropplan.CropPlan, error)
	Save(plan *cropplan.CropPlan) error
}
