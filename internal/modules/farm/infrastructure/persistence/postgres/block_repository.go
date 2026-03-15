package postgres

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/modules/farm/domain"
	"samurenkoroma/services/internal/modules/farm/domain/block"
)

type persistenceLandRow struct {
	id       string
	parentID *string
	rootID   string
	unitType string
	name     string
	geom
	properties map[string]interface{}
}
type blockRepository struct {
	tx *sql.Tx
}

func (b *blockRepository) Get(id domain.GrowingAreaID) (*field_block.FieldBlock, error) {
	//TODO implement me
	panic("implement me")
}

func (b *blockRepository) Save(unit *field_block.FieldBlock) error {
	query := `
    INSERT INTO land_structure(id,name,geom, properties)
    VALUES (
        $1,
        $2,
        ST_SetSRID(ST_GeomFromGeoJSON($3),4326),
		$4
    )
    `
	row := mapBedToRow(unit)
	_, err := b.tx.ExecContext(
		context.Background(),
		query,
		row.id,
		row.name,
		unit.Geometry.Coordinates,
		unit.P
	)

	return err
	return nil

}

func mapBedToRow(unit *field_block.FieldBlock) persistenceLandRow {

}

func NewBlockRepository(tx *sql.Tx) field_block.Repository {
	return &blockRepository{tx: tx}
}
