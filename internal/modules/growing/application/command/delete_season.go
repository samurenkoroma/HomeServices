package command

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/season"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/inmemory"
)

type deleteSeasonHandler struct {
	uowFactory repository.Factory
}

func (h *deleteSeasonHandler) Name() string {
	return "DeleteSeason"
}

func NewDeleteSeasonHandler(uowFactory repository.Factory) command.Handler {
	return &deleteSeasonHandler{uowFactory: uowFactory}
}

type DeleteSeasonCmd struct {
	SeasonId    string `json:"seasonId"`
	Permanently bool   `json:"permanently"`
}

func (h *deleteSeasonHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(*DeleteSeasonCmd)
	if !ok {
		return command.ErrInvalidCommandType
	}
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = uow.Execute(ctx, inmemory.NewRedisGrowingProvider, func(provider repository.RepositoryProvider) error {
		// Приводим провайдер к нужному типу
		growingProvider, ok := provider.(*inmemory.RedisGrowingProvider)
		if !ok {
			return fmt.Errorf("expected FarmProvider, got %T", provider)
		}

		obj, err := growingProvider.Seasons().FindByID(ctx, season.SeasonID(c.SeasonId))
		if err != nil {
			return fmt.Errorf("failed to find season: %w", err)
		}

		//TODO сделать проверки на наличие записей сезона

		if c.Permanently {
			err = growingProvider.Seasons().Delete(ctx, season.SeasonID(c.SeasonId))
		} else {
			obj.Delete()
			err = growingProvider.Seasons().Save(ctx, obj)
		}
		if err != nil {
			return fmt.Errorf("variety not found: %w", err)
		}

		uow.RegisterAggregate(obj)

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
