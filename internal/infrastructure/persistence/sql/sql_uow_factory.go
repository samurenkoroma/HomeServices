package sql

import (
	"context"
	"database/sql"
	repository2 "samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/port/messaging"
)

type sqlUnitOfWorkFactory struct {
	db  *sql.DB
	bus messaging.EventBus
}

func NewUnitOfWorkFactory(db *sql.DB, bus messaging.EventBus) repository2.Factory {
	return &sqlUnitOfWorkFactory{
		db:  db,
		bus: bus,
	}
}

func (f *sqlUnitOfWorkFactory) Begin(ctx context.Context) (repository2.UnitOfWork, error) {
	tx, err := f.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return newSQLUnitOfWork(ctx, tx, f.bus), nil
}
