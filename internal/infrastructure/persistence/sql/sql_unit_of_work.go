package sql

import (
	"context"
	"database/sql"
	"errors"
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/event"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/port/messaging"
	"samurenkoroma/services/internal/modules/growing/domain"
	"sync"
)

var (
	ErrAlreadyCommitted  = errors.New("unit of work already committed")
	ErrAlreadyRolledBack = errors.New("unit of work already rolled back")
)

func newSQLUnitOfWork(
	ctx context.Context,
	tx *sql.Tx,
	bus messaging.EventBus,
) repository.UnitOfWork {
	return &sqlUnitOfWork{}
}

func (u *sqlUnitOfWork) Execute(ctx context.Context, fn func(provider repository.RepositoryProvider) error) error {
	//TODO implement me
	panic("implement me")
}

type sqlUnitOfWork struct {
	ctx context.Context
	tx  *sql.Tx

	bus messaging.EventBus

	mu         sync.Mutex
	committed  bool
	rolledBack bool

	aggregates []aggregate.Aggregate
}

func (u *sqlUnitOfWork) RegisterAggregate(agg aggregate.Aggregate) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.aggregates = append(u.aggregates, agg)
}

func (u *sqlUnitOfWork) Commit() error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if u.committed {
		return ErrAlreadyCommitted
	}
	if u.rolledBack {
		return ErrAlreadyRolledBack
	}

	if err := u.tx.Commit(); err != nil {
		return err
	}

	return u.dispatchEvents()
}

func (u *sqlUnitOfWork) Rollback() error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if u.committed {
		return ErrAlreadyCommitted
	}
	if u.rolledBack {
		return ErrAlreadyRolledBack
	}
	u.rolledBack = true
	return u.tx.Rollback()
}

func (u *sqlUnitOfWork) EventBus() messaging.EventBus {
	return u.bus
}
func (u *sqlUnitOfWork) dispatchEvents() error {
	var allEvents []event.DomainEvent

	for _, agg := range u.aggregates {
		allEvents = append(allEvents, agg.PullEvents()...)
	}

	if len(allEvents) == 0 {
		return nil
	}

	return u.bus.Publish(repository2.WithUnitOfWork(u.ctx, u), allEvents)
}
