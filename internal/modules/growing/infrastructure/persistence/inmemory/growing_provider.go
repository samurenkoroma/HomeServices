package inmemory

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/phenology"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/task"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
	"samurenkoroma/services/internal/modules/growing/infrastructure/providers/weather"
	"sync"
)

type GrowingProvider struct {
	tx               *sql.Tx
	tasks            task.Repository
	cropplans        cropplan.Repository
	catalogs         catalog.Repository
	phenologyService phenology.PhenologyService

	oncePlan sync.Once
}

var (
	instance *GrowingProvider
	once     sync.Once
)

func (p *GrowingProvider) ProviderName() string {
	return "growing"
}

func NewGrowingProvider(tx *sql.Tx) repository.RepositoryProvider {

	return &GrowingProvider{
		tx: tx,
	}
}

func (p *GrowingProvider) PhenologyService() phenology.PhenologyService {
	if p.phenologyService == nil {

		p.phenologyService = phenology.NewPhenologyService(
			p.Catalogs(),
			weather.NewMockWeatherProvider(),
			true,
		)
	}
	return p.phenologyService
}

func (p *GrowingProvider) CropPlans() cropplan.Repository {
	p.oncePlan.Do(func() {
		p.cropplans = NewCropPlanRepo()
	})
	return p.cropplans
}

func (p *GrowingProvider) Catalogs() catalog.Repository {
	if p.catalogs == nil {
		p.catalogs = postgres.NewCatalogRepository(p.tx)
	}
	return p.catalogs
}
func (p *GrowingProvider) Tasks() task.Repository {
	if p.tasks == nil {
		p.tasks = NewTaskRepo()
	}
	return p.tasks
}
