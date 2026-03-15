package repository

import (
	"context"
)

type UnitOfWork interface {
	// Execute выполняет функцию в рамках транзакции
	Execute(ctx context.Context, fn func(RepositoryProvider) error) error
}
