package uow

import (
	"samurenkoroma/services/internal/common/domain"
	crop "samurenkoroma/services/internal/crop/application"
	growing "samurenkoroma/services/internal/growing/application"
)

type UnitOfWork interface {
	RegisterAggregate(agg domain.Aggregate)

	Commit() error
	Rollback() error

	EventBus() domain.EventBus
	GrowingFacilities() growing.GrowingFacilitiesRepository
	CropPlans() crop.CropPlanRepository
}
