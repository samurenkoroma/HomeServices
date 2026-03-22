package commands

import (
	"context"
	"errors"
	"fmt"
	"log"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"

	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/cropcycle"
	"samurenkoroma/services/internal/modules/growing/domain/croptemplate"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
	"samurenkoroma/services/internal/modules/growing/domain/season"
)

// StartCropCycleCmd — команда начала цикла
type StartCropCycleCmd struct {
	AreaID          string `json:"area_id" validate:"required"`
	SeasonID        string `json:"season_id" validate:"required"`
	CropPlanID      string `json:"crop_plan_id" validate:"required"`
	CropPlanVersion int    `json:"crop_plan_version" validate:"required,gt=0"`
}

// startCropCycleHandler — обработчик начала цикла
type startCropCycleHandler struct {
	uowFactory repository.Factory
}

func (h *startCropCycleHandler) Name() string {
	return "StartCropCycle"
}

func NewStartCropCycleHandler(uowFactory repository.Factory) command.Handler {
	return &startCropCycleHandler{
		uowFactory: uowFactory,
	}
}

func (h *startCropCycleHandler) Handle(ctx context.Context, command any) error {
	cmd, ok := command.(StartCropCycleCmd)
	if !ok {
		return errors.New("invalid command type")
	}
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	var cycle *cropcycle.CropCycle

	err = uow.Execute(ctx, postgres.NewGrowingProvider, func(provider repository.RepositoryProvider) error {
		growingProvider, ok := provider.(*postgres.GrowingProvider)
		if !ok {
			return fmt.Errorf("invalid provider type")
		}

		// 1. Проверяем сезон
		s, err := growingProvider.Seasons().FindByID(ctx, season.SeasonID(cmd.SeasonID))
		if err != nil {
			return fmt.Errorf("failed to find season: %w", err)
		}
		if s == nil {
			return season.ErrSeasonNotFound
		}
		if s.GetStatus() != season.SeasonStatusActive {
			return cropcycle.ErrSeasonNotActive
		}

		// 2. Проверяем место выращивания
		area, err := growingProvider.CultivationAreas().FindByID(ctx, cmd.AreaID)
		if err != nil {
			return fmt.Errorf("failed to find area: %w", err)
		}
		if area == nil {
			return cultivationarea.ErrAreaNotFound
		}

		// 3. Проверяем, что место настроено на этот сезон
		if !area.IsConfiguredForSeason(cmd.SeasonID) {
			return cultivationarea.ErrAreaNotConfiguredForSeason
		}

		// 4. Проверяем, что план культуры соответствует конфигурации
		configuredPlan, err := area.GetCropPlanForSeason(cmd.SeasonID)
		if err == nil && configuredPlan != cmd.CropPlanID {
			return cultivationarea.ErrCropPlanMismatch
		}

		// 5. Проверяем наличие активного цикла на этом месте
		exists, err := growingProvider.CropCycles().ExistsActiveForArea(ctx, cmd.AreaID)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("area already has an active cycle")
		}

		// 6. Получаем шаблон
		template, err := growingProvider.CropTemplates().FindByCropPlanIDAndVersion(ctx, cmd.CropPlanID, cmd.CropPlanVersion)
		if err != nil {
			return fmt.Errorf("failed to find template: %w", err)
		}
		if template == nil {
			return cropcycle.ErrTemplateNotFound
		}
		if template.GetStatus() != croptemplate.TemplateStatusPublished {
			return fmt.Errorf("template is not published")
		}

		// 7. Создаём цикл
		cycle = cropcycle.NewCropCycle(
			string(template.GetID()),
			cmd.AreaID,
			cmd.SeasonID,
			cmd.CropPlanID,
			cmd.CropPlanVersion,
		)

		// 8. Начинаем цикл
		if err := cycle.Start(); err != nil {
			return err
		}

		// 9. Сохраняем
		if err := growingProvider.CropCycles().Save(ctx, cycle); err != nil {
			return fmt.Errorf("failed to save cycle: %w", err)
		}

		uow.RegisterAggregate(cycle)
		return nil
	})

	if err != nil {
		return err
	}

	log.Printf("Crop cycle started: id=%s, area=%s, plan=%s version=%d",
		cycle.GetID(), cmd.AreaID, cmd.CropPlanID, cmd.CropPlanVersion)

	return nil
}
