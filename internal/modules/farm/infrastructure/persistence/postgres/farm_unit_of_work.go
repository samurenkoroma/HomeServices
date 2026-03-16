package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/event"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/port/messaging"
	"sync"
)

var (
	ErrAlreadyCommitted  = errors.New("unit of work already committed")
	ErrAlreadyRolledBack = errors.New("unit of work already rolled back")
)

type FarmUnitOfWork struct {
	tx  *sql.Tx
	ctx context.Context
	bus messaging.EventBus

	committed  bool
	rolledBack bool

	mu         sync.Mutex
	aggregates []aggregate.Aggregate
}

func (uow *FarmUnitOfWork) RegisterAggregate(agg aggregate.Aggregate) {
	uow.mu.Lock()
	defer uow.mu.Unlock()

	uow.aggregates = append(uow.aggregates, agg)
}

func (uow *FarmUnitOfWork) Commit() error {
	uow.mu.Lock()
	defer uow.mu.Unlock()

	if uow.committed {
		return ErrAlreadyCommitted
	}
	if uow.rolledBack {
		return ErrAlreadyRolledBack
	}

	if err := uow.tx.Commit(); err != nil {
		return err
	}

	return uow.dispatchEvents()
}

func (uow *FarmUnitOfWork) Rollback() error {
	uow.mu.Lock()
	defer uow.mu.Unlock()

	if uow.committed {
		return ErrAlreadyCommitted
	}
	if uow.rolledBack {
		return ErrAlreadyRolledBack
	}
	uow.rolledBack = true
	return uow.tx.Rollback()
}

func (uow *FarmUnitOfWork) dispatchEvents() error {
	var allEvents []event.DomainEvent

	for _, agg := range uow.aggregates {
		allEvents = append(allEvents, agg.PullEvents()...)
	}

	if len(allEvents) == 0 {
		return nil
	}

	return uow.bus.Publish(repository.WithUnitOfWork(uow.ctx, uow), allEvents)
}

func (uow *FarmUnitOfWork) Execute(ctx context.Context, fn func(repository.RepositoryProvider) error) error {
	// Создаем провайдер для этой транзакции
	provider := NewFarmProvider(uow.tx)

	// Выполняем бизнес-логику
	if err := fn(provider); err != nil {
		// В случае ошибки — откат
		if rbErr := uow.Rollback(); rbErr != nil {
			return fmt.Errorf("rollback error: %v, original error: %w", rbErr, err)
		}
		return err
	}

	// Пробуем закоммитить
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func NewFarmUnitOfWork(ctx context.Context, tx *sql.Tx, bus messaging.EventBus) repository.UnitOfWork {
	return &FarmUnitOfWork{
		tx:  tx,
		ctx: ctx,
		bus: bus,
	}
}
