package growing

import (
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/module"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/application/command/cropplan"
	"samurenkoroma/services/internal/modules/growing/application/command/cultivation"
	"samurenkoroma/services/internal/modules/growing/application/command/season"
	"samurenkoroma/services/internal/modules/growing/application/query/catalog"
	cropplanQuery "samurenkoroma/services/internal/modules/growing/application/query/cropplan"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
	"samurenkoroma/services/internal/modules/growing/infrastructure/projections"
	"samurenkoroma/services/pkg/utils"
)

type growingModule struct {
	Commands []module.CommandHandler
	Queries  []module.QueryHandler
}

func NewModule(uowFactory repository.Factory) module.Module {
	tx, _ := uowFactory.DB().Begin()
	pr := postgres.NewPostgresGrowingProvider(tx).(*postgres.PostgresGrowingProvider)
	projector := projections.NewGrowingProjectionsProvider(uowFactory.DB())
	return &growingModule{
		Commands: []module.CommandHandler{{
			Name:    "ActivateCropPlan",
			Handler: cropplan.NewCropPlanHandler(uowFactory).Activate,
			Decoder: utils.DecodeJSON[cropplan.ActivateCropPlanCmd],
		}, {
			Name:    "createCropPlan",
			Handler: cropplan.NewCropPlanHandler(uowFactory).Create,
			Decoder: utils.DecodeJSON[cropplan.CreateCropPlanCmd],
		}, {
			Name:    "createSeason",
			Handler: season.NewSeasonHandler(uowFactory).Create,
			Decoder: utils.DecodeJSON[season.CreateSeasonCmd],
		}, {
			Name:    "CreateCultivationPlan",
			Handler: cultivation.NewCultivationPlanHandler(uowFactory).Create,
			Decoder: utils.DecodeJSON[cultivation.CreateCultivationPlanCmd],
		}},
		Queries: []module.QueryHandler{{
			Name:    "getSeasons",
			Handler: catalog.NewCatalogHandler(projector.Catalog()).GetSeasons,
			Decoder: utils.DecodeJSON[catalog.SeasonsQuery],
		}, {
			Name:    "getCrops",
			Handler: catalog.NewCatalogHandler(projector.Catalog()).GetCrops,
			Decoder: utils.DecodeJSON[catalog.CropsQuery],
		}, {
			Name:    "getVarieties",
			Handler: catalog.NewCatalogHandler(projector.Catalog()).GetVarieties,
			Decoder: utils.DecodeJSON[catalog.VarietiesQuery],
		}, {
			Name:    "getCropPlans",
			Handler: cropplanQuery.NewCropPlanQueryHandler(projector, pr).GetCropPlans,
			Decoder: utils.DecodeJSON[cropplanQuery.CropPlansQuery],
		}, {
			Name:    "getCultivationPlans",
			Handler: cropplanQuery.NewCropPlanQueryHandler(projector, pr).GetCultivationPlan,
			Decoder: utils.DecodeJSON[cropplanQuery.CultivationPlansQuery],
		}, {
			Name:    "getCultivationAreas",
			Handler: cropplanQuery.NewCropPlanQueryHandler(projector, pr).GetCultivationAreas,
			Decoder: utils.DecodeJSON[cropplanQuery.CultivationAreasQuery],
		}},
	}

}
func (f *growingModule) RegisterCommands(router command.Router) {
	for _, cmd := range f.Commands {
		router.Register(cmd.Name, cmd.Handler, cmd.Decoder)
	}
}

func (f *growingModule) RegisterQueries(router query.Router) {
	for _, q := range f.Queries {
		router.Register(q.Name, q.Handler, q.Decoder)
	}
}
