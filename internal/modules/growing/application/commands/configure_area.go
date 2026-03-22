package commands

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
)

type ConfigureAreaCommand struct {
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

func NewConfigureAreaHandler(uowFactory repository.Factory) *ConfigureAreaHandler {
	return &ConfigureAreaHandler{uowFactory: uowFactory}
}

func (h *ConfigureAreaHandler) Handle(ctx context.Context, cmd ConfigureAreaCommand) error {
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return err
	}

	return uow.Execute(ctx, func(provider repository.RepositoryProvider) error {
		growingProvider := provider.(*postgres.GrowingProvider)

		// Получаем operational объект
		area, err := growingProvider.CultivationAreas().GetByFarmRefID(ctx, cmd.FarmRefID)
		if err != nil {
			return err
		}

		config := cultivationarea.AreaConfig{
			Name:       cmd.Name,
			Geometry:   cmd.Geometry,
			CropPlanID: cmd.CropPlanID,
			Metadata:   cmd.Metadata,
		}

		// Конфигурируем в зависимости от типа
		switch area.GetType() {
		case cultivationarea.AreaTypeField:
			fieldArea := area.(*cultivationarea.FieldArea)
			if cmd.UsageType == "polyculture" {
				return fieldArea.ConfigureAsPolyculture(cmd.SeasonID, cmd.Name, cmd.Geometry, cmd.Metadata)
			}
			if cmd.CropPlanID == nil {
				return errors.New("crop_plan_id required for monoculture field")
			}
			return fieldArea.ConfigureAsMonoculture(cmd.SeasonID, cmd.Name, cmd.Geometry, *cmd.CropPlanID, cmd.Metadata)

		case cultivationarea.AreaTypeBlock:
			blockArea := area.(*cultivationarea.Block)
			return blockArea.ConfigureForSeason(cmd.SeasonID, cmd.Name, cmd.Geometry, *cmd.CropPlanID, cmd.Metadata)

		case cultivationarea.AreaTypeBed:
			bedArea := area.(*cultivationarea.Bed)
			return bedArea.ConfigureForSeason(cmd.SeasonID, cmd.Name, cmd.Geometry, *cmd.CropPlanID, cmd.Metadata)

		case cultivationarea.AreaTypeGreenhouse:
			greenhouseArea := area.(*cultivationarea.GreenhouseArea)
			return greenhouseArea.ConfigureForSeason(cmd.SeasonID, cmd.Name, cmd.Geometry, cmd.Metadata)
		}

		return nil
	})
}
