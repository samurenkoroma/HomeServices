package command

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/core/port/repository"
	facility2 "samurenkoroma/services/internal/modules/growing/domain/facility"
	"samurenkoroma/services/internal/modules/growing/domain/valueobject"
)

type CreateFacilityHandler struct {
	UowFactory repository.Factory
}

func NewCreateFacilityHandler(uowFactory repository.Factory) *CreateFacilityHandler {
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

	uowObj, err := h.UowFactory.Begin(ctx)
	if err != nil {
		return err
	}

	defer uowObj.Rollback()

	dim, err := valueobject.NewDimension(c.Length, c.Width)
	if err != nil {
		fmt.Println("dim create command error.", err)
		return err
	}
	var unit *facility2.GrowingFacility

	switch c.FacilityType {
	case "FIELD":
		unit = facility2.NewFieldFacility(facility2.FacilityID(c.ID), c.Name, dim)
	case "GREENHOUSE":
		unit = facility2.NewGreenhouseFacility(facility2.FacilityID(c.ID), c.Name, dim)
	default:
		return facility2.ErrInvalidSpaceType
	}

	if err := uowObj.GrowingFacilities().Save(unit); err != nil {
		return err
	}
	uowObj.RegisterAggregate(unit)
	return uowObj.Commit()
}
