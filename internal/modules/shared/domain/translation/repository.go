package translation

import "context"

type Repository interface {
	Save(ctx context.Context, entity, latin, ru string) error
}
