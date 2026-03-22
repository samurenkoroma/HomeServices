package croptemplate

import (
	"context"
)

type Repository interface {
	// CRUD
	Save(ctx context.Context, template *CropTemplate) error
	FindByID(ctx context.Context, id TemplateID) (*CropTemplate, error)
	FindByCropPlanID(ctx context.Context, cropPlanID string) (*CropTemplate, error)
	FindByCropPlanIDAndVersion(ctx context.Context, cropPlanID string, version int) (*CropTemplate, error)
	FindAll(ctx context.Context) ([]*CropTemplate, error)
	FindPublished(ctx context.Context) ([]*CropTemplate, error)
	Delete(ctx context.Context, id TemplateID) error

	// Проверки
	Exists(ctx context.Context, cropPlanID string, version int) (bool, error)
	GetLatestVersion(ctx context.Context, cropPlanID string) (int, error)
}
