package postgres

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/farm/domain/field"
)

type GrowingFacilitiesRepoImp struct {
	tx *sql.Tx
}

func (g GrowingFacilitiesRepoImp) Get(id types.GrowingAreaID) (*field.Field, error) {
	//TODO implement me
	panic("implement me")
}

func (g GrowingFacilitiesRepoImp) Save(unit *field.Field) error {
	//TODO implement me
	panic("implement me")
}

//	type persistenceLandRow struct {
//		ID       string
//		ParentID *string
//		UnitType string
//		Name     string
//		Length   float64
//		Width    float64
//	}
func NewGrowingFacilitiesRepository(tx *sql.Tx) field.Repository {
	return &GrowingFacilitiesRepoImp{tx: tx}
}

//func (r *GrowingFacilitiesRepoImp) Get(id farm.FacilityID) (*field.Field, error) {
//	rows, err := r.tx.Query(`
//		SELECT id, parent_id, unit_type,  name, length, width
//		FROM land_structure
//		WHERE root_id = $1
//	`, id)
//
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//
//	var rootRow persistenceLandRow
//	var blocks []persistenceLandRow
//	var beds []persistenceLandRow
//
//	for rows.Next() {
//		var row persistenceLandRow
//
//		err := rows.Scan(
//			&row.ID,
//			&row.ParentID,
//			&row.UnitType,
//			&row.Name,
//			&row.Length,
//			&row.Width,
//		)
//
//		if err != nil {
//			return nil, err
//		}
//
//		if row.ParentID == nil {
//			rootRow = row
//		} else {
//			if row.UnitType == string(farm.BlockFacility) {
//				blocks = append(blocks, row)
//			} else if row.UnitType == string(farm.BedFacility) {
//				beds = append(beds, row)
//			}
//		}
//	}
//	dim, err := valueobject.NewDimension(rootRow.Length, rootRow.Width)
//	if err != nil {
//		return nil, err
//	}
//	unit := field.RehydrateGrowingFacility(
//		farm.FacilityID(rootRow.ID),
//		rootRow.Name,
//		dim,
//		nil,
//		nil,
//	)
//	//
//	//// rebuild tree
//	sectionMap := make(map[string]*domain3.FieldBlock)
//	//
//	for _, s := range blocks {
//
//		sDim, _ := valueobject.NewDimension(s.Length, s.Width)
//
//		sec := domain3.RehydrateBlock(
//			domain.GrowingAreaID(s.ID),
//			s.Name,
//			sDim,
//		)
//
//		unit.RehydrateAddBlock(&sec)
//		sectionMap[s.ID] = &sec
//	}
//	//
//	for _, b := range beds {
//		bDim, _ := valueobject.NewDimension(b.Length, b.Width)
//
//		bed := domain2.RehydrateBed(
//			domain.GrowingAreaID(b.ID),
//			b.Name,
//			bDim,
//		)
//
//		if *b.ParentID != string(unit.ID()) {
//			unit.RehydrateAddBedToSection(domain.GrowingAreaID(*b.ParentID), bed)
//		} else {
//			unit.RehydrateAddBed(bed)
//		}
//	}
//
//	return unit, nil
//}
//
//func (r *GrowingFacilitiesRepoImp) Save(unit *field.Field) error {
//	// remove old structure
//	_, err := r.tx.Exec(`DELETE FROM land_structure WHERE root_id = $1`, unit.ID())
//
//	if err != nil {
//		fmt.Println("Error deleting land_structure", err)
//		return err
//	}
//
//	// map aggregate → flat persistence rows
//	rows := mapLandUnitToRows(unit)
//
//	for _, row := range rows {
//		_, err = r.tx.Exec(`
//			INSERT INTO land_structure
//			(id, parent_id, root_id, unit_type, name, length, width, created_at)
//			VALUES ($1,$2,$3,$4,$5,$6,$7, NOW())
//		`,
//			row.ID,
//			row.ParentID,
//			row.RootID,
//			row.UnitType,
//			row.Name,
//			row.Length,
//			row.Width,
//		)
//
//		if err != nil {
//			if isUniqueViolation(err) {
//				return domain.ErrLandUnitAlreadyExists
//			}
//			return err
//		}
//	}
//
//	return nil
//}
