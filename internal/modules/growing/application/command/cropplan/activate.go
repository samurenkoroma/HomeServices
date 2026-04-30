package cropplan

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
)

type ActivateCropPlanCmd struct {
	PlanID string `json:"planId"`
}

func (h *CropPlanHandler) Activate(ctx context.Context, cmd any) (any, error) {
	c, ok := cmd.(*ActivateCropPlanCmd)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return uow.Execute(ctx, postgres.NewPostgresGrowingProvider, func(provider repository.RepositoryProvider) (any, error) {
		// Приводим провайдер к нужному типу
		growingProvider, ok := provider.(*postgres.PostgresGrowingProvider)
		if !ok {
			return nil, fmt.Errorf("expected FarmProvider, got %T", provider)
		}

		plan, err := growingProvider.CropPlans().FindByID(ctx, c.PlanID)
		if err != nil {
			return nil, err
		}

		if err := plan.Activate(); err != nil {
			return nil, err
		}

		if err := growingProvider.CropPlans().Save(ctx, plan); err != nil {
			return nil, err
		}

		uow.RegisterAggregate(plan)

		return nil, nil
	})
}
