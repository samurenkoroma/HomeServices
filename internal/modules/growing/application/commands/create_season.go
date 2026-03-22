package commands

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/season"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
	"time"
)

type CreateSeasonCmd struct {
	Name        string    `json:"name" validate:"required"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CreatedBy   string    `json:"created_by" validate:"required"`
}
type createSeasonHandler struct {
	uowFactory repository.Factory
}

func (h *createSeasonHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(CreateSeasonCmd)
	if !ok {
		return errors.New("invalid command type")
	}
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return err
	}

	return uow.Execute(ctx, postgres.NewGrowingProvider, func(provider repository.RepositoryProvider) error {
		growingProvider := provider.(*postgres.GrowingProvider)

		newSeason, err := season.NewSeason(c.Name, c.StartDate, c.EndDate, c.CreatedBy, c.Description)
		if err != nil {
			return err
		}

		if err := growingProvider.Seasons().Save(ctx, newSeason); err != nil {
			return err
		}
		uow.RegisterAggregate(newSeason)
		return nil
	})
}

func (h *createSeasonHandler) Name() string {
	return "CreateSeason"
}

func NewCreateSeasonCommand(uowFactory repository.Factory) command.Handler {
	return &createSeasonHandler{
		uowFactory: uowFactory,
	}
}
