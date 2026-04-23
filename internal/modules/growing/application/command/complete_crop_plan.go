package command

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/inmemory"
)

// CompleteCropPlanHandler команда завершения плана (сбор урожая)
type completeCropPlanHandler struct {
	uowFactory repository.Factory
}

type CompleteCropPlanCmd struct {
	PlanID    string  `json:"planId"`
	HarvestKg float64 `json:"harvestKg"`
}

func NewCompleteCropPlanHandler(uowFactory repository.Factory) command.Handler {
	return &completeCropPlanHandler{
		uowFactory: uowFactory,
	}
}

func (h *completeCropPlanHandler) Handle(ctx context.Context, cmd any) (any, error) {
	c, ok := cmd.(*CompleteCropPlanCmd)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return nil, err
	}

	err = uow.Execute(ctx, inmemory.NewRedisGrowingProvider, func(provider repository.RepositoryProvider) error {
		// Приводим провайдер к нужному типу
		growingProvider, ok := provider.(*inmemory.RedisGrowingProvider)
		if !ok {
			return fmt.Errorf("expected FarmProvider, got %T", provider)
		}
		// Получаем план
		plan, err := growingProvider.CropPlans().FindByID(ctx, c.PlanID)
		if err != nil {
			return fmt.Errorf("failed to find plan: %w", err)
		}

		if err := plan.Complete(c.HarvestKg); err != nil {
			return err
		}

		if err := growingProvider.CropPlans().Save(ctx, plan); err != nil {
			return err
		}

		uow.RegisterAggregate(plan)
		return nil

	})
	return nil, err
}
