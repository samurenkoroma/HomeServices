package application

type UnitOfWork interface {
	LandUnits() LandUnitRepository
	CropPlans() CropPlanRepository
	Commit() error
	Rollback() error
}
