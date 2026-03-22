package cropcycle

import (
	"context"
)

type Repository interface {
	// CRUD
	Save(ctx context.Context, cycle *CropCycle) error
	FindByID(ctx context.Context, id CycleID) (*CropCycle, error)
	FindByAreaID(ctx context.Context, areaID string) ([]*CropCycle, error)
	FindBySeasonID(ctx context.Context, seasonID string) ([]*CropCycle, error)
	FindByStatus(ctx context.Context, status Status) ([]*CropCycle, error)
	FindAll(ctx context.Context) ([]*CropCycle, error)
	Delete(ctx context.Context, id CycleID) error

	// Активные циклы
	FindActiveByArea(ctx context.Context, areaID string) (*CropCycle, error)
	FindActiveBySeason(ctx context.Context, seasonID string) ([]*CropCycle, error)

	// Проверки
	Exists(ctx context.Context, id CycleID) (bool, error)
	ExistsActiveForArea(ctx context.Context, areaID string) (bool, error)
}
