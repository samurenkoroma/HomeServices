package commands

import (
	"context"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/season"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
	"time"
)

type CreateSeasonCmd struct {
	Name        string `json:"name" validate:"required"`
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date" validate:"required"`
	Description string `json:"description" validate:"required"`
	CreatedBy   string `json:"created_by" validate:"required"`
}
type createSeasonHandler struct {
	uowFactory repository.Factory
}

func (h *createSeasonHandler) Handle(ctx context.Context, command any) error {
	cmd, ok := command.(*CreateSeasonCmd)
	if !ok {
		return errors.New("invalid command type")
	}
	// Парсим даты
	startDate, err := time.Parse("2006-01-02", cmd.StartDate)
	if err != nil {
		return fmt.Errorf("invalid start_date format, expected YYYY-MM-DD: %w", err)
	}

	endDate, err := time.Parse("2006-01-02", cmd.EndDate)
	if err != nil {
		return fmt.Errorf("invalid end_date format, expected YYYY-MM-DD: %w", err)
	}

	// Валидация
	if startDate.After(endDate) {
		return season.ErrInvalidPeriod
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return err
	}

	return uow.Execute(ctx, postgres.NewGrowingProvider, func(provider repository.RepositoryProvider) error {
		growingProvider := provider.(*postgres.GrowingProvider)

		newSeason, err := season.NewSeason(cmd.Name, startDate, endDate, cmd.CreatedBy, cmd.Description)
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
