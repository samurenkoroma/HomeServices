package application

import (
	"context"
	"samurenkoroma/services/internal/crop/domain"
)

type CropPlanRepository interface {
	ByID(context.Context, domain.CropPlanID) (*domain.CropPlan, error)
	Save(context.Context, *domain.CropPlan) error
}
