package application

import "samurenkoroma/services/internal/crop/domain"

type CropPlanRepository interface {
	ByID(id domain.CropTypeID) (*domain.CropPlan, error)
	Save(plan *domain.CropPlan) error
}
