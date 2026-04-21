package inmemory

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/phenology"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/task"
	"samurenkoroma/services/internal/modules/growing/domain/season"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
	"samurenkoroma/services/internal/modules/growing/infrastructure/providers/weather"
	"sync"
)

type RedisGrowingProvider struct {
	tx               *sql.Tx
	tasks            task.Repository
	cropplans        cropplan.Repository
	catalog          catalog.Repository
	phenologyService phenology.PhenologyService
	seasons          season.Repository

	oncePlan sync.Once
}

var (
	instance *RedisGrowingProvider
	once     sync.Once
)

func (p *RedisGrowingProvider) ProviderName() string {
	return "growing"
}

func NewRedisGrowingProvider(tx *sql.Tx) repository.RepositoryProvider {
	return &RedisGrowingProvider{
		tx: tx,
	}

}

func (p *RedisGrowingProvider) PhenologyService() phenology.PhenologyService {
	if p.phenologyService == nil {

		p.phenologyService = phenology.NewPhenologyService(
			p.Catalog(),
			weather.NewMockWeatherProvider(),
			true,
		)
	}
	return p.phenologyService
}

func (p *RedisGrowingProvider) CropPlans() cropplan.Repository {
	p.oncePlan.Do(func() {
		p.cropplans = NewCropPlanRepo()
	})
	return p.cropplans
}

func (p *RedisGrowingProvider) Seasons() season.Repository {
	if p.seasons == nil {
		p.seasons = postgres.NewSeasonRepository(p.tx)
	}
	return p.seasons
}

func (p *RedisGrowingProvider) Catalog() catalog.Repository {
	if p.catalog == nil {
		p.catalog = postgres.NewCatalogRepository(p.tx)
	}
	return p.catalog
}
func (p *RedisGrowingProvider) Tasks() task.Repository {
	if p.tasks == nil {
		p.tasks = NewTaskRepo()
	}
	return p.tasks
}
