package uow

import (
	"samurenkoroma/services/internal/common/domain"
	crop "samurenkoroma/services/internal/crop/domain"
	growing "samurenkoroma/services/internal/growing/domain"
)

type UnitOfWork interface {
	RegisterAggregate(agg domain.Aggregate)

	Commit() error
	Rollback() error

	EventBus() domain.EventBus
	GrowingFacilities() growing.GrowingFacilitiesRepository
	CropPlans() crop.CropPlanRepository
}
