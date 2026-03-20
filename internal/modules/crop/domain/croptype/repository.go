package croptype

import "context"

type Repository interface {
	Save(ctx context.Context, ct *CropType) error
	FindByID(ctx context.Context, id CropTypeID) (*CropType, error)
	FindByName(ctx context.Context, name string) (*CropType, error)
	FindAll(ctx context.Context) ([]*CropType, error)
	FindByCategory(ctx context.Context, category CropCategory) ([]*CropType, error)
	FindActive(ctx context.Context) ([]*CropType, error)
	Exists(ctx context.Context, name string) (bool, error)
	Delete(ctx context.Context, id CropTypeID) error
}
