package sql

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/core/port/messaging"
	"samurenkoroma/services/internal/core/port/repository"
)

type sqlUnitOfWorkFactory struct {
	db  *sql.DB
	bus messaging.EventBus
}

func NewUnitOfWorkFactory(db *sql.DB, bus messaging.EventBus) repository.Factory {
	return &sqlUnitOfWorkFactory{
		db:  db,
		bus: bus,
	}
}

func (f *sqlUnitOfWorkFactory) Begin(ctx context.Context) (repository.UnitOfWork, error) {
	tx, err := f.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return newSQLUnitOfWork(ctx, tx, f.bus), nil
}
