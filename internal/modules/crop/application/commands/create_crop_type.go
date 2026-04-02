package commands

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/crop/domain/croptype"
	"samurenkoroma/services/internal/modules/crop/infrastructure/persistence/postgres"
)

type CreateCropTypeCmd struct {
	Name        string  `json:"name" validate:"required"`
	Category    string  `json:"category" validate:"required"`
	CategoryRu  *string `json:"categoryRu,omitempty"`
	Family      string  `json:"family" validate:"required"`
	FamilyRu    *string `json:"familyRu,omitempty"`
	Icon        string  `json:"icon"`
	ImageURL    string  `json:"imageURL"`
	Description string  `json:"description"`
	IsPerennial bool    `json:"is_perennial"`
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
	c, ok := cmd.(CreateCropTypeCmd)
	if !ok {
		return command.ErrInvalidCommandType
	}
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	return uow.Execute(ctx, postgres.NewCropProvider, func(provider repository.RepositoryProvider) error {
		cropProvider := provider.(*postgres.CropProvider)

		// Создаем тип культуры
		cropType, err := croptype.NewCropType(
			c.Name,
			croptype.CropCategory(c.Category),
			croptype.CropFamily(c.Family),
			c.Description,
			c.IsPerennial,
		)
		if err != nil {
			return err
		}
		cropType.AddUI(c.Icon, c.ImageURL)

		// Сохраняем
		if err := cropProvider.CropTypes().Save(ctx, cropType); err != nil {

			return err
		}

		if c.CategoryRu != nil {
			cropProvider.Translations().Save(ctx, "crop_category", c.Category, *c.CategoryRu)
		}
		if c.FamilyRu != nil {
			cropProvider.Translations().Save(ctx, "crop_family", c.Family, *c.FamilyRu)
		}
		uow.RegisterAggregate(cropType)
		return nil
	})
}
