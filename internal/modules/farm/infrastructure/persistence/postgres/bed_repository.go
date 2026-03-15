package postgres

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/modules/farm/domain"
	"samurenkoroma/services/internal/modules/farm/domain/bed"
)

type bedRepository struct {
	tx *sql.Tx
}

func (b *bedRepository) Get(id domain.GrowingAreaID) (*bed.Bed, error) {
	//TODO implement me
	panic("implement me")
}

func (b *bedRepository) Save(unit *bed.Bed) error {
	query := `
    INSERT INTO land_structure(id,name,geom)
    VALUES (
        $1,
        $2,
        ST_SetSRID(ST_GeomFromGeoJSON($3),4326)
    )
    `

	_, err := b.tx.ExecContext(
		context.Background(),
		query,
		unit.ID(),
		unit.Name(),
		unit.Geometry().Coordinates,
	)

	return err

}

func NewBedRepository(tx *sql.Tx) bed.Repository {
	return &bedRepository{tx: tx}
}
