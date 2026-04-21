package postgres

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
	"samurenkoroma/services/internal/modules/growing/domain/season"
)

// PostgresGrowingProvider — провайдер репозиториев для контекста культур
type PostgresGrowingProvider struct {
	tx *sql.Tx

	// Кеш репозиториев
	seasonsRepo         season.Repository
	cultivationAreaRepo cultivationarea.Repository
	catalog             catalog.Repository
	cropPlans           cropplan.Repository
}

func (p *PostgresGrowingProvider) ProviderName() string {
	return "growing"
}

// Проверяем, что PostgresGrowingProvider реализует интерфейс RepositoryProvider
var _ repository.RepositoryProvider = (*PostgresGrowingProvider)(nil)

func NewPostgresGrowingProvider(tx *sql.Tx) repository.RepositoryProvider {
	return &PostgresGrowingProvider{
		tx: tx,
	}
}

// Seasons возвращает репозиторий всех объектов
func (p *PostgresGrowingProvider) Seasons() season.Repository {
	if p.seasonsRepo == nil {
		p.seasonsRepo = NewSeasonRepository(p.tx)
	}
	return p.seasonsRepo
}

// CultivationAreas возвращает репозиторий всех объектов
func (p *PostgresGrowingProvider) CultivationAreas() cultivationarea.Repository {
	if p.cultivationAreaRepo == nil {
		p.cultivationAreaRepo = NewCultivationAreaRepository(p.tx)
	}
	return p.cultivationAreaRepo
}

func (p *PostgresGrowingProvider) Catalog() catalog.Repository {
	if p.catalog == nil {
		p.catalog = NewCatalogRepository(p.tx)
	}
	return p.catalog
}
func (p *PostgresGrowingProvider) CropPlans() cropplan.Repository {
	if p.cropPlans == nil {
		p.cropPlans = NewCropPlanRepository(p.tx)
	}
	return p.cropPlans
}
