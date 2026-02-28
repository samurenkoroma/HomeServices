package command

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/field/application"
	"samurenkoroma/services/internal/field/domain/landunit"
	"samurenkoroma/services/internal/field/domain/valueobject"
)

type CreateLandUnitHandler struct {
	UowFactory application.UnitOfWorkFactory
}

func NewCreateLandUnitHandler(uowFactory application.UnitOfWorkFactory) *CreateLandUnitHandler {
	return &CreateLandUnitHandler{UowFactory: uowFactory}
}

type CreateLandUnitCmd struct {
	ID     string
	Name   string
	Type   string
	Length float64
	Width  float64
}

func (h *CreateLandUnitHandler) Handle(cmd CreateLandUnitCmd) error {
	uow, err := h.UowFactory.New(context.Background())
	if err != nil {
		return err
	}
	defer uow.Rollback()

	dim, err := valueobject.NewDimension(cmd.Length, cmd.Width)
	if err != nil {
		fmt.Println("dimcreate command error.", err)
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

	err = uow.LandUnits().Save(unit)
	if err != nil {
		fmt.Println("landunit save command error.", err)
		return err
	}

	return uow.Commit()
}
