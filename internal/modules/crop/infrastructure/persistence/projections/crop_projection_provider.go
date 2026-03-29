package projections

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/crop/domain/cropplan"
	"samurenkoroma/services/internal/modules/crop/domain/croptype"
	"samurenkoroma/services/internal/modules/crop/domain/variety"
)

type CropProjectionsProvider struct {
	db *sql.DB

	// Кеш репозиториев
	cropTypes croptype.Projections
	cropPlans cropplan.Projections
	varieties variety.Projections
}

func (provider *CropProjectionsProvider) ProviderName() string {
	return "crop"
}

var _ repository.ProjectionProvider = (*CropProjectionsProvider)(nil)

func NewCropProjectionsProvider(db *sql.DB) *CropProjectionsProvider {
	return &CropProjectionsProvider{
		db: db,
	}
}

// Varieties возвращает репозиторий всех объектов
func (provider *CropProjectionsProvider) Varieties() variety.Projections {
	if provider.varieties == nil {
		provider.varieties = NewVarietyProjections(provider.db)
	}
	return provider.varieties
}

func (provider *CropProjectionsProvider) CropTypes() croptype.Projections {
	if provider.cropTypes == nil {
		provider.cropTypes = NewCropTypeProjections(provider.db)
	}
	return provider.cropTypes
}

func (provider *CropProjectionsProvider) CropPlans() cropplan.Projections {
	if provider.cropPlans == nil {
		provider.cropPlans = NewCropPlanProjections(provider.db)
	}
	return provider.cropPlans
}
