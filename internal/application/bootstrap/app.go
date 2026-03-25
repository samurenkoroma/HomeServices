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
	cropCommands "samurenkoroma/services/internal/modules/crop/application/commands"
	cropQueries "samurenkoroma/services/internal/modules/crop/application/queries"
	postgres2 "samurenkoroma/services/internal/modules/crop/infrastructure/persistence/postgres"
	cropProjection "samurenkoroma/services/internal/modules/crop/infrastructure/persistence/projections"
	farmCommands "samurenkoroma/services/internal/modules/farm/application/commands"
	farmEventHandlers "samurenkoroma/services/internal/modules/farm/application/handlers"
	farmQueries "samurenkoroma/services/internal/modules/farm/application/queries"
	projections2 "samurenkoroma/services/internal/modules/farm/infrastructure/persistence/projections"
	growingCommands "samurenkoroma/services/internal/modules/growing/application/commands"
	growingEventHandlers "samurenkoroma/services/internal/modules/growing/application/eventhandlers"

	_ "github.com/lib/pq"
)

type App struct {
	DB            *sql.DB
	CommandRouter command.Router
	QueryRouter   query.Router
	HTTPHandler   http.Handler
}

func Build(ctx context.Context, db *sql.DB) (*App, error) {
	// ---------- Unit Of Work Factory ----------

	bus := inmemory.NewInMemoryEventBus()
	bus.Register("farm.field.created", farmEventHandlers.OnFieldCreated)
	bus.Register("farm.greenhouse.created", farmEventHandlers.OnGreenhouseCreated)
	bus.Register("crop.plan.published", growingEventHandlers.OnCropPlanPublished)

	uowFactory := repository.NewUnitOfWorkFactory(db, bus)

	// ---------- Routers ----------

	commandRouter := command.NewRouter()
	queryRouter := query.NewRouter()

	// ---------- Register Bounded Contexts ----------

	if err := registerGrowing(
		commandRouter,
		queryRouter,
		uowFactory,
		db,
	); err != nil {
		return nil, err
	}

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
	commandRouter.Register(farmCommands.NewCreatePhysicalObjectHandler(uowFactory), command.DecodeCmd[farmCommands.CreatePhysicalObjectCmd])
	commandRouter.Register(farmCommands.NewUpdatePhysicalObjectHandler(uowFactory), command.DecodeCmd[farmCommands.UpdatePhysicalObjectCommand])

	commandRouter.Register(cropCommands.NewCreateCropPlanHandler(uowFactory), command.DecodeCmd[cropCommands.CreateCropPlanCmd])
	commandRouter.Register(cropCommands.NewCreateCropTypeHandler(uowFactory), command.DecodeCmd[cropCommands.CreateCropTypeCmd])
	commandRouter.Register(cropCommands.NewCreateVarietyHandler(uowFactory), command.DecodeCmd[cropCommands.CreateVarietyCmd])
	commandRouter.Register(cropCommands.NewAddStageHandler(uowFactory), command.DecodeCmd[cropCommands.AddStageCmd])

	commandRouter.Register(growingCommands.NewCreateSeasonCommand(uowFactory), command.DecodeCmd[growingCommands.CreateSeasonCmd])
	commandRouter.Register(growingCommands.NewRecordOperationHandler(uowFactory), command.DecodeCmd[growingCommands.RecordOperationCmd])
	commandRouter.Register(growingCommands.NewConfigureAreaHandler(uowFactory), command.DecodeCmd[growingCommands.ConfigureAreaCmd])
	commandRouter.Register(growingCommands.NewStartCropCycleHandler(uowFactory), command.DecodeCmd[growingCommands.StartCropCycleCmd])

	//commandRouter.Register("CreateGreenhouse", fieldCommands.NewCreateGreenhouseHandler(uowFactory), command.DecodeCmd[fieldCommands.CreateGreenhouseCmd])
	//commandRouter.Register("CreateFieldBlock", fieldCommands.NewCreateFieldBlockHandler(uowFactory), command.DecodeCmd[fieldCommands.CreateFieldBlockCmd])
	//commandRouter.Register("CreateBed", fieldCommands.NewCreateBedHandler(uowFactory), command.DecodeCmd[fieldCommands.CreateBedCmd])
	//commandRouter.Register("CreateCropPlan", &cropCommand.CreateCropPlanHandler{UowFactory: uowFactory}, cropCommand.DecodeCreateCropPlan)

	// ---- Query Handlers ----
	//facilityReadRepo := postgres.NewGrowingFacilitiesRepository(uowFactory.Begin())
	//
	//getOverviewHandler := query2.NewGetFacilityOverviewHandler(facilityReadRepo)
	//getListHandler := queries.NewListObjectsOnMapHandler(facilityReadRepo)
	//

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	cropProjector := cropProjection.NewCropProjection(db)
	poProjector := projections2.NewPoProjection(db)
	queryRouter.Register(cropQueries.NewGetCropTypesHandler(postgres2.NewCropTypeRepository(tx)), query.DecodeJSON[cropQueries.GetCropTypesQuery])
	queryRouter.Register(cropQueries.NewGetVarietyHandler(cropProjector), query.DecodeJSON[cropQueries.GetVarietyQuery])
	queryRouter.Register(farmQueries.NewGetPhysicalObjectsHandler(poProjector), query.DecodeJSON[farmQueries.GetPhysicalObjectsQuery])

	return nil
}
