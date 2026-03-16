package field

import (
	"context"
	"samurenkoroma/services/internal/core/domain/types"
)

type Repository interface {
	Save(ctx context.Context, field *Field) error
	FindByID(ctx context.Context, id types.FieldId) (*Field, error)
	// FindAll Update(ctx context.Context, field *Field) error
	FindAll(ctx context.Context) ([]*Field, error)
	//FindAvailable(ctx context.Context) ([]*Field, error)
}
