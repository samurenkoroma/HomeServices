package commands

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/season"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
)

type ActivateSeasonCmd struct {
	Id string `json:"season_id" validate:"required"`
}
type activateSeasonHandler struct {
	uowFactory repository.Factory
}

func (h *activateSeasonHandler) Handle(ctx context.Context, command any) error {
	cmd, ok := command.(*ActivateSeasonCmd)
	if !ok {
		return errors.New("invalid command type")
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return err
	}

	return uow.Execute(ctx, postgres.NewGrowingProvider, func(provider repository.RepositoryProvider) error {
		growingProvider := provider.(*postgres.GrowingProvider)

		seasonObj, err2 := growingProvider.Seasons().FindByID(ctx, season.SeasonID(cmd.Id))
		if err2 != nil {
			return err2
		}
		err := seasonObj.Activate()
		if err != nil {
			return err
		}
		if err := growingProvider.Seasons().Save(ctx, seasonObj); err != nil {
			return err
		}
		uow.RegisterAggregate(seasonObj)
		return nil
	})
}

func (h *activateSeasonHandler) Name() string {
	return "ActivateSeason"
}

func NewActivateSeasonCommand(uowFactory repository.Factory) command.Handler {
	return &activateSeasonHandler{
		uowFactory: uowFactory,
	}
}
