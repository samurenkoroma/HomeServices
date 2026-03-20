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
	farmCommands "samurenkoroma/services/internal/modules/farm/application/commands"
	"samurenkoroma/services/internal/modules/farm/application/handlers"
	"samurenkoroma/services/internal/modules/farm/application/queries"
	"samurenkoroma/services/internal/modules/farm/infrastructure/persistence/postgres"

	_ "github.com/lib/pq"
)

type App struct {
	DB            *sql.DB
	CommandRouter command.Router
	QueryRouter   query.Router
	HTTPHandler   http.Handler
}

func Build(ctx context.Context, dsn string) (*App, error) {

	// ---------- Database ----------

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	// ---------- Unit Of Work Factory ----------

	bus := inmemory.NewInMemoryEventBus()
	bus.Register("farm.field.created", handlers.OnFieldCreated)
	bus.Register("farm.greenhouse.created", handlers.OnGreenhouseCreated)

	uowFactory := postgres.NewFarmUnitOfWorkFactory(db, bus)

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
	commandRouter.Register(farmCommands.NewCreateFieldHandler(uowFactory), command.DecodeCmd[farmCommands.CreatePhysicalObjectCmd])
	commandRouter.Register(cropCommands.NewCreateCropPlanHandler(uowFactory), command.DecodeCmd[cropCommands.CreateCropPlanCommand])

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
	farmQueries := queries.NewListObjectsOnMapHandler(postgres.NewPhysicalObjectRepository(tx))
	queryRouter.Register("ListObjects", query.DecodeJSON[queries.ListObjectsOnMapQuery], farmQueries.AsHandler())
	//queryRouter.Register("GetFacilitiesList", query.DecodeJSON[query2.GetFacilitiesListQuery], getListHandler.AsHandler())

	return nil
}
