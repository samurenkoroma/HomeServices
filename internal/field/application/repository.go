package application

import (
	"samurenkoroma/services/internal/field/domain/cropplan"
	"samurenkoroma/services/internal/field/domain/landunit"
)

type LandUnitRepository interface {
	Get(id landunit.LandUnitID) (*landunit.LandUnit, error)
	Save(unit *landunit.LandUnit) error
}

type CropPlanRepository interface {
	Get(id cropplan.CropPlanID) (*cropplan.CropPlan, error)
	Save(plan *cropplan.CropPlan) error
}
