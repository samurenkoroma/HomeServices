package postgres

import (
	"database/sql"
	"fmt"
	"samurenkoroma/services/internal/growing/application"
	"samurenkoroma/services/internal/growing/domain/facility"
)

type GrowingFacilitiesRepoImp struct {
	tx *sql.Tx
}

func NewGrowingFacilitiesRepository(tx *sql.Tx) application.GrowingFacilitiesRepository {
	return &GrowingFacilitiesRepoImp{tx: tx}
}
func (r *GrowingFacilitiesRepoImp) Get(id facility.FacilityID) (*facility.GrowingFacility, error) {
	//TODO implement me
	panic("implement me")
}

func (r *GrowingFacilitiesRepoImp) Save(unit *facility.GrowingFacility) error {
	// remove old structure
	_, err := r.tx.Exec(`DELETE FROM land_structure WHERE root_id = $1`, unit.ID())

	if err != nil {
		fmt.Println("Error deleting land_structure", err)
		return err
	}

	// map aggregate → flat persistence rows
	rows := mapLandUnitToRows(unit)

	for _, row := range rows {
		_, err = r.tx.Exec(`
			INSERT INTO land_structure
			(id, parent_id, root_id, unit_type, space_type, name, length, width, created_at)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8, NOW())
		`,
			row.ID,
			row.ParentID,
			row.RootID,
			row.UnitType,
			row.SpaceType,
			row.Name,
			row.Length,
			row.Width,
		)

		if err != nil {
			if isUniqueViolation(err) {
				return facility.ErrLandUnitAlreadyExists
			}
			return err
		}
	}

	return nil
}
