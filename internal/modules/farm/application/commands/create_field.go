package commands

import (
	"context"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/farm/domain/field"
	"samurenkoroma/services/internal/modules/farm/infrastructure/persistence/postgres"
)

type CreateFieldCmd struct {
	Name string          `json:"name"`
	Geom spatial.GeoJSON `json:"geom"`
}
type createFieldHandler struct {
	uowFactory repository.Factory
}

func NewCreateFieldHandler(uowFactory repository.Factory) command.Handler {
	return &createFieldHandler{uowFactory: uowFactory}
}

func (h *createFieldHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(CreateFieldCmd)
	if !ok {
		return errors.New("invalid command type")
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return err
	}

	err = uow.Execute(ctx, func(provider repository.RepositoryProvider) error {
		// Приводим провайдер к нужному типу
		farmProvider, ok := provider.(*postgres.FarmProvider)
		if !ok {
			return fmt.Errorf("expected FarmProvider, got %T", provider)
		}

		// 1. Создаем поле
		newField := field.NewField(c.Name, c.Geom)

		// 2. Сохраняем поле
		if err := farmProvider.Fields().Save(ctx, newField); err != nil {
			return fmt.Errorf("failed to save field: %w", err)
		}

		uow.RegisterAggregate(newField)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
