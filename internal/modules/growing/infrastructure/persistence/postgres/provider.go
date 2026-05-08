package postgres

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cultivation"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/phenology"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/worktask"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
	"samurenkoroma/services/internal/modules/growing/domain/season"
	"samurenkoroma/services/internal/modules/growing/infrastructure/providers/weather"
)

// PostgresGrowingProvider — провайдер репозиториев для контекста культур
type PostgresGrowingProvider struct {
	tx *sql.Tx

	cultivationAreaRepo cultivationarea.Repository
	catalog             catalog.Repository
	cropPlans           cropplan.Repository
	cultivationRepo     cultivation.Repository
	tasks               worktask.Repository
	phenologyService    phenology.PhenologyService
	seasons             season.Repository
}

func (p *PostgresGrowingProvider) Tasks() worktask.Repository {
	if p.tasks == nil {
		p.tasks = NewWorktaskRepository(p.tx)
	}
	return p.tasks
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
	if p.seasons == nil {
		p.seasons = NewSeasonRepository(p.tx)
	}
	return p.seasons
}

// CultivationAreas возвращает репозиторий всех объектов
func (p *PostgresGrowingProvider) CultivationAreas() cultivationarea.Repository {
	if p.cultivationAreaRepo == nil {
		p.cultivationAreaRepo = NewCultivationAreaRepository(p.tx)
	}
	return p.cultivationAreaRepo
}

func (p *PostgresGrowingProvider) Cultivation() cultivation.Repository {
	if p.cultivationRepo == nil {
		p.cultivationRepo = NewCultivationPlanRepository(p.tx)
	}
	return p.cultivationRepo
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

func (p *PostgresGrowingProvider) PhenologyService() phenology.PhenologyService {
	if p.phenologyService == nil {

		p.phenologyService = phenology.NewPhenologyService(
			p.Catalog(),
			weather.NewMockWeatherProvider(),
			true,
		)
	}
	return p.phenologyService
}
