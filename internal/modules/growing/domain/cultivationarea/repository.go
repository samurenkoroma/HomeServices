package cultivationarea

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, area CultivationArea) error
	FindById(ctx context.Context, id string) (CultivationArea, error)
}
