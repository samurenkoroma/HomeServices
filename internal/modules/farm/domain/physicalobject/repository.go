package physicalobject

import (
	"context"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/core/spatial"
)

type Repository interface {
	FindByID(context.Context, types.PhysicalObjectID) (*PhysicalObject, error)
	Save(context.Context, *PhysicalObject) error
	FindAll(ctx context.Context) ([]*PhysicalObject, error)

	FindInBounds(ctx context.Context, bounds spatial.BoundingBox) ([]*PhysicalObject, error)
	/*
		// Базовые CRUD
		Delete(ctx context.Context, id types.PhysicalObjectID) error

		// Специфичные запросы
		FindByType(ctx context.Context, objType ObjectType) ([]*PhysicalObject, error)
		FindByOwner(ctx context.Context, ownerID string) ([]*PhysicalObject, error)

		// Геопространственные запросы

		FindNearby(ctx context.Context, lat, lon, radius float64) ([]*PhysicalObject, error)
	*/
}
