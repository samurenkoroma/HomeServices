package commands

import (
	"context"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/farm/domain/greenhouse"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
	"samurenkoroma/services/internal/modules/farm/infrastructure/persistence/postgres"
)

type CreateGreenhouseCmd struct {
	Name   string          `json:"name"`
	Length float64         `json:"length"`
	Width  float64         `json:"width"`
	Geom   spatial.GeoJSON `json:"geom"`
}

type createGreenhouseHandler struct {
	uowFactory repository.Factory
}

func NewCreateGreenhouseHandler(uowFactory repository.Factory) command.Handler {
	return &createGreenhouseHandler{uowFactory: uowFactory}
}

func (h *createGreenhouseHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(CreateGreenhouseCmd)
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

		dimension, err := valueobject.NewDimension(c.Length, c.Width)
		if err != nil {
			return err
		}

		// 1. Создаем теплицу
		newGreenhouse := greenhouse.NewGreenhouse(c.Name, dimension, c.Geom)

		// 2. Сохраняем поле
		if err := farmProvider.Greenhouses().Save(ctx, newGreenhouse); err != nil {
			return fmt.Errorf("failed to save field: %w", err)
		}

		uow.RegisterAggregate(newGreenhouse)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
