package projections

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
)

type FarmProjectionsProvider struct {
	db *sql.DB

	objects physicalobject.ObjectProjections
}

func (provider *FarmProjectionsProvider) ProviderName() string {
	return "farm"
}

var _ repository.ProjectionProvider = (*FarmProjectionsProvider)(nil)

func NewFarmProjectionsProvider(db *sql.DB) *FarmProjectionsProvider {
	return &FarmProjectionsProvider{
		db: db,
	}
}

func (provider *FarmProjectionsProvider) Objects() physicalobject.ObjectProjections {
	if provider.objects == nil {
		provider.objects = NewPoProjection(provider.db)
	}
	return provider.objects
}
