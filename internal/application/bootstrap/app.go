package c

import (
	"context"
	"database/sql"
	"net/http"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/core/domain/repository"
	inmemory "samurenkoroma/services/internal/infrastructure/messaging/rabbitmq"
	uowSql "samurenkoroma/services/internal/infrastructure/persistence/sql"
	"samurenkoroma/services/internal/interfaces/httpapi"
	cropCommand "samurenkoroma/services/internal/modules/crop/application/commands"
	fieldCommands "samurenkoroma/services/internal/modules/farm/application/commands"
	"samurenkoroma/services/internal/modules/farm/application/handlers"

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
	bus.Register("FacilityCreated", handlers.OnFieldCreated)
	uowFactory := uowSql.NewUnitOfWorkFactory(db, bus)

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
	commandRouter.Register("CreateField", fieldCommands.NewCreateFieldHandler(uowFactory), fieldCommands.DecodeCreateField)
	commandRouter.Register("CreateGreenhouse", fieldCommands.NewCreateGreenhouseHandler(uowFactory), fieldCommands.DecodeCreateGreenhouse)
	commandRouter.Register("CreateBed", fieldCommands.NewCreateBedHandler(uowFactory), fieldCommands.DecodeCreateBed)
	commandRouter.Register("CreateFieldBlock", fieldCommands.NewCreateFieldBlockHandler(uowFactory), fieldCommands.DecodeCreateFieldBlock)
	commandRouter.Register("CreateCropPlan", &cropCommand.CreateCropPlanHandler{UowFactory: uowFactory}, cropCommand.DecodeCreateCropPlan)

	// ---- Query Handlers ----
	//facilityReadRepo := postgres.NewGrowingFacilitiesRepository(uowFactory.Begin())
	//
	//getOverviewHandler := query2.NewGetFacilityOverviewHandler(facilityReadRepo)
	//getListHandler := query2.NewGetFacilitiesListHandler(facilityReadRepo)
	//
	//queryRouter.Register("GetFacilityOverview", query.DecodeJSON[query2.GetFacilityOverviewQuery], getOverviewHandler.AsHandler())
	//queryRouter.Register("GetFacilitiesList", query.DecodeJSON[query2.GetFacilitiesListQuery], getListHandler.AsHandler())

	return nil
}
