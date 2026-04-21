package c

import (
	"context"
	"database/sql"
	"net/http"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/module"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/core/domain/repository"
	inmemory "samurenkoroma/services/internal/infrastructure/messaging/rabbitmq"
	"samurenkoroma/services/internal/interfaces/httpapi"
	"samurenkoroma/services/internal/modules/farm"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
	"samurenkoroma/services/internal/modules/growing"
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
	bus := inmemory.NewInMemoryEventBus()
	uowFactory := repository.NewUnitOfWorkFactory(db, bus)

	commandRouter := command.NewRouter()
	queryRouter := query.NewRouter()

	modules := []module.Module{
		farm.NewModule(uowFactory),
		growing.NewModule(uowFactory),
	}

	for _, m := range modules {
		m.RegisterCommands(commandRouter)
		m.RegisterQueries(queryRouter)
	}

	bus.Register("farm.field.created", growingEventHandlers.OnFarmObjectCreated)
	bus.Register(physicalobject.FarmObjectSchemaUpdatedEvent, growingEventHandlers.OnFarmObjectSchemaUpdated)
	//bus.Register("crop.plan.published", growingEventHandlers.OnCropPlanPublished)

	httpHandler := httpapi.NewRouter(commandRouter, queryRouter)

	return &App{
		DB:            db,
		CommandRouter: commandRouter,
		QueryRouter:   queryRouter,
		HTTPHandler:   httpHandler,
	}, nil
}
