package repository

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/core/port/messaging"
)

type Factory interface {
	Begin(ctx context.Context, providerName string) (UnitOfWork, error)
}

type UoWFactory struct {
	db  *sql.DB
	bus messaging.EventBus
}

func NewUnitOfWorkFactory(db *sql.DB, bus messaging.EventBus) Factory {
	return &UoWFactory{
		db:  db,
		bus: bus,
	}
}

func (f *UoWFactory) Begin(ctx context.Context, providerName string) (UnitOfWork, error) {
	tx, err := f.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return NewUnitOfWork(ctx, tx, f.bus), nil
}
