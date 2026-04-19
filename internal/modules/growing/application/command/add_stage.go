package command

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/inmemory"
)

// AddStageHandler команда добавления этапа в план
type addStageHandler struct {
	uowFactory repository.Factory
}

func (h *addStageHandler) Name() string {
	return "AddStage"
}

// AddStageCmd структура команды
type AddStageCmd struct {
	PlanID      string `json:"planId"`
	Name        string `json:"name"`
	Type        string `json:"type"` // soil_preparation, sowing, etc.
	Description string `json:"description"`
	BBCHStart   int    `json:"bbchStart"`
	BBCHEnd     int    `json:"bbchEnd"`
	Order       int    `json:"order"`
}

func NewAddStageHandler(uowFactory repository.Factory) command.Handler {
	return &addStageHandler{
		uowFactory: uowFactory,
	}
}

// Handle выполняет команду
func (h *addStageHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(*AddStageCmd)
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

		// Создаем этап
		stage, err := cropplan.NewStage(
			types.NewUUID(),
			c.PlanID,
			cropplan.StageType(c.Type),
			c.Name,
			c.Order,
			c.BBCHStart,
			c.BBCHEnd,
		)
		if err != nil {
			return fmt.Errorf("failed to create stage: %w", err)
		}
		stage.Description = c.Description

		// Добавляем этап в план
		if err := plan.AddStage(*stage); err != nil {
			return fmt.Errorf("failed to add stage to plan: %w", err)
		}

		// Сохраняем изменения
		if err := growingProvider.CropPlans().Update(ctx, plan); err != nil {
			return fmt.Errorf("failed to update plan: %w", err)
		}

		uow.RegisterAggregate(plan)

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
