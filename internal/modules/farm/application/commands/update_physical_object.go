package commands

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/modules/farm/infrastructure/persistence/postgres"

	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
)

type UpdatePhysicalObjectCommand struct {
	ID          string                 `json:"id" validate:"required"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Status      *string                `json:"status,omitempty"` // active, inactive
	Geometry    *spatial.GeoJSON       `json:"geometry,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
}
type updatePhysicalObjectHandler struct {
	uowFactory repository.Factory
}

func (h *updatePhysicalObjectHandler) Name() string {
	return "UpdatePhysicalObject"
}

func NewUpdatePhysicalObjectHandler(uowFactory repository.Factory) command.Handler {
	return &updatePhysicalObjectHandler{uowFactory: uowFactory}
}

// Handle обрабатывает команду
func (h *updatePhysicalObjectHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(UpdatePhysicalObjectCommand)
	if !ok {
		return command.ErrInvalidCommandType
	}

	if c.Status != nil {
		if *c.Status != "active" && *c.Status != "inactive" {
			return fmt.Errorf("invalid status: %s, must be 'active' or 'inactive'", *c.Status)
		}
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = uow.Execute(ctx, postgres.NewFarmProvider, func(provider repository.RepositoryProvider) error {
		farmProvider, ok := provider.(*postgres.FarmProvider)
		if !ok {
			return fmt.Errorf("invalid provider type")
		}

		// Получаем объект
		obj, err := farmProvider.Objects().FindByID(ctx, physicalobject.PhysicalObjectID(c.ID))
		if err != nil {
			return fmt.Errorf("failed to find physical object: %w", err)
		}
		if obj == nil {
			return physicalobject.ErrPhysicalObjectNotFound
		}

		// Обновляем поля
		if c.Name != nil {
			obj.SetName(*c.Name)
		}

		if c.Description != nil {
			obj.SetDescription(*c.Description)
		}

		if c.Status != nil {
			if *c.Status == "active" {
				if err := obj.Activate(); err != nil {
					return err
				}
			} else if *c.Status == "inactive" {
				if err := obj.Deactivate(); err != nil {
					return err
				}
			}
		}

		if c.Geometry != nil {
			if err := obj.SetGeometry(*c.Geometry); err != nil {
				return fmt.Errorf("failed to set geometry: %w", err)
			}
		}

		if c.Attributes != nil {
			obj.SetAttributes(c.Attributes)
		}

		// Сохраняем
		if err := farmProvider.Objects().Save(ctx, obj); err != nil {
			return fmt.Errorf("failed to save physical object: %w", err)
		}

		uow.RegisterAggregate(obj)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
