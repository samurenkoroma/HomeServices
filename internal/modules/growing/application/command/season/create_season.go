package season

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/season"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/inmemory"
	"time"
)

type CreateSeasonCmd struct {
	StartDate   string `json:"startDate,format:date" validate:"required"`
	EndDate     string `json:"endDate,format:date" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Status      string `json:"status" validate:"required"`
	Description string `json:"description"`
}

func (h *SeasonHandler) Create(ctx context.Context, cmd any) (any, error) {
	c, ok := cmd.(*CreateSeasonCmd)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}
	startDate, err := time.Parse(time.RFC3339, c.StartDate)
	if err != nil {
		return nil, err
	}
	endDate, err := time.Parse(time.RFC3339, c.EndDate)
	if err != nil {
		return nil, err
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return uow.Execute(ctx, inmemory.NewRedisGrowingProvider, func(provider repository.RepositoryProvider) (any, error) {
		orgId, ok := ctx.Value("organization_id").(string)
		growingProvider, ok := provider.(*inmemory.RedisGrowingProvider)
		if !ok {
			return nil, fmt.Errorf("expected FarmProvider, got %T", provider)
		}
		newSeason, err := season.NewSeason(c.Name, startDate, endDate, season.SeasonStatus(c.Status), orgId, c.Description)
		if err != nil {
			return nil, err
		}
		err = growingProvider.Seasons().Save(ctx, newSeason)
		if err != nil {
			return nil, err
		}

		uow.RegisterAggregate(newSeason)

		return nil, nil
	})
}
