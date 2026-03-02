package c

import (
	"context"
	"database/sql"
	"net/http"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/common/application/uow"
	growingCommand "samurenkoroma/services/internal/growing/application/command"
	"samurenkoroma/services/internal/growing/application/events"

	growingQuery "samurenkoroma/services/internal/growing/application/query"
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
	bus.Register("FacilityCreated", events.OnFacilityCreated)
	uowFactory := uowSql.NewUnitOfWorkFactory(db, bus)

	// ---------- Routers ----------

	commandRouter := command.NewRouter()
	queryRouter := query.NewRouter()

	// ---------- Register Bounded Contexts ----------

	if err := registerGrowing(
		commandRouter,
		queryRouter,
		uowFactory,
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

func registerGrowing(
	commandRouter command.Router,
	queryRouter query.Router,
	uowFactory uow.Factory,
) error {

	// ---- Command Registration ----
	commandRouter.Register(
		"CreateFacility",
		growingCommand.NewCreateFacilityHandler(uowFactory),
		growingCommand.DecodeCreateField,
	)
	commandRouter.Register(
		"AddBed",
		growingCommand.NewAddBedHandler(uowFactory),
		growingCommand.DecodeAddBed,
	)
	commandRouter.Register(
		"AddBlock",
		growingCommand.NewAddBlockHandler(uowFactory),
		growingCommand.DecodeAddBlock, // нужно будет реализовать аналогично
	)

	// ---- Query Handlers ----
	getOverviewHandler := growingQuery.NewGetFacilityOverviewHandler()

	queryRouter.Register(
		"GetFacilityOverview",
		query.DecodeJSON[growingQuery.GetFacilityOverviewQuery],
		getOverviewHandler.AsHandler(),
	)

	return nil
}
