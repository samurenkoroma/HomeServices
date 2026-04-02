package commands

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
	"samurenkoroma/services/internal/modules/farm/infrastructure/persistence/postgres"

	"github.com/google/uuid"
)

type CreatePhysicalObjectCmd struct {
	Type    string          `json:"type" validate:"required,oneof=field greenhouse"`
	Name    string          `json:"name" validate:"required"`
	GeoJSON spatial.GeoJSON `json:"geometry" validate:"required"`
	OwnerID string          `json:"ownerId" validate:"required"`

	// Специфические поля
	SoilType       *string  `json:"soilType,omitempty"`       // для field
	GreenhouseType *string  `json:"greenhouseType,omitempty"` // для greenhouse
	Width          *float64 `json:"width,omitempty"`          // для greenhouse
	Height         *float64 `json:"height,omitempty"`         // для greenhouse
	Length         *float64 `json:"length,omitempty"`         // для greenhouse
}
type createPhysicalHandler struct {
	uowFactory repository.Factory
}

func (h *createPhysicalHandler) Name() string {
	return "CreatePhysicalObject"
}

func NewCreatePhysicalObjectHandler(uowFactory repository.Factory) command.Handler {
	return &createPhysicalHandler{uowFactory: uowFactory}
}

func (h *createPhysicalHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(*CreatePhysicalObjectCmd)
	if !ok {
		return command.ErrInvalidCommandType
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = uow.Execute(ctx, postgres.NewFarmProvider, func(provider repository.RepositoryProvider) error {
		// Приводим провайдер к нужному типу
		farmProvider, ok := provider.(*postgres.FarmProvider)
		if !ok {
			return fmt.Errorf("expected FarmProvider, got %T", provider)
		}
		var newObj *physicalobject.PhysicalObject
		// 1. Создаем поле
		switch c.Type {
		case "field":
			area := spatial.CalculateAreaFromGeometry(uow.Tx(), c.GeoJSON)
			newObj = physicalobject.NewField(c.Name, c.GeoJSON, area, *c.SoilType, uuid.New().String())
		case "greenhouse":
			newObj = physicalobject.NewGreenhouse(c.Name, types.Dimension{
				Length: c.Length,
				Width:  c.Width,
				Height: c.Height,
			}, c.GeoJSON, "cntrkj", uuid.New().String())

		}

		// 2. Сохраняем поле
		if err := farmProvider.Objects().Save(ctx, newObj); err != nil {
			return fmt.Errorf("failed to save field: %w", err)
		}

		uow.RegisterAggregate(newObj)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
