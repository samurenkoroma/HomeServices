package growing

import (
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/module"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/application/command/cropplan"
	"samurenkoroma/services/internal/modules/growing/application/command/season"
	"samurenkoroma/services/internal/modules/growing/application/command/stage"
	growingQueries "samurenkoroma/services/internal/modules/growing/application/query"
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
	seasons := season.NewSeasonHandler(uowFactory)
	stages := stage.NewStageHandler(uowFactory)
	plans := cropplan.NewCropPlanHandler(uowFactory)
	return &growingModule{
		Commands: []module.CommandHandler{{
			Name:    "ActivateCropPlan",
			Handler: plans.Activate,
			Decoder: utils.DecodeJSON[cropplan.ActivateCropPlanCmd],
		}, {
			Name:    "CreateCropPlan",
			Handler: plans.Create,
			Decoder: utils.DecodeJSON[cropplan.CreateCropPlanCmd],
		}, {
			Name:    "CompleteCropPlan",
			Handler: plans.Complete,
			Decoder: utils.DecodeJSON[cropplan.CompleteCropPlanCmd],
		}, {
			Name:    "AddStage",
			Handler: stages.Add,
			Decoder: utils.DecodeJSON[stage.AddStageCmd],
		}, {
			Name:    "CompleteStage",
			Handler: stages.Complete,
			Decoder: utils.DecodeJSON[stage.CompleteStageCmd],
		}, {
			Name:    "SkipStage",
			Handler: stages.Skip,
			Decoder: utils.DecodeJSON[stage.SkipStageCmd],
		}, {
			Name:    "StartStage",
			Handler: stages.Start,
			Decoder: utils.DecodeJSON[stage.StartStageCmd],
		}, {
			Name:    "CreateSeason",
			Handler: seasons.Create,
			Decoder: utils.DecodeJSON[season.CreateSeasonCmd],
		}},
		Queries: []module.QueryHandler{{
			Name:    "ListSeasons",
			Handler: growingQueries.ListSeasonsHandler(pr.Seasons()),
			Decoder: utils.DecodeJSON[growingQueries.ListSeasonsQuery],
		}, {
			Name:    "ListVarieties",
			Handler: growingQueries.NewListVarietiesHandler(uowFactory),
			Decoder: utils.DecodeJSON[growingQueries.ListVarietiesQuery],
		}, {
			Name:    "GetSpecie",
			Handler: growingQueries.NewGetSpeciesHandler(pr.Catalog()),
			Decoder: utils.DecodeJSON[growingQueries.GetSpeciesQuery],
		}, {
			Name:    "GetVariety",
			Handler: growingQueries.NewGetVarietyHandler(pr.Catalog()),
			Decoder: utils.DecodeJSON[growingQueries.GetVarietyQuery],
		}, {
			Name:    "GetCropPlan",
			Handler: growingQueries.NewGetCropPlanHandler(pr.CropPlans()),
			Decoder: utils.DecodeJSON[growingQueries.GetCropPlanQuery],
		}, {
			Name:    "ListCropPlan",
			Handler: growingQueries.NewListCropPlansHandler(pr.CropPlans()),
			Decoder: utils.DecodeJSON[growingQueries.ListCropPlansQuery],
		}, {
			Name:    "ListSpecies",
			Handler: growingQueries.NewListSpeciesHandler(pr.Catalog()),
			Decoder: utils.DecodeJSON[growingQueries.ListSpeciesQuery],
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
