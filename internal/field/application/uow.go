package application

import "context"

type UnitOfWork interface {
	LandUnits() LandUnitRepository
	CropPlans() CropPlanRepository
	Commit() error
	Rollback() error
}

type UnitOfWorkFactory interface {
	New(ctx context.Context) (UnitOfWork, error)
}
