package projections

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
	"samurenkoroma/services/internal/modules/growing/domain/season"
)

type GrowingProjectionsProvider struct {
	db *sql.DB

	// Кеш репозиториев
	seasons season.Projections
	areas   cultivationarea.Projections
}

func (provider *GrowingProjectionsProvider) ProviderName() string {
	return "crop"
}

var _ repository.ProjectionProvider = (*GrowingProjectionsProvider)(nil)

func NewGrowingProjectionsProvider(db *sql.DB) *GrowingProjectionsProvider {
	return &GrowingProjectionsProvider{
		db: db,
	}
}

func (provider *GrowingProjectionsProvider) Seasons() season.Projections {
	if provider.seasons == nil {
		provider.seasons = NewSeasonProjection(provider.db)
	}
	return provider.seasons
}

func (provider *GrowingProjectionsProvider) Areas() cultivationarea.Projections {
	if provider.areas == nil {
		provider.areas = NewAreaProjection(provider.db)
	}
	return provider.areas
}
