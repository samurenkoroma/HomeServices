package cultivationarea

import (
	"context"
)

type Filter struct {
	FarmId   string
	Type     AreaType
	ParentId string
}

type Repository interface {
	Save(ctx context.Context, area CultivationArea) error
	FindBy(ctx context.Context, id string) (CultivationArea, error)
	FindAllBy(context.Context, Filter) ([]CultivationArea, error)
}
