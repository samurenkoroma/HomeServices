package command

import (
	"context"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/crop/domain/cropplan"
)

type CreateCropPlanCommand struct {
	CropTypeID  string `json:"crop_type_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Duration    int    `json:"duration" validate:"required,gt=0"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by" validate:"required"`
}

type CreateCropPlanHandler struct {
	uowFactory repository.Factory
}

func NewCreateCropPlanHandler(uowFactory repository.Factory) *CreateCropPlanHandler {
	return &CreateCropPlanHandler{uowFactory: uowFactory}
}

func (h *CreateCropPlanHandler) Handle(ctx context.Context, cmd CreateCropPlanCommand) error {
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return err
	}

	return uow.Execute(ctx, func(provider repository.RepositoryProvider) error {
		cropProvider := provider.(*postgres.CropProvider)

		// Создаем план
		plan, err := cropplan.NewCropPlan(
			cmd.CropTypeID,
			cmd.Name,
			cmd.Duration,
			cmd.CreatedBy,
		)
		if err != nil {
			return err
		}

		plan.Description = cmd.Description

		// Сохраняем
		if err := cropProvider.CropPlans().Save(ctx, plan); err != nil {
			return err
		}

		uow.RegisterAggregate(plan)
		return nil
	})
}
