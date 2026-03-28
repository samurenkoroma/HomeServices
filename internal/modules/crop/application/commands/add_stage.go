package commands

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/crop/domain/cropplan"
	"samurenkoroma/services/internal/modules/crop/domain/valueobject"
	"samurenkoroma/services/internal/modules/crop/infrastructure/persistence/postgres"
)

type AddStageCmd struct {
	PlanID      string  `json:"plan_id" validate:"required"`
	Order       int     `json:"order" validate:"required,gt=0"`    // порядковый номер этапа
	Name        string  `json:"name" validate:"required"`          // название этапа
	Duration    int     `json:"duration" validate:"required,gt=0"` // длительность в днях
	MinTemp     float64 `json:"min_temp"`                          // мин. температура
	MaxTemp     float64 `json:"max_temp"`                          // макс. температура
	OptimalTemp float64 `json:"optimal_temp"`                      // оптимальная температура
	WaterPerDay float64 `json:"water_per_day"`                     // полив л/м² в день
}
type addStageHandler struct {
	uowFactory repository.Factory
}

func (h *addStageHandler) Name() string {
	return "AddStage"
}

func NewAddStageHandler(uowFactory repository.Factory) command.Handler {
	return &addStageHandler{
		uowFactory: uowFactory,
	}
}
func (h *addStageHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(AddStageCmd)
	if !ok {
		return command.ErrInvalidCommandType
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	return uow.Execute(ctx, postgres.NewCropProvider, func(provider repository.RepositoryProvider) error {
		cropProvider := provider.(*postgres.CropProvider)

		// Получаем план
		plan, err := cropProvider.CropPlans().GetByID(ctx, cropplan.PlanID(c.PlanID))
		if err != nil {
			return err
		}

		// Создаем этап
		stage, err := cropplan.NewGrowthStage(
			c.Order,
			c.Name,
			valueobject.Duration(c.Duration),
			c.MinTemp,
			c.MaxTemp,
			c.OptimalTemp,
			c.WaterPerDay,
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
