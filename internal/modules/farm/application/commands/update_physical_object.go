package commands

import (
	"context"
	"errors"
	"fmt"
	"log"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/modules/farm/infrastructure/persistence/postgres"

	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
)

// UpdatePhysicalObjectCommand — команда обновления физического объекта
type UpdatePhysicalObjectCommand struct {
	ID          string                 `json:"id" validate:"required"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Status      *string                `json:"status,omitempty"` // active, inactive
	Geometry    *spatial.GeoJSON       `json:"geometry,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
}

// UpdatePhysicalObjectHandler — обработчик обновления
type updatePhysicalObjectHandler struct {
	uowFactory repository.Factory
}

func (h *updatePhysicalObjectHandler) Name() string {
	return "UpdatePhysicalObject"
}

// NewUpdatePhysicalObjectHandler создаёт новый обработчик
func NewUpdatePhysicalObjectHandler(uowFactory repository.Factory) command.Handler {
	return &updatePhysicalObjectHandler{
		uowFactory: uowFactory,
	}
}

// Handle обрабатывает команду
func (h *updatePhysicalObjectHandler) Handle(ctx context.Context, command any) error {

	cmd, ok := command.(UpdatePhysicalObjectCommand)
	if !ok {
		return errors.New("invalid command type")
	} // Валидация статуса
	if cmd.Status != nil {
		if *cmd.Status != "active" && *cmd.Status != "inactive" {
			return fmt.Errorf("invalid status: %s, must be 'active' or 'inactive'", *cmd.Status)
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
		obj, err := farmProvider.Objects().FindByID(ctx, physicalobject.PhysicalObjectID(cmd.ID))
		if err != nil {
			return fmt.Errorf("failed to find physical object: %w", err)
		}
		if obj == nil {
			return physicalobject.ErrPhysicalObjectNotFound
		}

		// Обновляем поля
		if cmd.Name != nil {
			obj.SetName(*cmd.Name)
		}

		if cmd.Description != nil {
			obj.SetDescription(*cmd.Description)
		}

		if cmd.Status != nil {
			if *cmd.Status == "active" {
				if err := obj.Activate(); err != nil {
					return err
				}
			} else if *cmd.Status == "inactive" {
				if err := obj.Deactivate(); err != nil {
					return err
				}
			}
		}

		if cmd.Geometry != nil {
			if err := obj.SetGeometry(*cmd.Geometry); err != nil {
				return fmt.Errorf("failed to set geometry: %w", err)
			}
		}

		if cmd.Attributes != nil {
			obj.SetAttributes(cmd.Attributes)
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

	log.Printf("Physical object updated: id=%s", cmd.ID)
	return nil
}
