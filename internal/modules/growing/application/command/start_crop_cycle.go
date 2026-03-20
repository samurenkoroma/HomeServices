package command

import (
	"context"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/cropcycle"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
)

type StartCropCycleCommand struct {
	AreaID     string `json:"area_id" validate:"required"`
	SeasonID   string `json:"season_id" validate:"required"`
	CropPlanID string `json:"crop_plan_id" validate:"required"`
}

type StartCropCycleHandler struct {
	uowFactory repository.Factory
}

func NewStartCropCycleHandler(uowFactory repository.Factory) *StartCropCycleHandler {
	return &StartCropCycleHandler{uowFactory: uowFactory}
}

func (h *StartCropCycleHandler) Handle(ctx context.Context, cmd StartCropCycleCommand) error {
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return err
	}

	return uow.Execute(ctx, func(provider repository.RepositoryProvider) error {
		growingProvider := provider.(*postgres.GrowingProvider)

		// Проверяем, что место настроено на этот сезон
		area, err := growingProvider.CultivationAreas().GetByID(ctx, cmd.AreaID)
		if err != nil {
			return err
		}

		if !area.IsConfiguredForSeason(cmd.SeasonID) {
			return cultivationarea.ErrAreaNotConfiguredForSeason
		}

		// Проверяем, что план культуры соответствует конфигурации
		configuredPlan, err := area.GetCropPlanForSeason(cmd.SeasonID)
		if err == nil && configuredPlan != cmd.CropPlanID {
			return cultivationarea.ErrCropPlanMismatch
		}

		// Создаем цикл
		cycle := cropcycle.NewCropCycle(
			cmd.AreaID,
			cmd.SeasonID,
			cmd.CropPlanID,
		)

		if err := cycle.Start(); err != nil {
			return err
		}

		// Сохраняем
		if err := growingProvider.CropCycles().Save(ctx, cycle); err != nil {
			return err
		}

		uow.RegisterAggregate(cycle)
		return nil
	})
}
