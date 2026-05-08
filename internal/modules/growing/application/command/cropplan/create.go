package cropplan

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"samurenkoroma/services/internal/modules/growing/domain/utils"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
	"time"

	"github.com/google/uuid"
)

type CreateCropPlanCmd struct {
	CropID            string  `json:"cropId"`
	VarietyID         *string `json:"varietyId,omitempty"`
	CultivationPlanID string  `json:"cultivationPlanId"`
	SeasonID          string  `json:"seasonId"`
	AreaID            string  `json:"areaId"`

	Status    string    `json:"status"`
	StartDate time.Time `json:"startDate,format:date"`
}

func (h *PlanHandler) Create(ctx context.Context, cmd any) (any, error) {

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

		tpl, err := growingProvider.Cultivation().GetLatest(ctx, c.CultivationPlanID)
		if err != nil {
			return nil, err
		}

		plan := &cropplan.Plan{
			ID: uuid.New().String(),

			CropKey: "tomato",

			CultivationPlanID:      tpl.ID,
			CultivationPlanVersion: tpl.Version,

			Snapshot: utils.BuildSnapshot(tpl),
		}
		if err := growingProvider.CropPlans().Save(ctx, plan); err != nil {
			return nil, err
		}

		uow.RegisterAggregate(plan)
		return nil, nil
	})

}
