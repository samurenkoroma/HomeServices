package postgres

import (
	"samurenkoroma/services/internal/field/domain/cropplan"
	"samurenkoroma/services/internal/field/domain/landunit"
	"samurenkoroma/services/internal/field/domain/valueobject"
)

type landUnitRow struct {
	ID       string
	RootID   string
	ParentID *string
	Name     string
	UnitType string
	LandType string
	Length   float64
	Width    float64
}
type sectionRow landUnitRow
type bedRow landUnitRow

type cropPlanRow struct {
	ID        string
	BedID     string
	CropName  string
	Status    string
	HarvestKg float64
}

// LAND UNIT → DB
func mapLandUnitToRows(unit *landunit.LandUnit) []landUnitRow {
	rootId := string(unit.ID())
	lu := landUnitRow{
		ID:       string(unit.ID()),
		RootID:   rootId,
		Name:     unit.Name(),
		LandType: string(unit.Type()),
		UnitType: "land_unit",
		Length:   unit.Dimension().Length,
		Width:    unit.Dimension().Width,
	}

	var rows = []landUnitRow{lu}

	for _, s := range unit.Sections() {
		rows = append(rows, landUnitRow{
			ID:       string(s.ID()),
			ParentID: &rootId,
			RootID:   rootId,
			UnitType: "section",
			Name:     s.Name(),
			Length:   s.Dimension().Length,
			Width:    s.Dimension().Width,
		})

		for _, b := range s.Beds() {
			rows = append(rows, landUnitRow{
				ID:       string(b.ID()),
				ParentID: &rootId,
				RootID:   rootId,
				UnitType: "bed",
				Name:     b.Name(),
				Length:   b.Dimension().Length,
				Width:    b.Dimension().Width,
			})
		}
	}

	// beds directly under greenhouse
	for _, b := range unit.Beds() {
		rows = append(rows, landUnitRow{
			ID:       string(b.ID()),
			ParentID: &rootId,
			RootID:   rootId,
			Name:     b.Name(),
			Length:   b.Dimension().Length,
			Width:    b.Dimension().Width,
		})
	}

	return rows
}

// DB → LAND UNIT
func mapRowsToLandUnit(
	lu landUnitRow,
	sections []sectionRow,
	beds []bedRow,
) (*landunit.LandUnit, error) {

	dim, err := valueobject.NewDimension(lu.Length, lu.Width)
	if err != nil {
		return nil, err
	}

	var unit *landunit.LandUnit

	switch lu.LandType {
	case "field":
		unit = landunit.NewField(
			landunit.LandUnitID(lu.ID),
			lu.Name,
			dim,
		)
	case "greenhouse":
		unit = landunit.NewGreenhouse(
			landunit.LandUnitID(lu.ID),
			lu.Name,
			dim,
		)
	default:
		return nil, landunit.ErrInvalidUnitType
	}

	// создаём map sections
	sectionMap := make(map[string]*landunit.Section)

	for _, s := range sections {
		sDim, _ := valueobject.NewDimension(s.Length, s.Width)

		sec := landunit.NewSection(
			landunit.SectionID(s.ID),
			s.Name,
			sDim,
		)

		unit.AddSection(sec)
		sectionMap[s.ID] = sec
	}

	// добавляем beds
	for _, b := range beds {
		bDim, _ := valueobject.NewDimension(b.Length, b.Width)
		bed := landunit.NewBed(
			landunit.BedID(b.ID),
			b.Name,
			bDim,
		)

		if b.ParentID != nil {
			sectionMap[*b.ParentID].AddBed(bed)
		} else {
			unit.AddBedToGreenhouse(bed)
		}
	}

	return unit, nil
}

// CropPlan → DB
func mapCropPlanToRow(p *cropplan.CropPlan) cropPlanRow {
	return cropPlanRow{
		ID:        string(p.ID()),
		BedID:     string(p.BedID()),
		CropName:  p.CropName(),
		Status:    string(p.Status()),
		HarvestKg: p.HarvestKg(),
	}
}

// DB → CropPlan
func mapRowToCropPlan(r cropPlanRow) *cropplan.CropPlan {

	plan := cropplan.New(
		cropplan.CropPlanID(r.ID),
		cropplan.BedID(r.BedID),
		r.CropName,
	)

	switch r.Status {
	case "growing":
		plan.StartGrowing()
	case "harvested":
		plan.StartGrowing()
		plan.Harvest(r.HarvestKg)
	}

	return plan
}
