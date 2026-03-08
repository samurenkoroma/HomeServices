package postgres

import (
	"samurenkoroma/services/internal/modules/farm/field/domain"
)

type facilityRow struct {
	ID       string
	RootID   string
	ParentID *string
	Name     string
	UnitType string
	Length   float64
	Width    float64
}

// LAND UNIT → DB
func mapLandUnitToRows(unit *domain.Field) []facilityRow {
	rootId := string(unit.ID())
	lu := facilityRow{
		ID:       string(unit.ID()),
		RootID:   rootId,
		Name:     unit.Name(),
		UnitType: string(unit.FacilityType()),
		Length:   unit.Dimension().Length(),
		Width:    unit.Dimension().Width(),
	}

	var rows = []facilityRow{lu}
	//
	//for _, s := range unit.Blocks() {
	//	rows = append(rows, facilityRow{
	//		ID:        string(s.ID()),
	//		ParentID:  &rootId,
	//		RootID:    rootId,
	//		UnitType:  "section",
	//		SpaceType: string(unit.SpaceType()),
	//		Name:      s.Name(),
	//		Length:    s.Dimension().Length(),
	//		Width:     s.Dimension().Width(),
	//	})
	//
	//	for _, b := range s.Beds() {
	//		rows = append(rows, landUnitRow{
	//			ID:        string(b.ID()),
	//			ParentID:  &rootId,
	//			RootID:    rootId,
	//			SpaceType: string(unit.SpaceType()),
	//			UnitType:  "bed",
	//			Name:      b.Name(),
	//			Length:    b.Dimension().Length,
	//			Width:     b.Dimension().Width,
	//		})
	//	}
	//}
	//
	//// beds directly under greenhouse
	//for _, b := range unit.Beds() {
	//	rows = append(rows, landUnitRow{
	//		ID:        string(b.ID()),
	//		ParentID:  &rootId,
	//		RootID:    rootId,
	//		SpaceType: string(unit.SpaceType()),
	//		UnitType:  "bed",
	//		Name:      b.Name(),
	//		Length:    b.Dimension().Length,
	//		Width:     b.Dimension().Width,
	//	})
	//}

	return rows
}

//
//// DB → LAND UNIT
//func mapRowsToLandUnit(
//	lu landUnitRow,
//	sections []sectionRow,
//	beds []bedRow,
//) (*facility.Field, error) {
//
//	dim, err := valueobject.NewDimension(lu.Length, lu.Width)
//	if err != nil {
//		return nil, err
//	}
//
//	var unit *facility.Field
//
//	switch lu.SpaceType {
//	case "facility":
//		unit = facility.NewGreenhouseFacility(
//			facility.AreaID(lu.ID),
//			lu.Name,
//			dim,
//		)
//	case "greenhouse":
//		unit = facility.NewGreenhouseFacility(
//			facility.AreaID(lu.ID),
//			lu.Name,
//			dim,
//		)
//	default:
//		return nil, facility.ErrInvalidSpaceType
//	}
//
//	// создаём map sections
//	sectionMap := make(map[string]facility.FieldBlock)
//
//	for _, s := range sections {
//		sDim, _ := valueobject.NewDimension(s.Length, s.Width)
//
//		sec := facility.NewFieldBlock(
//			facility.GrowingAreaID(s.ID),
//			s.Name,
//			sDim,
//		)
//
//		unit.AddBlock(sec)
//		sectionMap[s.ID] = sec
//	}
//
//	// добавляем beds
//	for _, b := range beds {
//		bDim, _ := valueobject.NewDimension(b.Length, b.Width)
//		facility.NewBed(
//			facility.GrowingAreaID(b.ID),
//			b.Name,
//			bDim,
//		)
//
//		//if b.ParentID != nil {
//		//	sectionMap[*b.ParentID].AddBed(bed)
//		//} else {
//		//	unit.AddBed(bed)
//		//}
//	}
//
//	return unit, nil
//}
//
//// CropPlan → DB
//func mapCropPlanToRow(p *cropplan.CropPlan) cropPlanRow {
//	return cropPlanRow{
//		ID:        string(p.ID()),
//		BedID:     string(p.BedID()),
//		CropName:  p.CropName(),
//		Status:    string(p.Status()),
//		HarvestKg: p.HarvestKg(),
//	}
//}
//
//// DB → CropPlan
//func mapRowToCropPlan(r cropPlanRow) *cropplan.CropPlan {
//
//	plan := cropplan.New(
//		cropplan.CropPlanID(r.ID),
//		cropplan.BedID(r.BedID),
//		r.CropName,
//	)
//
//	switch r.Status {
//	case "facility":
//		plan.StartGrowing()
//	case "harvested":
//		plan.StartGrowing()
//		plan.Harvest(r.HarvestKg)
//	}
//
//	return plan
//}
