package commands

import (
	"context"
	"errors"
	"fmt"
	"log"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
)

type ConfigureAreaCmd struct {
	FarmRefID  string                 `json:"farm_ref_id" validate:"required"`
	SeasonID   string                 `json:"season_id" validate:"required"`
	Name       string                 `json:"name"`
	Geometry   spatial.GeoJSON        `json:"geometry"`
	CropPlanID *string                `json:"crop_plan_id,omitempty"`
	UsageType  string                 `json:"usage_type"` // monoculture/polyculture (для field)
	Metadata   map[string]interface{} `json:"metadata"`
}

type ConfigureAreaHandler struct {
	uowFactory repository.Factory
}

func (h *ConfigureAreaHandler) Name() string {
	return "ConfigureArea"
}

func NewConfigureAreaHandler(uowFactory repository.Factory) command.Handler {
	return &ConfigureAreaHandler{uowFactory: uowFactory}
}

func (h *ConfigureAreaHandler) Handle(ctx context.Context, command any) error {
	cmd, ok := command.(ConfigureAreaCmd)
	if !ok {
		return errors.New("invalid command type")
	}
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return err
	}
	var area cultivationarea.CultivationArea
	err = uow.Execute(ctx, postgres.NewGrowingProvider, func(provider repository.RepositoryProvider) error {
		growingProvider, ok := provider.(*postgres.GrowingProvider)
		if !ok {
			return fmt.Errorf("invalid provider type")
		}

		// Получаем operational объект
		area, err = growingProvider.CultivationAreas().FindByFarmRefID(ctx, cmd.FarmRefID)
		if err != nil {
			return fmt.Errorf("failed to find cultivation area: %w", err)
		}

		if area == nil {
			return cultivationarea.ErrAreaNotFound
		}

		// Создаём конфигурацию
		config := cultivationarea.AreaConfig{
			Name:       cmd.Name,
			Geometry:   cmd.Geometry,
			CropPlanID: cmd.CropPlanID,
			Metadata:   cmd.Metadata,
		}

		// Используем config для настройки места
		// Конфигурируем в зависимости от типа
		switch area.GetType() {
		case cultivationarea.AreaTypeField:
			fieldArea := area.(*cultivationarea.FieldArea)

			// Если указан тип использования
			if cmd.UsageType == "polyculture" {
				// Поликультура — поле с участками
				err = fieldArea.ConfigureAsPolyculture(
					cmd.SeasonID,
					config.Name,
					config.Geometry,
					config.Metadata,
				)
			} else {
				// Монокультура — одна культура на всём поле
				if config.CropPlanID == nil {
					return fmt.Errorf("crop_plan_id required for monoculture field")
				}
				err = fieldArea.ConfigureAsMonoculture(
					cmd.SeasonID,
					config.Name,
					config.Geometry,
					*config.CropPlanID,
					config.Metadata,
				)
			}

		case cultivationarea.AreaTypeBlock:
			blockArea := area.(*cultivationarea.Block)
			if config.CropPlanID == nil {
				return fmt.Errorf("crop_plan_id required for block")
			}
			err = blockArea.ConfigureForSeason(cmd.SeasonID, config)

		case cultivationarea.AreaTypeBed:
			bedArea := area.(*cultivationarea.Bed)
			if config.CropPlanID == nil {
				return fmt.Errorf("crop_plan_id required for bed")
			}
			err = bedArea.ConfigureForSeason(cmd.SeasonID, config)

		case cultivationarea.AreaTypeGreenhouse:
			greenhouseArea := area.(*cultivationarea.GreenhouseArea)
			err = greenhouseArea.ConfigureForSeason(cmd.SeasonID, config)

		default:
			return fmt.Errorf("unsupported area type: %s", area.GetType())
		}

		if err != nil {
			return fmt.Errorf("failed to configure area: %w", err)
		}

		// Сохраняем конфигурацию
		if err := growingProvider.CultivationAreas().SaveSeasonConfig(ctx, area, cmd.SeasonID); err != nil {
			return fmt.Errorf("failed to save season config: %w", err)
		}

		uow.RegisterAggregate(area)
		return nil
	})
	if err != nil {
		return err
	}

	log.Printf("Area configured: farm_ref_id=%s, season_id=%s, type=%s, name=%s",
		cmd.FarmRefID, cmd.SeasonID, area.GetType(), cmd.Name)

	return nil
}
