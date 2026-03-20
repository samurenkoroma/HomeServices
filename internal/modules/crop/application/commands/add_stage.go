package commands

import (
	"context"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/crop/domain/cropplan"
	"samurenkoroma/services/internal/modules/crop/domain/valueobject"
	"samurenkoroma/services/internal/modules/crop/infrastructure/persistence/postgres"
)

type AddStageCommand struct {
	PlanID      string  `json:"plan_id" validate:"required"`
	Order       int     `json:"order" validate:"required,gt=0"`
	Name        string  `json:"name" validate:"required"`
	Duration    int     `json:"duration" validate:"required,gt=0"`
	MinTemp     float64 `json:"min_temp"`
	MaxTemp     float64 `json:"max_temp"`
	OptimalTemp float64 `json:"optimal_temp"`
	WaterPerDay float64 `json:"water_per_day"`
}

type AddStageHandler struct {
	uowFactory repository.Factory
}

func (h *AddStageHandler) Handle(ctx context.Context, cmd AddStageCommand) error {
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return err
	}

	return uow.Execute(ctx, func(provider repository.RepositoryProvider) error {
		cropProvider := provider.(*postgres.CropProvider)

		// Получаем план
		plan, err := cropProvider.CropPlans().GetByID(ctx, cropplan.PlanID(cmd.PlanID))
		if err != nil {
			return err
		}

		// Создаем этап
		stage, err := cropplan.NewGrowthStage(
			cmd.Order,
			cmd.Name,
			valueobject.Duration(cmd.Duration),
			cmd.MinTemp,
			cmd.MaxTemp,
			cmd.OptimalTemp,
			cmd.WaterPerDay,
		)
		if err != nil {
			return err
		}

		// Добавляем
		if err := plan.AddStage(stage); err != nil {
			return err
		}

		// Сохраняем
		if err := cropProvider.CropPlans().Save(ctx, plan); err != nil {
			return err
		}

		uow.RegisterAggregate(plan)
		return nil
	})
}
