package repository

import (
	"context"
)

type Factory interface {
	Begin(ctx context.Context) (UnitOfWork, error)
}
