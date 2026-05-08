package cropplan

import "context"

type Repository interface {
	Save(context.Context, *Plan) error
	GetByID(context.Context, string) (*Plan, error)
	All(context.Context, Filter) ([]*Plan, error)
}

type Filter struct {
	SeasonID *string
	AreaID   *string
	OwnerID  *string
	Status   *string
}
