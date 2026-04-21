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

type CreateFarmObjectCmd struct {
	Type    string          `json:"type" validate:"required,oneof=field greenhouse plot"`
	Name    string          `json:"name" validate:"required"`
	GeoJSON spatial.GeoJSON `json:"geometry" validate:"required"`

	Attributes struct {
		Length float64 `json:"length" validate:"required"`
		Width  float64 `json:"width" validate:"required"`
	} `json:"attributes,omitempty"`
}
type createPhysicalHandler struct {
	uowFactory repository.Factory
}

func NewCreateFarmObjectHandler(uowFactory repository.Factory) command.Handler {
	return &createPhysicalHandler{uowFactory: uowFactory}
}

func (h *createPhysicalHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(*CreateFarmObjectCmd)
	if !ok {
		return command.ErrInvalidCommandType
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = uow.Execute(ctx, postgres.NewPostgresFarmProvider, func(provider repository.RepositoryProvider) error {
		// Приводим провайдер к нужному типу
		farmProvider, ok := provider.(*postgres.FarmProvider)
		if !ok {
			return fmt.Errorf("expected FarmProvider, got %T", provider)
		}
		var newObj *physicalobject.PhysicalObject
		dim := types.Dimension{
			Length: c.Attributes.Length,
			Width:  c.Attributes.Width,
		}
		area := spatial.CalculateAreaFromGeometry(uow.Tx(), c.GeoJSON)
		// 1. Создаем поле
		switch c.Type {
		case "field":
			newObj = physicalobject.NewField(c.Name, c.GeoJSON, area, dim, "Чернозём", uuid.New().String())
		case "plot":
			newObj = physicalobject.NewPlot(c.Name, c.GeoJSON, area, dim, uuid.New().String())
		case "greenhouse":
			newObj = physicalobject.NewGreenhouse(c.Name, dim, c.GeoJSON, "film", uuid.New().String())

		}

		// 2. Сохраняем поле
		if err := farmProvider.Objects().Save(ctx, newObj); err != nil {
			return fmt.Errorf("failed to save field: %w", err)
		}

		uow.RegisterAggregate(newObj)
		return nil
	})

	return err
}
