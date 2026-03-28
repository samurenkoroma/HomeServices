package commands

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/crop/domain/variety"
	"samurenkoroma/services/internal/modules/crop/infrastructure/persistence/postgres"
)

type CreateVarietyCmd struct {
	Name        string `json:"name" validate:"required"`
	Crop        string `json:"crop" validate:"required"`
	Description string `json:"description"`

	VegetationDays     string   `json:"vegetation_days,omitempty"`
	YieldPotential     string   `json:"yield_potential,omitempty"`
	RecommendedRegions []string `json:"recommended_regions,omitempty"`
}

type createVarietyHandler struct {
	uowFactory repository.Factory
}

func (h *createVarietyHandler) Name() string {
	return "CreateVariety"
}

func NewCreateVarietyHandler(uowFactory repository.Factory) command.Handler {
	return &createVarietyHandler{
		uowFactory: uowFactory,
	}
}

func (h *createVarietyHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(CreateVarietyCmd)
	if !ok {
		return command.ErrInvalidCommandType
	}
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = uow.Execute(ctx, postgres.NewCropProvider, func(provider repository.RepositoryProvider) error {
		// Приводим провайдер к нужному типу
		cropProvider, ok := provider.(*postgres.CropProvider)
		if !ok {
			return fmt.Errorf("expected CropProvider, got %T", provider)
		}
		crop, err := cropProvider.CropTypes().FindByName(ctx, c.Crop)
		if err != nil {
			return err
		}
		newVariety, err := variety.NewVariety(string(crop.Id), c.Name, c.VegetationDays, c.YieldPotential)
		if err != nil {
			return err
		}

		if err := cropProvider.Varieties().Save(ctx, newVariety); err != nil {
			return err
		}
		uow.RegisterAggregate(newVariety)
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
