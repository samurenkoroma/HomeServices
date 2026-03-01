package sql

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/common/application/uow"
	"samurenkoroma/services/internal/common/domain"
)

type sqlUnitOfWorkFactory struct {
	db  *sql.DB
	bus domain.EventBus
}

func NewUnitOfWorkFactory(
	db *sql.DB,
	bus domain.EventBus,
) uow.Factory {
	return &sqlUnitOfWorkFactory{
		db:  db,
		bus: bus,
	}
}

func (f *sqlUnitOfWorkFactory) Begin(
	ctx context.Context,
) (uow.UnitOfWork, error) {

	tx, err := f.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return newSQLUnitOfWork(ctx, tx, f.bus), nil
}
