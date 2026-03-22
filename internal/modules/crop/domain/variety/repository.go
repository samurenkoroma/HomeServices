package variety

import "context"

type Repository interface {
	Save(ctx context.Context, obj *Variety) error
	GetByID(ctx context.Context, id VarietyID) (*Variety, error)
	List(ctx context.Context) ([]*Variety, error)
}
