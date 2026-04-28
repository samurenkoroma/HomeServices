package commands

import (
	"context"
	"encoding/json"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/modules/farm/infrastructure/persistence/postgres"

	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
)

type UpdateFarmObjectCommand struct {
	ID          string                 `json:"id" validate:"required"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Status      *string                `json:"status,omitempty"` // active, inactive
	Geometry    *spatial.GeoJSON       `json:"geometry,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
	Schema      json.RawMessage        `json:"schema,omitempty"`
}

func (h *FarmObjectHandler) Update(ctx context.Context, cmd any) (any, error) {
	c, ok := cmd.(*UpdateFarmObjectCommand)
	if !ok {
		return nil, command.ErrInvalidCommandType
	}

	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return uow.Execute(ctx, postgres.NewPostgresFarmProvider, func(provider repository.RepositoryProvider) (any, error) {
		farmProvider, ok := provider.(*postgres.FarmProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}

		// Получаем объект
		obj, err := farmProvider.Objects().FindByID(ctx, physicalobject.PhysicalObjectID(c.ID))
		if err != nil {
			return nil, err
		}
		if obj == nil {
			return nil, physicalobject.ErrPhysicalObjectNotFound
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
					return nil, err
				}
			} else if *c.Status == "inactive" {
				if err := obj.Deactivate(); err != nil {
					return nil, err
				}
			}
		}

		if c.Geometry != nil {
			if err := obj.SetGeometry(*c.Geometry); err != nil {
				return nil, err
			}
		}

		if c.Attributes != nil {
			obj.SetAttributes(c.Attributes)
		}
		if c.Schema != nil {
			obj.SetSchema(c.Schema)
		}

		// Сохраняем
		if err := farmProvider.Objects().Save(ctx, obj); err != nil {
			return nil, err
		}

		uow.RegisterAggregate(obj)
		return nil, nil
	})
}
