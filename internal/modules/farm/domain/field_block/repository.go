package field_block

import (
	"context"
	"samurenkoroma/services/internal/core/domain/types"
)

type Repository interface {
	Save(ctx context.Context, block *FieldBlock) error
	//Update(ctx context.Context, block *FieldBlock) error
	FindByID(ctx context.Context, id types.FieldBlockId) (*FieldBlock, error)
	//FindByFieldID(ctx context.Context, fieldID string) ([]*FieldBlock, error)
	//FindAvailable(ctx context.Context) ([]*FieldBlock, error)
	//FindByStatus(ctx context.Context, status valueobject.AreaStatus) ([]*FieldBlock, error)
	//FindByCropCycleID(ctx context.Context, cropCycleID string) (*FieldBlock, error)
}
