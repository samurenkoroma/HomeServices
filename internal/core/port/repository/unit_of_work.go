package repository

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/port/messaging"
	crop "samurenkoroma/services/internal/modules/crop/domain"
	domain2 "samurenkoroma/services/internal/modules/farm/field/domain"
	"samurenkoroma/services/internal/modules/growing/domain"
)

type UnitOfWork interface {
	RegisterAggregate(agg aggregate.Aggregate)

	Commit() error
	Rollback() error

	EventBus() messaging.EventBus
	GrowingFacilities() domain2.FieldRepository
	CropPlans() crop.CropPlanRepository
	CropCycles() domain.CropCycleRepository
	CropTemplates() domain.CropTemplateRepository
}
