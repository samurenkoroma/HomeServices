package cropplan

import "context"

type Repository interface {
	Save(ctx context.Context, plan *CropPlan) error
	GetByID(ctx context.Context, id PlanID) (*CropPlan, error)
}
