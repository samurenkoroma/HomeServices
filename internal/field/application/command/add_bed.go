package command

import (
	"samurenkoroma/services/internal/field/application"
	"samurenkoroma/services/internal/field/domain/landunit"
	"samurenkoroma/services/internal/field/domain/valueobject"
)

type AddBed struct {
	Uow application.UnitOfWork
}

type AddBedCmd struct {
	LandUnitID string
	SectionID  string
	BedID      string
	Name       string
	Length     float64
	Width      float64
}

func (h *AddBed) Handle(cmd AddBedCmd) error {
	unit, err := h.Uow.LandUnits().Get(landunit.LandUnitID(cmd.LandUnitID))
	if err != nil {
		return err
	}

	dim, _ := valueobject.NewDimension(cmd.Length, cmd.Width)
	bed := landunit.NewBed(landunit.BedID(cmd.BedID), cmd.Name, dim)

	if cmd.SectionID != "" {
		err = unit.AddBedToSection(landunit.SectionID(cmd.SectionID), bed)
	} else {
		err = unit.AddBedToGreenhouse(bed)
	}
	if err != nil {
		return err
	}

	err = h.Uow.LandUnits().Save(unit)
	if err != nil {
		return err
	}

	return h.Uow.Commit()
}
