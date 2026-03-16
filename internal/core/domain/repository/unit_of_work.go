package repository

import (
	"context"
	"samurenkoroma/services/internal/core/domain/aggregate"
)

type UnitOfWork interface {
	// Execute выполняет функцию в рамках транзакции
	Execute(ctx context.Context, fn func(RepositoryProvider) error) error

	RegisterAggregate(agg aggregate.Aggregate)
	Commit() error
	Rollback() error
}
