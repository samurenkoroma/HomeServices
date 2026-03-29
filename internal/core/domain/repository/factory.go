package repository

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/core/port/messaging"
)

type Factory interface {
	Begin(ctx context.Context) (UnitOfWork, error)
	DB() *sql.DB
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

func (f *UoWFactory) DB() *sql.DB {
	return f.db
}

func (f *UoWFactory) Begin(ctx context.Context) (UnitOfWork, error) {
	tx, err := f.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return NewUnitOfWork(ctx, tx, f.db, f.bus), nil
}
