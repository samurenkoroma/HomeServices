package command

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/common/application/uow"
	"samurenkoroma/services/internal/growing/domain/facility"
	"samurenkoroma/services/internal/growing/domain/valueobject"
)

type CreateFacilityHandler struct {
	UowFactory uow.Factory
}

func NewCreateFacilityHandler(uowFactory uow.Factory) *CreateFacilityHandler {
	return &CreateFacilityHandler{UowFactory: uowFactory}
}

type CreateFacilityCmd struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	FacilityType string  `json:"type"`
	Length       float64 `json:"length"`
	Width        float64 `json:"width"`
}

func DecodeCreateField(data []byte) (any, error) {

	var cmd CreateFacilityCmd

	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (h *CreateFacilityHandler) Handle(ctx context.Context, cmd any) error {

	c, ok := cmd.(CreateFacilityCmd)
	if !ok {
		return errors.New("invalid command type")
	}

	uow, err := h.UowFactory.Begin(ctx)
	if err != nil {
		return err
	}

	defer uow.Rollback()

	dim, err := valueobject.NewDimension(c.Length, c.Width)
	if err != nil {
		fmt.Println("dim create command error.", err)
		return err
	}
	var unit *facility.GrowingFacility

	switch c.FacilityType {
	case "FIELD":
		unit = facility.NewFieldFacility(facility.FacilityID(c.ID), c.Name, dim)
	case "GREENHOUSE":
		unit = facility.NewGreenhouseFacility(facility.FacilityID(c.ID), c.Name, dim)
	default:
		return facility.ErrInvalidSpaceType
	}

	if err := uow.GrowingFacilities().Save(unit); err != nil {
		return err
	}
	uow.RegisterAggregate(unit)
	return uow.Commit()
}
