package commands

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/crop/domain/croptype"
	"samurenkoroma/services/internal/modules/crop/infrastructure/persistence/postgres"
)

type CreateCropTypeCommand struct {
	Name        string `json:"name" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Description string `json:"description"`
	IsPerennial bool   `json:"is_perennial"`
}

type createCropTypeHandler struct {
	uowFactory repository.Factory
}

func (h *createCropTypeHandler) Name() string {
	return "CreateCropType"
}

func NewCreateCropTypeHandler(uowFactory repository.Factory) command.Handler {
	return &createCropTypeHandler{uowFactory: uowFactory}
}

func (h *createCropTypeHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(CreateCropTypeCommand)
	if !ok {
		return errors.New("invalid command type")
	}
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return err
	}

	return uow.Execute(ctx, postgres.NewCropProvider, func(provider repository.RepositoryProvider) error {
		cropProvider := provider.(*postgres.CropProvider)

		// Создаем тип культуры
		cropType, err := croptype.NewCropType(
			c.Name,
			croptype.CropCategory(c.Category),
			c.Description,
			c.IsPerennial,
		)
		if err != nil {
			return err
		}

		// Сохраняем
		if err := cropProvider.CropTypes().Save(ctx, cropType); err != nil {
			return err
		}

		uow.RegisterAggregate(cropType)
		return nil
	})
}
