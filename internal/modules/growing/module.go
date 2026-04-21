package growing

import (
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/module"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/core/domain/repository"
	growingCommands "samurenkoroma/services/internal/modules/growing/application/command"
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
	return &growingModule{
		Commands: []module.CommandHandler{{
			Name:    "ActivateCropPlan",
			Handler: growingCommands.NewActivateCropPlanCmd(uowFactory),
			Decoder: utils.DecodeJSON[growingCommands.ActivateCropPlanCmd],
		}, {
			Name:    "AddStage",
			Handler: growingCommands.NewAddStageHandler(uowFactory),
			Decoder: utils.DecodeJSON[growingCommands.AddStageCmd],
		}, {
			Name:    "CompleteCropPlan",
			Handler: growingCommands.NewCompleteCropPlanHandler(uowFactory),
			Decoder: utils.DecodeJSON[growingCommands.CompleteCropPlanCmd],
		}, {
			Name:    "CompleteStage",
			Handler: growingCommands.NewCompleteStageCommand(uowFactory),
			Decoder: utils.DecodeJSON[growingCommands.CompleteStageCmd],
		}, {
			Name:    "CreateCropPlan",
			Handler: growingCommands.NewCreateCropPlanHandler(uowFactory),
			Decoder: utils.DecodeJSON[growingCommands.CreateCropPlanCmd],
		}, {
			Name:    "SkipStage",
			Handler: growingCommands.NewSkipStageCommand(uowFactory),
			Decoder: utils.DecodeJSON[growingCommands.SkipStageCmd],
		}, {
			Name:    "StartStage",
			Handler: growingCommands.NewStartStageCommand(uowFactory),
			Decoder: utils.DecodeJSON[growingCommands.StartStageCmd],
		}, {
			Name:    "CreateSeason",
			Handler: growingCommands.NewCreateSeasonHandler(uowFactory),
			Decoder: utils.DecodeJSON[growingCommands.CreateSeasonCmd],
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
			Name:    "GetSpecies",
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
