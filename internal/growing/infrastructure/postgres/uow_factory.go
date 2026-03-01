package postgres

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/application/shared"
	"samurenkoroma/services/internal/growing/application"
)

type PgUnitOfWorkFactory struct {
	db       *sql.DB
	eventBus shared.EventBus
}

func NewPgUnitOfWorkFactory(db *sql.DB, eventBus shared.EventBus) *PgUnitOfWorkFactory {
	return &PgUnitOfWorkFactory{
		db:       db,
		eventBus: eventBus,
	}
}

func (f *PgUnitOfWorkFactory) New(ctx context.Context) (application.UnitOfWork, error) {

	tx, err := f.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return newPgUnitOfWork(tx, f.eventBus), nil
}
