package repository

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/port/messaging"
	crop "samurenkoroma/services/internal/modules/crop/domain"
	"samurenkoroma/services/internal/modules/growing/domain"
)

type UnitOfWork interface {
	RegisterAggregate(agg aggregate.Aggregate)

	Commit() error
	Rollback() error

	EventBus() messaging.EventBus
	GrowingFacilities() domain.GrowingFacilitiesRepository
	CropPlans() crop.CropPlanRepository
	CropCycles() domain.CropCycleRepository
	CropTemplates() domain.CropTemplateRepository
}
