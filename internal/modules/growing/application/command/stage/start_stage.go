package stage

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/inmemory"
)

type StartStageCmd struct {
	PlanID      string `json:"planId"`
	StageID     string `json:"stageId"`
	CurrentBBCH int    `json:"currentBbch"` // текущая BBCH фаза (из phenology)
}

func (h *StageHandler) Start(ctx context.Context, cmd any) (any, error) {
	c, ok := cmd.(*StartStageCmd)
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

		if err := plan.StartStage(c.StageID, c.CurrentBBCH); err != nil {
			return nil, err
		}

		if err := growingProvider.CropPlans().Save(ctx, plan); err != nil {
			return nil, err
		}

		uow.RegisterAggregate(plan)
		return nil, nil
	})
}
