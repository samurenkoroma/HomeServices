package command

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/inmemory"
)

type startStageHandler struct {
	uowFactory repository.Factory
}

func (h *startStageHandler) Name() string {
	return "StartStage"
}

// StartStageCmd структура команды
type StartStageCmd struct {
	PlanID      string `json:"planId"`
	StageID     string `json:"stageId"`
	CurrentBBCH int    `json:"currentBbch"` // текущая BBCH фаза (из phenology)
}

func NewStartStageCommand(uowFactory repository.Factory) command.Handler {
	return &startStageHandler{
		uowFactory: uowFactory,
	}
}

// Handle выполняет команду
func (h *startStageHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(*StartStageCmd)
	if !ok {
		return command.ErrInvalidCommandType
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
		// Получаем план
		plan, err := growingProvider.CropPlans().FindByID(ctx, c.PlanID)
		if err != nil {
			return fmt.Errorf("failed to find plan: %w", err)
		}

		if err := plan.StartStage(c.StageID, c.CurrentBBCH); err != nil {
			return fmt.Errorf("failed to complete stage: %w", err)
		}

		if err := growingProvider.CropPlans().Update(ctx, plan); err != nil {
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
