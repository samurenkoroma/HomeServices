package growing

import (
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/module"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/application/command/cropplan"
	"samurenkoroma/services/internal/modules/growing/application/command/season"
	"samurenkoroma/services/internal/modules/growing/application/command/stage"
	"samurenkoroma/services/internal/modules/growing/application/query/catalog"
	cropplanQuery "samurenkoroma/services/internal/modules/growing/application/query/cropplan"
	"samurenkoroma/services/internal/modules/growing/application/query/seasons"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
	"samurenkoroma/services/pkg/utils"
)

type growingModule struct {
	Commands []module.CommandHandler
	Queries  []module.QueryHandler
}

func NewModule(uowFactory repository.Factory) module.Module {
	tx, _ := uowFactory.DB().Begin()
	pr := postgres.NewPostgresGrowingProvider(tx).(*postgres.PostgresGrowingProvider)
	return &growingModule{
		Commands: []module.CommandHandler{{
			Name:    "ActivateCropPlan",
			Handler: cropplan.NewCropPlanHandler(uowFactory).Activate,
			Decoder: utils.DecodeJSON[cropplan.ActivateCropPlanCmd],
		}, {
			Name:    "CreateCropPlan",
			Handler: cropplan.NewCropPlanHandler(uowFactory).Create,
			Decoder: utils.DecodeJSON[cropplan.CreateCropPlanCmd],
		}, {
			Name:    "CompleteCropPlan",
			Handler: cropplan.NewCropPlanHandler(uowFactory).Complete,
			Decoder: utils.DecodeJSON[cropplan.CompleteCropPlanCmd],
		}, {
			Name:    "AddStage",
			Handler: stage.NewStageHandler(uowFactory).Add,
			Decoder: utils.DecodeJSON[stage.AddStageCmd],
		}, {
			Name:    "CompleteStage",
			Handler: stage.NewStageHandler(uowFactory).Complete,
			Decoder: utils.DecodeJSON[stage.CompleteStageCmd],
		}, {
			Name:    "SkipStage",
			Handler: stage.NewStageHandler(uowFactory).Skip,
			Decoder: utils.DecodeJSON[stage.SkipStageCmd],
		}, {
			Name:    "StartStage",
			Handler: stage.NewStageHandler(uowFactory).Start,
			Decoder: utils.DecodeJSON[stage.StartStageCmd],
		}, {
			Name:    "CreateSeason",
			Handler: season.NewSeasonHandler(uowFactory).Create,
			Decoder: utils.DecodeJSON[season.CreateSeasonCmd],
		}},
		Queries: []module.QueryHandler{{
			Name:    "ListSeasons",
			Handler: seasons.NewSeasonsHandler(pr.Seasons()).List,
			Decoder: utils.DecodeJSON[seasons.ListSeasonsQuery],
		}, {
			Name:    "GetCrops",
			Handler: catalog.NewCatalogHandler(pr.Catalog()).GetCrops,
			Decoder: utils.DecodeJSON[catalog.CropsQuery],
		}, {
			Name:    "GetVarieties",
			Handler: catalog.NewCatalogHandler(pr.Catalog()).GetVarieties,
			Decoder: utils.DecodeJSON[catalog.VarietiesQuery],
		}, {
			Name:    "GetCropPlans",
			Handler: cropplanQuery.NewQueryCropPlanHandler(pr.CropPlans(), pr.Catalog(), pr.Tasks(), pr.PhenologyService()).GetCropPlans,
			Decoder: utils.DecodeJSON[cropplanQuery.GetCropPlanQuery],
		},
		},
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
