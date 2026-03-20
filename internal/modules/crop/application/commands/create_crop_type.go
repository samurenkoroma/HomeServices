package commands

import (
	"context"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/crop/domain/croptype"
	"samurenkoroma/services/internal/modules/crop/infrastructure/persistence/postgres"
)

type CreateCropTypeCommand struct {
	Name           string `json:"name" validate:"required"`
	ScientificName string `json:"scientific_name"`
	Category       string `json:"category" validate:"required"`
	VegetationDays int    `json:"vegetation_days" validate:"required,gt=0"`
	Description    string `json:"description"`
}

type CreateCropTypeHandler struct {
	uowFactory repository.Factory
}

func (h *CreateCropTypeHandler) Handle(ctx context.Context, cmd CreateCropTypeCommand) error {
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return err
	}

	return uow.Execute(ctx, func(provider repository.RepositoryProvider) error {
		cropProvider := provider.(*postgres.CropProvider)

		// Создаем тип культуры
		cropType, err := croptype.NewCropType(
			cmd.Name,
			cmd.ScientificName,
			croptype.CropCategory(cmd.Category),
			cmd.VegetationDays,
		)
		if err != nil {
			return err
		}

		cropType.Description = cmd.Description

		// Сохраняем
		if err := cropProvider.CropTypes().Save(ctx, cropType); err != nil {
			return err
		}

		uow.RegisterAggregate(cropType)
		return nil
	})
}
