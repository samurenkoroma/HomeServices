package command

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/inmemory"
)

// SkipStageHandler команда пропуска этапа
type skipStageHandler struct {
	uowFactory repository.Factory
}

func (h *skipStageHandler) Name() string {
	return "SkipStage"
}

// SkipStageCmd структура команды
type SkipStageCmd struct {
	PlanID  string `json:"plan_id"`
	StageID string `json:"stage_id"`
	Reason  string `json:"reason"`
}

func NewSkipStageCommand(uowFactory repository.Factory) command.Handler {
	return &skipStageHandler{
		uowFactory: uowFactory,
	}
}

// Handle выполняет команду
func (h *skipStageHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(*SkipStageCmd)
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

		// Пропускаем этап
		if err := plan.SkipStage(c.StageID); err != nil {
			return fmt.Errorf("failed to skip stage: %w", err)
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
