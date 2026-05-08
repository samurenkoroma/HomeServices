package cultivation

import "context"

type Repository interface {
	Save(ctx context.Context, p *CultivationPlan) error
	Get(ctx context.Context, id string, version int) (*CultivationPlan, error)
	GetLatest(ctx context.Context, id string) (*CultivationPlan, error)
	List(ctx context.Context, cropKey string) ([]*CultivationPlan, error)
}
