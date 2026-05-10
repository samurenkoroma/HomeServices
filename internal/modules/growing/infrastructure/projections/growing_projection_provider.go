package projections

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
)

type GrowingProjectionsProvider struct {
	db *sql.DB

	// Кеш репозиториев
	catalog catalog.CatalogProjections
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

func (provider *GrowingProjectionsProvider) Catalog() catalog.CatalogProjections {
	if provider.catalog == nil {
		provider.catalog = NewCatalogProjection(provider.db)
	}
	return provider.catalog
}

func (provider *GrowingProjectionsProvider) Areas() cultivationarea.Projections {
	if provider.areas == nil {
		provider.areas = NewAreaProjection(provider.db)
	}
	return provider.areas
}
