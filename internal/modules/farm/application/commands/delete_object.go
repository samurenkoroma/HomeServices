package commands

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/modules/farm/infrastructure/persistence/postgres"

	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
)

type DeleteFarmObjectCommand struct {
	ID string `json:"id" validate:"required"`
}
type deleteFarmObjectHandler struct {
	uowFactory repository.Factory
}

func NewDeleteFarmObjectHandler(uowFactory repository.Factory) command.Handler {
	return &deleteFarmObjectHandler{uowFactory: uowFactory}
}

// Handle обрабатывает команду
func (h *deleteFarmObjectHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(*DeleteFarmObjectCommand)
	if !ok {
		return command.ErrInvalidCommandType
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = uow.Execute(ctx, postgres.NewPostgresFarmProvider, func(provider repository.RepositoryProvider) error {
		farmProvider, ok := provider.(*postgres.FarmProvider)
		if !ok {
			return fmt.Errorf("invalid provider type")
		}

		// Получаем объект
		obj, err := farmProvider.Objects().FindByID(ctx, physicalobject.PhysicalObjectID(c.ID))
		if err != nil {
			return fmt.Errorf("failed to find physical object: %w", err)
		}
		if err := obj.Delete(); err != nil {
			return fmt.Errorf("Ошибка удаления %w", err)
		}

		// Сохраняем
		if err := farmProvider.Objects().Delete(ctx, obj); err != nil {
			return fmt.Errorf("failed to save physical object: %w", err)
		}

		uow.RegisterAggregate(obj)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
