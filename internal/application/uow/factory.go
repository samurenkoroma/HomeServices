package uow

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/shared/application"
)

type PgUnitOfWorkFactory struct {
	db       *sql.DB
	eventBus application.EventBus
}

func NewPgUnitOfWorkFactory(db *sql.DB, eventBus application.EventBus) *PgUnitOfWorkFactory {
	return &PgUnitOfWorkFactory{
		db:       db,
		eventBus: eventBus,
	}
}

func (f *PgUnitOfWorkFactory) New(ctx context.Context) (UnitOfWork, error) {

	tx, err := f.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return newPgUnitOfWork(tx, f.eventBus), nil
}
