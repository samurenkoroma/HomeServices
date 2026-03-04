package domain

import (
	"context"
)

type CropPlanRepository interface {
	ByID(context.Context, CropPlanID) (*CropPlan, error)
	Save(context.Context, *CropPlan) error
}
