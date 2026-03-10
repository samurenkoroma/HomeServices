package repository

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/port/messaging"
	crop "samurenkoroma/services/internal/modules/crop/domain"
	"samurenkoroma/services/internal/modules/farm/domain/bed"
	"samurenkoroma/services/internal/modules/farm/domain/block"
	"samurenkoroma/services/internal/modules/farm/domain/field"
	"samurenkoroma/services/internal/modules/farm/domain/greenhouse"
	"samurenkoroma/services/internal/modules/growing/domain"
)

type UnitOfWork interface {
	RegisterAggregate(agg aggregate.Aggregate)

	Commit() error
	Rollback() error

	EventBus() messaging.EventBus
	Fields() field.Repository
	Beds() bed.Repository
	Blocks() block.Repository
	Greenhouses() greenhouse.Repository
	CropPlans() crop.CropPlanRepository
	CropCycles() domain.CropCycleRepository
	CropTemplates() domain.CropTemplateRepository
}
