package postgres

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/farm/domain/bed"
)

type bedRepository struct {
	tx *sql.Tx
}

func (b bedRepository) FindByID(ctx context.Context, id types.BedId) (*bed.Bed, error) {
	//TODO implement me
	panic("implement me")
}

func (b bedRepository) Save(ctx context.Context, b2 *bed.Bed) error {
	//TODO implement me
	panic("implement me")
}

func (b bedRepository) FindAll(ctx context.Context) ([]*bed.Bed, error) {
	//TODO implement me
	panic("implement me")
}

func NewBedRepository(tx *sql.Tx) bed.Repository {
	return &bedRepository{tx: tx}
}
