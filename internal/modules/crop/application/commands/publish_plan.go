package command

import (
	"context"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/crop/domain/cropplan"
	"samurenkoroma/services/internal/modules/crop/infrastructure/integration"
)

type PublishPlanCommand struct {
	PlanID string `json:"plan_id" validate:"required"`
}

type PublishPlanHandler struct {
	uowFactory    repository.Factory
	growingClient *integration.GrowingClient
}

func NewPublishPlanHandler(
	uowFactory repository.Factory,
	growingClient *integration.GrowingClient,
) *PublishPlanHandler {
	return &PublishPlanHandler{
		uowFactory:    uowFactory,
		growingClient: growingClient,
	}
}

func (h *PublishPlanHandler) Handle(ctx context.Context, cmd PublishPlanCommand) error {
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return err
	}

	return uow.Execute(ctx, func(provider repository.RepositoryProvider) error {
		cropProvider := provider.(*postgres.CropProvider)

		// Получаем план
		plan, err := cropProvider.CropPlans().GetByID(ctx, cropplan.PlanID(cmd.PlanID))
		if err != nil {
			return err
		}

		// Публикуем
		if err := plan.Publish(); err != nil {
			return err
		}

		// Сохраняем
		if err := cropProvider.CropPlans().Save(ctx, plan); err != nil {
			return err
		}

		uow.RegisterAggregate(plan)
		return nil
	})
}
