package commands

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/farm/application"
	"samurenkoroma/services/internal/modules/farm/domain/field_block"
	"samurenkoroma/services/internal/modules/farm/infrastructure/persistence/postgres"
)

type CreateFieldBlockCmd struct {
	Name     string          `json:"name"`
	Geom     spatial.GeoJSON `json:"geom"`
	ParentId string          `json:"parent_id"`
}

type createFieldBlockHandler struct {
	uowFactory repository.Factory
}

func NewCreateFieldBlockHandler(uowFactory repository.Factory) command.Handler {
	return &createFieldBlockHandler{uowFactory: uowFactory}
}

func (h *createFieldBlockHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(CreateFieldBlockCmd)
	if !ok {
		return application.ErrInvalidCommandType
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
		newFieldBlock := field_block.NewFieldBlock(types.FieldId(c.ParentId), c.Name, c.Geom)

		// 2. Сохраняем поле
		if err := farmProvider.FieldBlocks().Save(ctx, newFieldBlock); err != nil {
			return fmt.Errorf("failed to save field: %w", err)
		}

		uow.RegisterAggregate(newFieldBlock)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
