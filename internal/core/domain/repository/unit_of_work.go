package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/event"
	"samurenkoroma/services/internal/core/port/messaging"
	"sync"
)

var (
	ErrAlreadyCommitted  = errors.New("unit of work already committed")
	ErrAlreadyRolledBack = errors.New("unit of work already rolled back")
)

type UnitOfWork interface {
	// Execute выполняет функцию в рамках транзакции
	Execute(ctx context.Context, build func(tx *sql.Tx) RepositoryProvider, fn func(RepositoryProvider) (any, error)) (any, error)
	Tx() *sql.Tx
	RegisterAggregate(agg aggregate.Aggregate)
	Commit() error
	Rollback() error
}

type unitOfWork struct {
	committed  bool
	rolledBack bool

	mu         sync.Mutex
	aggregates []aggregate.Aggregate
	tx         *sql.Tx
	factory    Factory
	ctx        context.Context
	bus        messaging.EventBus
}

func NewUnitOfWork(ctx context.Context, tx *sql.Tx, factory Factory, bus messaging.EventBus) UnitOfWork {
	return &unitOfWork{
		tx:      tx,
		ctx:     ctx,
		bus:     bus,
		factory: factory,
	}
}

func (uow *unitOfWork) Tx() *sql.Tx {
	return uow.tx
}

func (uow *unitOfWork) Execute(ctx context.Context, build func(tx *sql.Tx) RepositoryProvider, fn func(RepositoryProvider) (any, error)) (any, error) {
	// Создаем провайдер для этой транзакции
	provider := build(uow.tx)

	// Выполняем бизнес-логику
	data, err := fn(provider)
	if err != nil {
		// В случае ошибки — откат
		if rbErr := uow.Rollback(); rbErr != nil {
			return nil, fmt.Errorf("rollback error: %v, original error: %w", rbErr, err)
		}
		return nil, err
	}

	// Пробуем закоммитить
	if err := uow.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return data, nil
}

func (uow *unitOfWork) RegisterAggregate(agg aggregate.Aggregate) {
	uow.mu.Lock()
	defer uow.mu.Unlock()

	uow.aggregates = append(uow.aggregates, agg)
}

func (uow *unitOfWork) Commit() error {
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

func (uow *unitOfWork) Rollback() error {
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

func (uow *unitOfWork) dispatchEvents() error {
	var allEvents []event.DomainEvent

	for _, agg := range uow.aggregates {
		allEvents = append(allEvents, agg.PullEvents()...)
	}

	if len(allEvents) == 0 {
		return nil
	}

	return uow.bus.Publish(WithFactory(uow.ctx, uow.factory), allEvents)
}
