package commands

import (
	"context"
	"errors"
	"fmt"
	"log"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"

	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
)

// ConfigureAreaCommand — команда настройки места на сезон
type ConfigureAreaCmd struct {
	FarmRefID  string                 `json:"farm_ref_id" validate:"required"`
	SeasonID   string                 `json:"season_id" validate:"required"`
	Name       string                 `json:"name"`
	Geometry   spatial.GeoJSON        `json:"geometry"`
	CropPlanID *string                `json:"crop_plan_id,omitempty"`
	UsageType  string                 `json:"usage_type"` // monoculture/polyculture (для field)
	Metadata   map[string]interface{} `json:"metadata"`
}

// ConfigureAreaHandler — обработчик настройки места
type ConfigureAreaHandler struct {
	uowFactory repository.Factory
}

func (h *ConfigureAreaHandler) Name() string {
	return "ConfigureArea"
}

// NewConfigureAreaHandler создаёт новый обработчик
func NewConfigureAreaHandler(uowFactory repository.Factory) command.Handler {
	return &ConfigureAreaHandler{
		uowFactory: uowFactory,
	}
}

// Handle обрабатывает команду
func (h *ConfigureAreaHandler) Handle(ctx context.Context, command any) error {
	cmd, ok := command.(*ConfigureAreaCmd)
	if !ok {
		return errors.New("invalid command type")
	} // Валидация
	if cmd.FarmRefID == "" {
		return fmt.Errorf("farm_ref_id is required")
	}
	if cmd.SeasonID == "" {
		return fmt.Errorf("season_id is required")
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	var area cultivationarea.CultivationArea

	err = uow.Execute(ctx, postgres.NewGrowingProvider, func(provider repository.RepositoryProvider) error {
		growingProvider, ok := provider.(*postgres.GrowingProvider)
		if !ok {
			return fmt.Errorf("invalid provider type")
		}

		// Получаем operational объект по ссылке на farm
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

		// Если имя не указано, используем текущее имя места
		if config.Name == "" {
			config.Name = area.GetName()
		}

		// Если геометрия не указана, используем текущую геометрию места
		if config.Geometry.Type == "" {
			config.Geometry = area.GetGeometry()
		}

		// Конфигурируем в зависимости от типа места
		switch area.GetType() {
		case cultivationarea.AreaTypeField:
			fieldArea, ok := area.(*cultivationarea.FieldArea)
			if !ok {
				return fmt.Errorf("failed to cast to FieldArea")
			}

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
			blockArea, ok := area.(*cultivationarea.Block)
			if !ok {
				return fmt.Errorf("failed to cast to Block")
			}
			if config.CropPlanID == nil {
				return cultivationarea.ErrCropPlanRequiredForBlock
			}
			err = blockArea.ConfigureForSeason(
				cmd.SeasonID,
				config,
			)

		case cultivationarea.AreaTypeBed:
			bedArea, ok := area.(*cultivationarea.Bed)
			if !ok {
				return fmt.Errorf("failed to cast to Bed")
			}
			if config.CropPlanID == nil {
				return cultivationarea.ErrCropPlanRequiredForBed
			}
			err = bedArea.ConfigureForSeason(
				cmd.SeasonID,
				config,
			)

		case cultivationarea.AreaTypeGreenhouse:
			greenhouseArea, ok := area.(*cultivationarea.GreenhouseArea)
			if !ok {
				return fmt.Errorf("failed to cast to GreenhouseArea")
			}
			err = greenhouseArea.ConfigureForSeason(
				cmd.SeasonID,
				config,
			)

		default:
			return fmt.Errorf("unsupported area type: %s", area.GetType())
		}

		if err != nil {
			return fmt.Errorf("failed to configure area: %w", err)
		}

		// Сохраняем конфигурацию в БД
		seasonConfig, err := area.GetSeasonConfig(cmd.SeasonID)
		if err != nil {
			return fmt.Errorf("failed to get season config: %w", err)
		}

		if err := growingProvider.CultivationAreas().SaveSeasonConfig(ctx, area.GetID(), *seasonConfig); err != nil {
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
