package uow

import (
	"samurenkoroma/services/internal/common/domain"
	"samurenkoroma/services/internal/growing/application"
)

type UnitOfWork interface {
	RegisterAggregate(agg domain.Aggregate)

	Commit() error
	Rollback() error

	EventBus() domain.EventBus
	GrowingFacilities() application.GrowingFacilitiesRepository
	CropPlans() application.CropPlanRepository
}
