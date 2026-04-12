package command

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/inmemory"
	"time"

	"github.com/google/uuid"
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
	BedID        string    `json:"bed_id"`
	Name         string    `json:"name"`
	VarietyID    string    `json:"variety_id"`
	SeasonStart  time.Time `json:"season_start"`
	SeasonEnd    time.Time `json:"season_end"`
	PlantingDate time.Time `json:"planting_date"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	AssignedTo   string    `json:"assigned_to"`
	AssignedName string    `json:"assigned_name"`
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

	err = uow.Execute(ctx, inmemory.NewGrowingProvider, func(provider repository.RepositoryProvider) error {
		// Приводим провайдер к нужному типу
		growingProvider, ok := provider.(*inmemory.GrowingProvider)
		if !ok {
			return fmt.Errorf("expected FarmProvider, got %T", provider)
		}
		// Получаем сорт из каталога
		variety, err := growingProvider.Catalogs().GetVariety(ctx, "", c.VarietyID)
		if err != nil {
			return fmt.Errorf("variety not found: %w", err)
		}

		// Создаем план
		planID := uuid.New().String()
		plan, err := cropplan.NewCropPlan(
			planID,
			c.BedID,
			c.Name,
			c.VarietyID,
			variety.Name,
			variety.SpeciesName,
			c.SeasonStart,
			c.SeasonEnd,
			c.PlantingDate,
			c.Latitude,
			c.Longitude,
			c.AssignedTo,
			c.AssignedName,
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
