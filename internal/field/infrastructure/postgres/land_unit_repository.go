package postgres

import (
	"database/sql"
	"fmt"
	"samurenkoroma/services/internal/field/application"
	"samurenkoroma/services/internal/field/domain/landunit"
)

type LandUnitRepoImp struct {
	tx  *sql.Tx
	uow *PgUow
}

func NewLandUnitRepository(tx *sql.Tx, uow *PgUow) application.LandUnitRepository {
	return &LandUnitRepoImp{tx: tx, uow: uow}
}

func (r *LandUnitRepoImp) Save(unit *landunit.LandUnit) error {
	// remove old structure
	_, err := r.tx.Exec(`
		DELETE FROM land_structure
		WHERE root_id = $1
	`, unit.ID())

	if err != nil {
		fmt.Println("Error deleting land_structure", err)
		return err
	}

	// map aggregate → flat persistence rows
	rows := mapLandUnitToRows(unit)

	for _, row := range rows {
		_, err = r.tx.Exec(`
			INSERT INTO land_structure
			(id, parent_id, root_id, unit_type, land_type, name, length, width, created_at)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8, NOW())
		`,
			row.ID,
			row.ParentID,
			row.RootID,
			row.UnitType,
			row.LandType,
			row.Name,
			row.Length,
			row.Width,
		)

		if err != nil {
			if isUniqueViolation(err) {
				return landunit.ErrLandUnitAlreadyExists
			}
			return err
		}
	}

	r.uow.registerAggregate(unit)
	return nil
}
