package repository

import (
	"context"
)

type ctxKey struct{}

func WithUnitOfWork(ctx context.Context, uow UnitOfWork) context.Context {
	return context.WithValue(ctx, ctxKey{}, uow)
}

func FromContext(ctx context.Context) (UnitOfWork, bool) {
	uow, ok := ctx.Value(ctxKey{}).(UnitOfWork)
	return uow, ok
}
