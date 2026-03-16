package bed

import (
	"context"
	"samurenkoroma/services/internal/core/domain/types"
)

type Repository interface {
	FindByID(context.Context, types.BedId) (*Bed, error)
	Save(context.Context, *Bed) error
	FindAll(ctx context.Context) ([]*Bed, error)
}
