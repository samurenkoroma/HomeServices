package greenhouse

import (
	"context"
	"samurenkoroma/services/internal/core/domain/types"
)

type Repository interface {
	FindByID(context.Context, types.GreenhouseId) (*Greenhouse, error)
	Save(context.Context, *Greenhouse) error
	FindAll(ctx context.Context) ([]*Greenhouse, error)
}
