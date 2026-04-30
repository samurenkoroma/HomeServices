package cropplan

import (
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/phenology"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/task"
)

type QueryHandler struct {
	PlanRepo         cropplan.Repository
	CatalogRepo      catalog.Repository
	PhenologyService phenology.PhenologyService
	TaskRepo         task.Repository
}

func NewQueryCropPlanHandler(
	PlanRepo cropplan.Repository,
	CatalogRepo catalog.Repository,
	TaskRepo task.Repository,
	PhenologyService phenology.PhenologyService) *QueryHandler {
	return &QueryHandler{
		TaskRepo:         TaskRepo,
		PlanRepo:         PlanRepo,
		CatalogRepo:      CatalogRepo,
		PhenologyService: PhenologyService,
	}
}
