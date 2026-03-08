package c

import (
	"context"
	"database/sql"
	"net/http"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/core/port/repository"
	cropCommand "samurenkoroma/services/internal/modules/crop/application/commands"
	"samurenkoroma/services/internal/modules/farm/bed/application/commands"
	bedCommands "samurenkoroma/services/internal/modules/farm/bed/application/commands"
	blockCommands "samurenkoroma/services/internal/modules/farm/block/application/commands"
	fieldCommands "samurenkoroma/services/internal/modules/farm/field/application/commands"
	"samurenkoroma/services/internal/modules/farm/field/application/handlers"
	greenhouseCommands "samurenkoroma/services/internal/modules/farm/greenhouse/application/commands"
	query2 "samurenkoroma/services/internal/modules/growing/application/query"
	"samurenkoroma/services/internal/modules/growing/infrastructure/postgres"

	"samurenkoroma/services/internal/infrastructure/eventbus"
	uowSql "samurenkoroma/services/internal/infrastructure/persistence/sql"
	"samurenkoroma/services/internal/interfaces/httpapi"

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

	bus := eventbus.NewInMemoryEventBus()
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
	commandRouter.Register("CreateGreenhouse", greenhouseCommands.NewCreateGreenhouseHandler(uowFactory), greenhouseCommands.DecodeCreateField)
	commandRouter.Register("AddBed", bedCommands.NewAddBedHandler(uowFactory), commands.DecodeAddBed)
	commandRouter.Register("AddBlock", blockCommands.NewAddBlockHandler(uowFactory), blockCommands.DecodeAddBlock)
	commandRouter.Register("CreateCropPlan", &cropCommand.CreateCropPlanHandler{UowFactory: uowFactory}, cropCommand.DecodeCreateCropPlan)

	// ---- Query Handlers ----
	facilityReadRepo := postgres.NewFacilityReadRepository(db)

	getOverviewHandler := query2.NewGetFacilityOverviewHandler(facilityReadRepo)
	getListHandler := query2.NewGetFacilitiesListHandler(facilityReadRepo)

	queryRouter.Register("GetFacilityOverview", query.DecodeJSON[query2.GetFacilityOverviewQuery], getOverviewHandler.AsHandler())
	queryRouter.Register("GetFacilitiesList", query.DecodeJSON[query2.GetFacilitiesListQuery], getListHandler.AsHandler())

	return nil
}
