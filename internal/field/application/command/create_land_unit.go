package command

import (
	"samurenkoroma/services/internal/field/application"
	"samurenkoroma/services/internal/field/domain/landunit"
	"samurenkoroma/services/internal/field/domain/valueobject"
)

type CreateLandUnitHandler struct {
	Uow application.UnitOfWork
}

type CreateLandUnitCmd struct {
	ID     string
	Name   string
	Type   string
	Length float64
	Width  float64
}

func (h *CreateLandUnitHandler) Handle(cmd CreateLandUnitCmd) error {
	dim, err := valueobject.NewDimension(cmd.Length, cmd.Width)
	if err != nil {
		return err
	}

	var unit *landunit.LandUnit

	switch cmd.Type {
	case "field":
		unit = landunit.NewField(landunit.LandUnitID(cmd.ID), cmd.Name, dim)
	case "greenhouse":
		unit = landunit.NewGreenhouse(landunit.LandUnitID(cmd.ID), cmd.Name, dim)
	default:
		return landunit.ErrInvalidUnitType
	}

	err = h.Uow.LandUnits().Save(unit)
	if err != nil {
		return err
	}

	return h.Uow.Commit()
}
