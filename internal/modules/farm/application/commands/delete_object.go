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

func (h *FarmObjectHandler) Delete(ctx context.Context, cmd any) (any, error) {
	c, ok := cmd.(*DeleteFarmObjectCommand)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return uow.Execute(ctx, postgres.NewPostgresFarmProvider, func(provider repository.RepositoryProvider) (any, error) {
		farmProvider, ok := provider.(*postgres.FarmProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}

		// Получаем объект
		obj, err := farmProvider.Objects().FindByID(ctx, physicalobject.PhysicalObjectID(c.ID))
		if err != nil {
			return nil, err
		}
		if err := obj.Delete(); err != nil {
			return nil, err
		}

		// Сохраняем
		if err := farmProvider.Objects().Delete(ctx, obj); err != nil {
			return nil, err
		}

		uow.RegisterAggregate(obj)
		return nil, nil
	})
}
