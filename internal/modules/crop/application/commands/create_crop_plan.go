package commands

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/crop/domain/cropplan"
	"samurenkoroma/services/internal/modules/crop/infrastructure/persistence/postgres"
)

type CreateCropPlanCommand struct {
	CropTypeID  string `json:"crop_type_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Duration    int    `json:"duration" validate:"required,gt=0"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by" validate:"required"`
}

type createCropPlanHandler struct {
	uowFactory repository.Factory
}

func (h *createCropPlanHandler) Name() string {
	return "CreateCropPlan"
}

func NewCreateCropPlanHandler(uowFactory repository.Factory) command.Handler {
	return &createCropPlanHandler{uowFactory: uowFactory}
}

func (h *createCropPlanHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(CreateCropPlanCommand)
	if !ok {
		return errors.New("invalid command type")
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return err
	}

	return uow.Execute(ctx, postgres.NewCropProvider, func(provider repository.RepositoryProvider) error {
		cropProvider := provider.(*postgres.CropProvider)

		// Создаем план
		plan, err := cropplan.NewCropPlan(
			c.CropTypeID,
			c.Name,
			c.Duration,
			c.CreatedBy,
		)
		if err != nil {
			return err
		}

		plan.Description = c.Description

		// Сохраняем
		if err := cropProvider.CropPlans().Save(ctx, plan); err != nil {
			return err
		}

		uow.RegisterAggregate(plan)
		return nil
	})
}
