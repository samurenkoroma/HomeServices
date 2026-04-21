package c

import (
	"context"
	"database/sql"
	"net/http"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/core/domain/repository"
	inmemory "samurenkoroma/services/internal/infrastructure/messaging/rabbitmq"
	"samurenkoroma/services/internal/interfaces/httpapi"
	farmCommands "samurenkoroma/services/internal/modules/farm/application/commands"
	farmQueries "samurenkoroma/services/internal/modules/farm/application/queries"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
	farmProjections "samurenkoroma/services/internal/modules/farm/infrastructure/persistence/projections"
	growingCommands "samurenkoroma/services/internal/modules/growing/application/command"
	growingEventHandlers "samurenkoroma/services/internal/modules/growing/application/eventhandlers"
	growingQueries "samurenkoroma/services/internal/modules/growing/application/query"
	//inmemory2 "samurenkoroma/services/internal/modules/growing/infrastructure/persistence/inmemory"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
	"samurenkoroma/services/pkg/utils"

	_ "github.com/lib/pq"
)

type App struct {
	DB            *sql.DB
	CommandRouter command.Router
	QueryRouter   query.Router
	HTTPHandler   http.Handler
}

func Build(ctx context.Context, db *sql.DB) (*App, error) {

	bus := inmemory.NewInMemoryEventBus()
	bus.Register("farm.field.created", growingEventHandlers.OnFarmObjectCreated)
	bus.Register(physicalobject.FarmObjectSchemaUpdatedEvent, growingEventHandlers.OnFarmObjectSchemaUpdated)
	//bus.Register("crop.plan.published", growingEventHandlers.OnCropPlanPublished)

	// ---------- Unit Of Work Factory ----------
	uowFactory := repository.NewUnitOfWorkFactory(db, bus)

	// ---------- Routers ----------

	commandRouter := command.NewRouter()
	queryRouter := query.NewRouter()

	// ---------- Register Bounded Contexts ----------

	if err := registerGrowing(commandRouter, queryRouter, uowFactory, db); err != nil {
		return nil, err
	}
	if err := registerFarm(commandRouter, queryRouter, uowFactory, db); err != nil {
		return nil, err
	}
	//if err := registerCrop(commandRouter, queryRouter, uowFactory, db); err != nil {
	//	return nil, err
	//}
	// можно добавить:
	// registerCrop(...)
	// registerSeason(...)

	// ---------- HTTP Layer ----------

	httpHandler := httpapi.NewRouter(
		commandRouter,
		queryRouter,
	)

	return &App{
		DB:            db,
		CommandRouter: commandRouter,
		QueryRouter:   queryRouter,
		HTTPHandler:   httpHandler,
	}, nil
}

func registerGrowing(commandRouter command.Router, queryRouter query.Router, uowFactory repository.Factory, db *sql.DB) error {
	// ---- Command Registration ----
	commandRouter.Register(growingCommands.NewActivateCropPlanCmd(uowFactory), utils.DecodeJSON[growingCommands.ActivateCropPlanCmd])
	commandRouter.Register(growingCommands.NewAddStageHandler(uowFactory), utils.DecodeJSON[growingCommands.AddStageCmd])
	commandRouter.Register(growingCommands.NewCompleteCropPlanHandler(uowFactory), utils.DecodeJSON[growingCommands.CompleteCropPlanCmd])
	commandRouter.Register(growingCommands.NewCompleteStageCommand(uowFactory), utils.DecodeJSON[growingCommands.CompleteStageCmd])
	commandRouter.Register(growingCommands.NewCreateCropPlanHandler(uowFactory), utils.DecodeJSON[growingCommands.CreateCropPlanCmd])
	commandRouter.Register(growingCommands.NewSkipStageCommand(uowFactory), utils.DecodeJSON[growingCommands.SkipStageCmd])
	commandRouter.Register(growingCommands.NewStartStageCommand(uowFactory), utils.DecodeJSON[growingCommands.StartStageCmd])
	commandRouter.Register(growingCommands.NewCreateSeasonHandler(uowFactory), utils.DecodeJSON[growingCommands.CreateSeasonCmd])
	//commandRouter.Register(growingCommands.NewActivateSeasonCommand(uowFactory), utils.DecodeJSON[growingCommands.ActivateSeasonCmd])
	//commandRouter.Register(growingCommands.NewRecordOperationHandler(uowFactory), utils.DecodeJSON[growingCommands.RecordOperationCmd])
	//commandRouter.Register(growingCommands.NewConfigureAreaHandler(uowFactory), utils.DecodeJSON[growingCommands.ConfigureAreaCmd])
	//commandRouter.Register(growingCommands.NewCreateBedsBatchHandler(uowFactory), utils.DecodeJSON[growingCommands.CreateBedsBatchCommand])
	//commandRouter.Register(growingCommands.NewStartCropCycleHandler(uowFactory), utils.DecodeJSON[growingCommands.StartCropCycleCmd])

	//growingProvider := projections.NewGrowingProjectionsProvider(db)
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	pr := postgres.NewPostgresGrowingProvider(tx).(*postgres.PostgresGrowingProvider)
	queryRouter.Register(growingQueries.ListSeasonsHandler(pr.Seasons()), utils.DecodeJSON[growingQueries.ListSeasonsQuery])
	queryRouter.Register(growingQueries.NewListVarietiesHandler(uowFactory), utils.DecodeJSON[growingQueries.ListVarietiesQuery])
	queryRouter.Register(growingQueries.NewGetSpeciesHandler(pr.Catalog()), utils.DecodeJSON[growingQueries.GetSpeciesQuery])
	queryRouter.Register(growingQueries.NewGetVarietyHandler(pr.Catalog()), utils.DecodeJSON[growingQueries.GetVarietyQuery])
	queryRouter.Register(growingQueries.NewGetCropPlanHandler(pr.CropPlans()), utils.DecodeJSON[growingQueries.GetCropPlanQuery])
	queryRouter.Register(growingQueries.NewListCropPlansHandler(pr.CropPlans()), utils.DecodeJSON[growingQueries.ListCropPlansQuery])
	queryRouter.Register(growingQueries.NewListSpeciesHandler(pr.Catalog()), utils.DecodeJSON[growingQueries.ListSpeciesQuery])
	//queryRouter.Register(growingQueries.NewGetPlanStatisticsHandler(pr.CropPlans(), pr.Tasks()), utils.DecodeJSON[growingQueries.GetPlanStatisticsQuery])
	//queryRouter.Register(growingQueries.NewGetCurrentPhenologyHandler(pr.CropPlans(), pr.Catalog(), pr.PhenologyService()), utils.DecodeJSON[growingQueries.GetCurrentPhenologyQuery])
	//queryRouter.Register(growingQueries.NewGetCultivationAreasHandler(growingProvider.Areas()), utils.DecodeJSON[growingQueries.GetCultivationAreasQuery])
	return nil
}

