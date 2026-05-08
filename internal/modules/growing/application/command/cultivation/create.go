package cultivation

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cultivation"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
)

type CreateCultivationPlanCmd struct {
	Version   int     `json:"version" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	CropKey   string  `json:"cropKey" validate:"required"`
	VarietyId *string `json:"varietyId" `
	Steps     []struct {
		Type    string `json:"type" validate:"required"`
		Title   string `json:"title" validate:"required"`
		Trigger struct {
			Type  string         `json:"type" validate:"required"`
			Value map[string]any `json:"value" validate:"required,min=1,dive,keys,min=3,endkeys,required"`
		} `json:"trigger" validate:"required"`
	} `json:"steps" validate:"required"`
}

func (h *PlanHandler) Create(ctx context.Context, cmd any) (any, error) {

	c, ok := cmd.(*CreateCultivationPlanCmd)
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

		cultivationPlan := cultivation.NewCultivationPlan(c.Name, c.CropKey, c.VarietyId, c.Version)

		for i, step := range c.Steps {
			step := cultivation.NewStep(
				uint8(i+1),
				step.Type,
				cultivation.Trigger{
					Type:  cultivation.TriggerType(step.Trigger.Type),
					Value: step.Trigger.Value,
				})
			cultivationPlan.AddStep(step)
		}

		err := growingProvider.Cultivation().Save(ctx, cultivationPlan)
		if err != nil {
			return nil, err
		}

		//uow.RegisterAggregate(tpl)
		return nil, nil
	})

}
