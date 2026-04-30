package stage

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/inmemory"
)

type AddStageCmd struct {
	PlanID      string `json:"planId"`
	Name        string `json:"name"`
	Type        string `json:"type"` // soil_preparation, sowing, etc.
	Description string `json:"description"`
	BBCHStart   int    `json:"bbchStart"`
	BBCHEnd     int    `json:"bbchEnd"`
	Order       int    `json:"order"`
}

func (h *StageHandler) Add(ctx context.Context, cmd any) (any, error) {
	c, ok := cmd.(*AddStageCmd)
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
			return nil, err
		}
		stage.Description = c.Description

		// Добавляем этап в план
		if err := plan.AddStage(*stage); err != nil {
			return nil, err
		}

		// Сохраняем изменения
		if err := growingProvider.CropPlans().Save(ctx, plan); err != nil {
			return nil, err
		}
		uow.RegisterAggregate(plan)

		return nil, nil
	})
}
