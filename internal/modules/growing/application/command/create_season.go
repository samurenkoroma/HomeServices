package command

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/growing/domain/season"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/inmemory"
	"time"
)

type createSeasonHandler struct {
	uowFactory repository.Factory
}

func (h *createSeasonHandler) Name() string {
	return "CreateSeason"
}

func NewCreateSeasonHandler(uowFactory repository.Factory) command.Handler {
	return &createSeasonHandler{uowFactory: uowFactory}
}

type CreateSeasonCmd struct {
	StartDate   string `json:"startDate,format:date" validate:"required"`
	EndDate     string `json:"endDate"  validate:"required"`
	Name        string `json:"name"  validate:"required"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func (h *createSeasonHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(*CreateSeasonCmd)
	if !ok {
		return command.ErrInvalidCommandType
	}
	startDate, err := time.Parse(time.RFC3339, c.StartDate)
	if err != nil {
		return fmt.Errorf("invalid start date %s", err)
	}
	endDate, err := time.Parse(time.RFC3339, c.StartDate)
	if err != nil {
		return fmt.Errorf("invalid start date %s", err)
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
		// Получаем сорт из каталога
		newSeason, err := season.NewSeason(c.Name, startDate, endDate, season.SeasonStatus(c.Status), types.NewUUID(), c.Description)
		if err != nil {
			return fmt.Errorf("failed to create season: %w", err)
		}
		err = growingProvider.Seasons().Save(ctx, newSeason)
		if err != nil {
			return fmt.Errorf("variety not found: %w", err)
		}

		uow.RegisterAggregate(newSeason)

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
