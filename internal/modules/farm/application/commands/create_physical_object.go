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

	Attributes map[string]interface{} `json:"attributes,omitempty"`
}
type createPhysicalHandler struct {
	uowFactory repository.Factory
}

func (h *createPhysicalHandler) Name() string {
	return "CreateObject"
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

	err = uow.Execute(ctx, postgres.NewFarmProvider, func(provider repository.RepositoryProvider) error {
		// Приводим провайдер к нужному типу
		farmProvider, ok := provider.(*postgres.FarmProvider)
		if !ok {
			return fmt.Errorf("expected FarmProvider, got %T", provider)
		}
		var newObj *physicalobject.PhysicalObject
		length := c.Attributes["length"].(float64)
		width := c.Attributes["width"].(float64)
		dim := types.Dimension{
			Length: length,
			Width:  width,
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

	if err != nil {
		return err
	}

	return nil
}
