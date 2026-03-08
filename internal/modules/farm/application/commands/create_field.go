package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/core/port/repository"
	domain2 "samurenkoroma/services/internal/modules/farm/domain"
	"samurenkoroma/services/internal/modules/farm/domain/field"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
)

type CreateFieldHandler struct {
	UowFactory repository.Factory
}

func NewCreateFieldHandler(uowFactory repository.Factory) *CreateFieldHandler {
	return &CreateFieldHandler{UowFactory: uowFactory}
}

type CreateFieldCmd struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
}

func DecodeCreateField(data []byte) (any, error) {

	var cmd CreateFieldCmd

	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (h *CreateFieldHandler) Handle(ctx context.Context, cmd any) error {

	c, ok := cmd.(CreateFieldCmd)
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

	unit := field.NewField(domain2.GrowingAreaID(c.ID), c.Name, dim)

	if err := uowObj.GrowingFacilities().Save(unit); err != nil {
		return err
	}
	uowObj.RegisterAggregate(unit)
	return uowObj.Commit()
}
