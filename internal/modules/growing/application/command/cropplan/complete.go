package cropplan

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/inmemory"
)

type CompleteCropPlanCmd struct {
	PlanID    string  `json:"planId"`
	HarvestKg float64 `json:"harvestKg"`
}

func (h *CropPlanHandler) Complete(ctx context.Context, cmd any) (any, error) {
	c, ok := cmd.(*CompleteCropPlanCmd)
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
		// Получаем план
		plan, err := growingProvider.CropPlans().FindByID(ctx, c.PlanID)
		if err != nil {
			return nil, err
		}

		if err := plan.Complete(c.HarvestKg); err != nil {
			return nil, err
		}

		if err := growingProvider.CropPlans().Save(ctx, plan); err != nil {
			return nil, err
		}

		uow.RegisterAggregate(plan)
		return nil, nil
	})
}
