package command

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
)

// ActivateCropPlanHandler команда активации плана
type activateCropPlanHandler struct {
	uowFactory repository.Factory
}

func (h *activateCropPlanHandler) Name() string {
	return "ActivateCropPlan"
}

// ActivateCropPlanCmd структура команды
type ActivateCropPlanCmd struct {
	PlanID string `json:"planId"`
}

func NewActivateCropPlanCmd(uowFactory repository.Factory) command.Handler {
	return &activateCropPlanHandler{uowFactory: uowFactory}
}

func (h *activateCropPlanHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(*ActivateCropPlanCmd)
	if !ok {
		return command.ErrInvalidCommandType
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = uow.Execute(ctx, postgres.NewPostgresGrowingProvider, func(provider repository.RepositoryProvider) error {
		// Приводим провайдер к нужному типу
		growingProvider, ok := provider.(*postgres.PostgresGrowingProvider)
		if !ok {
			return fmt.Errorf("expected FarmProvider, got %T", provider)
		}

		plan, err := growingProvider.CropPlans().FindByID(ctx, c.PlanID)
		if err != nil {
			return err
		}

		if err := plan.Activate(); err != nil {
			return err
		}

		if err := growingProvider.CropPlans().Save(ctx, plan); err != nil {
			return err
		}

		uow.RegisterAggregate(plan)

		return nil
	})

	if err != nil {
		return err
	}
	return nil

}
