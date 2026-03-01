package application

import "context"

type UnitOfWork interface {
	GrowingFacilities() GrowingFacilitiesRepository
	CropPlans() CropPlanRepository
	Commit() error
	Rollback() error
}

type UnitOfWorkFactory interface {
	New(ctx context.Context) (UnitOfWork, error)
}
