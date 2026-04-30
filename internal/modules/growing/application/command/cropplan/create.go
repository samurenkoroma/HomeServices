package cropplan

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"samurenkoroma/services/internal/modules/growing/domain/season"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
	"time"
)

type CreateCropPlanCmd struct {
	AreaID       string    `json:"areaID"`
	Name         string    `json:"name"`
	VarietyID    string    `json:"varietyId"`
	SpeciesKey   string    `json:"speciesKey"`
	SeasonId     string    `json:"seasonId"`
	PlantingDate time.Time `json:"plantingDate"`
	AssignedTo   string    `json:"assignedTo"`
}

func (h *CropPlanHandler) Create(ctx context.Context, cmd any) (any, error) {
	c, ok := cmd.(*CreateCropPlanCmd)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return uow.Execute(ctx, postgres.NewPostgresGrowingProvider, func(provider repository.RepositoryProvider) (any, error) {
		// Приводим провайдер к нужному типу
		growingProvider, ok := provider.(*postgres.PostgresGrowingProvider)
		if !ok {
			return nil, fmt.Errorf("expected FarmProvider, got %T", provider)
		}
		// Получаем сорт из каталога
		variety, err := growingProvider.Catalog().GetVariety(ctx, c.VarietyID)
		if err != nil {
			return nil, err
		}

		seasons, err := growingProvider.Seasons().FindByID(ctx, season.SeasonID(c.SeasonId))
		if err != nil {
			return nil, err
		}
		area, err := growingProvider.CultivationAreas().FindByID(ctx, c.AreaID)
		if err != nil {
			return nil, err
		}
		// Создаем план
		plan, err := cropplan.NewCropPlan(
			types.NewUUID(),
			c.Name,
			c.PlantingDate,
			area,
			seasons,
			variety,
			c.AssignedTo,
		)
		if err != nil {
			return nil, err
		}

		// Добавляем этапы из шаблонов
		templates := catalog.GetStageTemplatesForSpecies(variety.SpeciesKey)
		if err := plan.AddStagesFromTemplates(templates); err != nil {
			return nil, err
		}

		// Сохраняем план

		if err := growingProvider.CropPlans().Save(ctx, plan); err != nil {
			return nil, err
		}

		uow.RegisterAggregate(plan)

		return nil, nil
	})
}
