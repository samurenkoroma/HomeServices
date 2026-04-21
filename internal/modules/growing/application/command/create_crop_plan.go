package command

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

// CreateCropPlanHandler команда создания плана
type createCropPlanHandler struct {
	uowFactory repository.Factory
}

func (h *createCropPlanHandler) Name() string {
	return "CreateCropPlan"
}

func NewCreateCropPlanHandler(uowFactory repository.Factory) command.Handler {
	return &createCropPlanHandler{uowFactory: uowFactory}
}

// CreateCropPlanCmd структура команды
type CreateCropPlanCmd struct {
	AreaID       string    `json:"areaID"`
	Name         string    `json:"name"`
	VarietyID    string    `json:"varietyId"`
	SpeciesKey   string    `json:"speciesKey"`
	SeasonId     string    `json:"seasonId"`
	PlantingDate time.Time `json:"plantingDate"`
	AssignedTo   string    `json:"assignedTo"`
}

// Handle выполняет команду
func (h *createCropPlanHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(*CreateCropPlanCmd)
	if !ok {
		return command.ErrInvalidCommandType
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = uow.Execute(ctx, postgres.NewPostgresGrowingProvider, func(provider repository.RepositoryProvider) error {
		// Приводим провайдер к нужному типу
		growingProvider, ok := provider.(*postgres.PostgresGrowingProvider)
		if !ok {
			return fmt.Errorf("expected FarmProvider, got %T", provider)
		}
		// Получаем сорт из каталога
		variety, err := growingProvider.Catalog().GetVariety(ctx, c.SpeciesKey, c.VarietyID)
		if err != nil {
			return fmt.Errorf("variety not found: %w", err)
		}

		seasons, err := growingProvider.Seasons().FindByID(ctx, season.SeasonID(c.SeasonId))
		if err != nil {
			return fmt.Errorf("variety not found: %w", err)
		}
		area, err := growingProvider.CultivationAreas().FindByID(ctx, c.AreaID)
		if err != nil {
			return fmt.Errorf("variety not found: %w", err)
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
			return fmt.Errorf("failed to create plan: %w", err)
		}

		// Добавляем этапы из шаблонов
		templates := catalog.GetStageTemplatesForSpecies(variety.SpeciesKey)
		if err := plan.AddStagesFromTemplates(templates); err != nil {
			return fmt.Errorf("failed to add stages: %w", err)
		}

		// Сохраняем план

		if err := growingProvider.CropPlans().Save(ctx, plan); err != nil {
			return fmt.Errorf("failed to save plan: %w", err)
		}

		uow.RegisterAggregate(plan)

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
