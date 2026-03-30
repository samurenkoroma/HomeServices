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
	log.Printf("=== CONFIGURE AREA COMMAND ===")
	log.Printf("Command: %+v", command)

	cmd, ok := command.(*ConfigureAreaCmd)
	if !ok {
		return errors.New("invalid command")
	}
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return fmt.Errorf("failed to begin transaction: %w", err)
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

		// Получаем геометрию из farm объекта, если не передана
		geometry := cmd.Geometry
		if geometry.Type == "" {
			// Берём геометрию из существующего operational объекта
			geometry = area.GetGeometry()
			log.Printf("Using geometry from existing area: type=%s", geometry.Type)
		}

		config := cultivationarea.AreaConfig{
			Name:       cmd.Name,
			Geometry:   geometry,
			CropPlanID: cmd.CropPlanID,
			Metadata:   cmd.Metadata,
		}

		// Если имя не указано, используем имя из operational объекта
		if config.Name == "" {
			config.Name = area.GetName()
		}

		// Конфигурируем
		if err := area.ConfigureForSeason(cmd.SeasonID, config); err != nil {
			log.Printf("ConfigureForSeason error: %v", err)
			return err
		}

		// Сохраняем конфигурацию
		seasonConfig, err := area.GetSeasonConfig(cmd.SeasonID)
		if err != nil {
			log.Printf("GetSeasonConfig error: %v", err)
			return err
		}

		// Убеждаемся, что геометрия не пустая
		if seasonConfig.Geometry.Type == "" {
			return fmt.Errorf("geometry is empty for area %s", area.GetID())
		}

		log.Printf("Saving season config with geometry type: %s", seasonConfig.Geometry.Type)

		if err := growingProvider.CultivationAreas().SaveSeasonConfig(ctx, area.GetID(), *seasonConfig); err != nil {
			log.Printf("SaveSeasonConfig error: %v", err)
			return fmt.Errorf("failed to save season config: %w", err)
		}

		uow.RegisterAggregate(area)
		return nil
	})

	if err != nil {
		log.Printf("Execute error: %v", err)
		return err
	}

	log.Printf("Area configured successfully: farm_ref_id=%s, season_id=%s", cmd.FarmRefID, cmd.SeasonID)
	return nil
}
