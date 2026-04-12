package inmemory

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/task"
)

type GrowingProvider struct {
	tasks     task.Repository
	cropplans cropplan.Repository
	catalogs  catalog.Repository
}

func (p *GrowingProvider) ProviderName() string {
	return "growing"
}

func NewGrowingProvider(tx *sql.Tx) repository.RepositoryProvider {
	return &GrowingProvider{}
}

func (p *GrowingProvider) CropPlans() cropplan.Repository {
	if p.cropplans == nil {
		p.cropplans = NewCropPlanRepo()
	}
	return p.cropplans
}

func (p *GrowingProvider) Catalogs() catalog.Repository {
	if p.catalogs == nil {
		p.catalogs = NewInMemoryCatalogRepository()
	}
	return p.catalogs
}
func (p *GrowingProvider) Tasks() task.Repository {
	if p.tasks == nil {
		p.tasks = NewTaskRepo()
	}
	return p.tasks
}
