package postgres

import (
	"database/sql"
	"fmt"
	"samurenkoroma/services/internal/growing/application"
	"samurenkoroma/services/internal/growing/domain/facility"
	"samurenkoroma/services/internal/growing/domain/valueobject"
)

type GrowingFacilitiesRepoImp struct {
	tx *sql.Tx
}

type persistenceLandRow struct {
	ID       string
	ParentID *string
	UnitType string
	Name     string
	Length   float64
	Width    float64
}

func NewGrowingFacilitiesRepository(tx *sql.Tx) application.GrowingFacilitiesRepository {
	return &GrowingFacilitiesRepoImp{tx: tx}
}
func (r *GrowingFacilitiesRepoImp) Get(id facility.FacilityID) (*facility.GrowingFacility, error) {
	rows, err := r.tx.Query(`
		SELECT id, parent_id, unit_type,  name, length, width
		FROM land_structure
		WHERE root_id = $1
	`, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rootRow persistenceLandRow
	var blocks []persistenceLandRow
	var beds []persistenceLandRow

	for rows.Next() {
		var row persistenceLandRow

		err := rows.Scan(
			&row.ID,
			&row.ParentID,
			&row.UnitType,
			&row.Name,
			&row.Length,
			&row.Width,
		)

		if err != nil {
			return nil, err
		}

		if row.ParentID == nil {
			rootRow = row
		} else {
			if row.UnitType == string(facility.BlockFacility) {
				blocks = append(blocks, row)
			} else if row.UnitType == string(facility.BedFacility) {
				beds = append(beds, row)
			}
		}
	}
	dim, err := valueobject.NewDimension(rootRow.Length, rootRow.Width)
	if err != nil {
		return nil, err
	}
	unit := facility.RehydrateGrowingFacility(
		facility.FacilityID(rootRow.ID),
		rootRow.Name,
		facility.FacilityType(rootRow.UnitType),
		dim,
		nil,
		nil,
	)
	//
	//// rebuild tree
	sectionMap := make(map[string]*facility.FieldBlock)
	//
	for _, s := range blocks {

		sDim, _ := valueobject.NewDimension(s.Length, s.Width)

		sec := facility.RehydrateBlock(
			facility.GrowingAreaID(s.ID),
			s.Name,
			sDim,
		)

		unit.RehydrateAddBlock(&sec)
		sectionMap[s.ID] = &sec
	}
	//
	for _, b := range beds {
		bDim, _ := valueobject.NewDimension(b.Length, b.Width)

		bed := facility.RehydrateBed(
			facility.GrowingAreaID(b.ID),
			b.Name,
			bDim,
		)

		if *b.ParentID != string(unit.ID()) {
			unit.RehydrateAddBedToSection(facility.GrowingAreaID(*b.ParentID), bed)
		} else {
			unit.RehydrateAddBed(bed)
		}
	}

	return unit, nil
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
			(id, parent_id, root_id, unit_type, name, length, width, created_at)
			VALUES ($1,$2,$3,$4,$5,$6,$7, NOW())
		`,
			row.ID,
			row.ParentID,
			row.RootID,
			row.UnitType,
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
