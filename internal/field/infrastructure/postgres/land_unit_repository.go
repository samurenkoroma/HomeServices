package postgres

import (
	"database/sql"
	"samurenkoroma/services/internal/field/application"
	"samurenkoroma/services/internal/field/domain/landunit"
	"samurenkoroma/services/internal/field/domain/valueobject"
)

type LandUnitRepoImp struct {
	db *sql.DB
}

func (r *LandUnitRepoImp) Get(id landunit.LandUnitID) (*landunit.LandUnit, error) {

	rows, err := r.db.Query(`
		SELECT id, parent_id, unit_type, name, length, width
		FROM land_structure
		WHERE root_id = $1
	`, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rootRow landUnitRow
	var sections []landUnitRow
	var beds []landUnitRow

	for rows.Next() {
		var row landUnitRow

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
			if row.UnitType == "section" {
				sections = append(sections, row)
			} else {
				beds = append(beds, row)
			}
		}
	}

	dim, err := valueobject.NewDimension(rootRow.Length, rootRow.Width)
	if err != nil {
		return nil, err
	}

	unit := landunit.RehydrateLandUnit(
		landunit.LandUnitID(rootRow.ID),
		rootRow.Name,
		landunit.LandUnitType(rootRow.UnitType),
		dim,
		nil,
		nil,
	)

	// rebuild tree
	//sectionMap := make(map[string]*landunit.Section)

	//for _, s := range sections {
	//
	//	sDim, _ := valueobject.NewDimension(s.Length, s.Width)
	//
	//	sec := landunit.RehydrateSection(
	//		landunit.SectionID(s.ID),
	//		s.Name,
	//		sDim,
	//	)
	//
	//	unit.RehydrateAddSection(sec)
	//	sectionMap[s.ID] = sec
	//}
	//
	//for _, b := range beds {
	//
	//	bDim, _ := valueobject.NewDimension(b.Length, b.Width)
	//
	//	bed := landunit.RehydrateBed(
	//		landunit.BedID(b.ID),
	//		b.Name,
	//		bDim,
	//	)
	//
	//	if b.ParentID != nil {
	//		if sec, ok := sectionMap[*b.ParentID]; ok {
	//			sec.RehydrateAddBed(bed)
	//		}
	//	} else {
	//		unit.RehydrateAddBed(bed)
	//	}
	//}

	return unit, nil
}

func (r *LandUnitRepoImp) Save(unit *landunit.LandUnit) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// remove old structure
	_, err = tx.Exec(`
		DELETE FROM land_structure
		WHERE root_id = $1
	`, unit.ID())

	if err != nil {
		return err
	}

	// map aggregate → flat persistence rows
	rows := mapLandUnitToRows(unit)

	for _, row := range rows {
		_, err = tx.Exec(`
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
			return err
		}
	}

	return tx.Commit()
}

func NewLandRepo(db *sql.DB) application.LandUnitRepository {
	return &LandUnitRepoImp{
		db: db,
	}
}
