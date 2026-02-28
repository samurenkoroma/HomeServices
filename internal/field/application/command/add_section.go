package command

import (
	"samurenkoroma/services/internal/field/application"
	"samurenkoroma/services/internal/field/domain/landunit"
	"samurenkoroma/services/internal/field/domain/valueobject"
)

type AddSection struct {
	Uow application.UnitOfWork
}

type AddSectionCmd struct {
	LandUnitID string
	SectionID  string
	Name       string
	Length     float64
	Width      float64
}

func (h *AddSection) Handle(cmd AddBedCmd) error {
	unit, err := h.Uow.LandUnits().Get(landunit.LandUnitID(cmd.LandUnitID))
	if err != nil {
		return err
	}

	dim, _ := valueobject.NewDimension(cmd.Length, cmd.Width)
	section := landunit.NewSection(landunit.SectionID(cmd.SectionID), cmd.Name, dim)

	err = unit.AddSection(section)
	if err != nil {
		return err
	}

	err = h.Uow.LandUnits().Save(unit)
	if err != nil {
		return err
	}

	return h.Uow.Commit()
}
