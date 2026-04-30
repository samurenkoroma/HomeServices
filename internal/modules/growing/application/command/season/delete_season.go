package season

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/season"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/inmemory"
)

type DeleteSeasonCmd struct {
	SeasonId    string `json:"seasonId"`
	Permanently bool   `json:"permanently"`
}

func (h *SeasonHandler) Delete(ctx context.Context, cmd any) (any, error) {
	c, ok := cmd.(*DeleteSeasonCmd)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return uow.Execute(ctx, inmemory.NewRedisGrowingProvider, func(provider repository.RepositoryProvider) (any, error) {
		// Приводим провайдер к нужному типу
		growingProvider, ok := provider.(*inmemory.RedisGrowingProvider)
		if !ok {
			return nil, fmt.Errorf("expected FarmProvider, got %T", provider)
		}

		obj, err := growingProvider.Seasons().FindByID(ctx, season.SeasonID(c.SeasonId))
		if err != nil {
			return nil, err
		}

		//TODO сделать проверки на наличие записей сезона

		if c.Permanently {
			err = growingProvider.Seasons().Delete(ctx, season.SeasonID(c.SeasonId))
		} else {
			obj.Delete()
			err = growingProvider.Seasons().Save(ctx, obj)
		}
		if err != nil {
			return nil, err
		}

		uow.RegisterAggregate(obj)

		return nil, nil
	})
}
