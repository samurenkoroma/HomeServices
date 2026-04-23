package commands

import (
	"context"
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
func (h *deleteFarmObjectHandler) Handle(ctx context.Context, cmd any) (any, error) {
	c, ok := cmd.(*DeleteFarmObjectCommand)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return nil, err
	}

	err = uow.Execute(ctx, postgres.NewPostgresFarmProvider, func(provider repository.RepositoryProvider) error {
		farmProvider, ok := provider.(*postgres.FarmProvider)
		if !ok {
			return repository.ErrInvalidProviderType
		}

		// Получаем объект
		obj, err := farmProvider.Objects().FindByID(ctx, physicalobject.PhysicalObjectID(c.ID))
		if err != nil {
			return err
		}
		if err := obj.Delete(); err != nil {
			return err
		}

		// Сохраняем
		if err := farmProvider.Objects().Delete(ctx, obj); err != nil {
			return err
		}

		uow.RegisterAggregate(obj)
		return nil
	})

	return nil, err
}