//	func registerCrop(commandRouter command.Router, queryRouter query.Router, uowFactory repository.Factory, db *sql.DB) error {
//		// ---- Command Registration ----
//		//commandRouter.Register(cropCommands.NewCreateCropPlanHandler(uowFactory), utils.DecodeJSON[cropCommands.CreateCropPlanCmd])
//		commandRouter.Register(cropCommands.NewCreateCropTypeHandler(uowFactory), utils.DecodeJSON[cropCommands.CreateCropTypeCmd])
//		commandRouter.Register(cropCommands.NewCreateVarietyHandler(uowFactory), utils.DecodeJSON[cropCommands.CreateVarietyCmd])
//		commandRouter.Register(cropCommands.NewAddStageHandler(uowFactory), utils.DecodeJSON[cropCommands.AddStageCmd])
//
//		cropProvider := cropProjections.NewCropProjectionsProvider(db)
//		queryRouter.Register(cropQueries.NewGetCropTypesHandler(cropProvider.CropTypes()), utils.DecodeJSON[cropQueries.GetCropTypesQuery])
//		queryRouter.Register(cropQueries.NewGetCropPlanHandler(cropProvider.CropPlans()), utils.DecodeJSON[cropQueries.GetCropPlanQuery])
//		queryRouter.Register(cropQueries.NewGetVarietyHandler(cropProvider.Varieties()), utils.DecodeJSON[cropQueries.GetVarietyQuery])
//		queryRouter.Register(cropQueries.NewGetCategoriesHandler(cropProvider), utils.DecodeJSON[cropQueries.GetCategoriesQuery])
//
//		return nil
//	}
func registerFarm(commandRouter command.Router, queryRouter query.Router, uowFactory repository.Factory, db *sql.DB) error {

	// ---- Command Registration ----
	commandRouter.Register(farmCommands.NewCreateFarmObjectHandler(uowFactory), utils.DecodeJSON[farmCommands.CreateFarmObjectCmd])
	commandRouter.Register(farmCommands.NewUpdateFarmObjectHandler(uowFactory), utils.DecodeJSON[farmCommands.UpdateFarmObjectCommand])
	commandRouter.Register(farmCommands.NewDeleteFarmObjectHandler(uowFactory), utils.DecodeJSON[farmCommands.DeleteFarmObjectCommand])

	farmProvider := farmProjections.NewFarmProjectionsProvider(db)
	queryRouter.Register(farmQueries.NewGetCurrentFarmHandler(farmProvider.Objects()), utils.DecodeJSON[farmQueries.GetCurrentFarmQuery])
	queryRouter.Register(farmQueries.NewGetPhysicalObjectsHandler(farmProvider.Objects()), utils.DecodeJSON[farmQueries.GetPhysicalObjectsQuery])

	return nil
}
