package postgres

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/port/messaging"
)

type farmUoWFactory struct {
	db  *sql.DB
	bus messaging.EventBus
}

func NewFarmUnitOfWorkFactory(db *sql.DB, bus messaging.EventBus) repository.Factory {
	return &farmUoWFactory{
		db:  db,
		bus: bus,
	}
}

func (f *farmUoWFactory) Begin(ctx context.Context) (repository.UnitOfWork, error) {
	tx, err := f.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return NewFarmUnitOfWork(ctx, tx, f.bus), nil
}
