package repository

import (
	"context"
)

type ctxKey struct{}
type factoryKey struct{}

func WithUnitOfWork(ctx context.Context, uow UnitOfWork) context.Context {
	return context.WithValue(ctx, ctxKey{}, uow)
}
func FromContext(ctx context.Context) (UnitOfWork, bool) {
	uow, ok := ctx.Value(ctxKey{}).(UnitOfWork)
	return uow, ok
}

func FactoryFromContext(ctx context.Context) (Factory, bool) {
	factory, ok := ctx.Value(factoryKey{}).(Factory)
	return factory, ok
}

func WithFactory(ctx context.Context, factory Factory) context.Context {
	return context.WithValue(ctx, factoryKey{}, factory)
}
