package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/core/port/repository"
	domain2 "samurenkoroma/services/internal/modules/farm/domain"
	"samurenkoroma/services/internal/modules/farm/domain/greenhouse"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
)

type CreateGreenhouseHandler struct {
	UowFactory repository.Factory
}

func NewCreateGreenhouseHandler(uowFactory repository.Factory) *CreateGreenhouseHandler {
	return &CreateGreenhouseHandler{UowFactory: uowFactory}
}

type CreateGreenhouseCmd struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
}

func DecodeCreateField(data []byte) (any, error) {

	var cmd CreateGreenhouseCmd

	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (h *CreateGreenhouseHandler) Handle(ctx context.Context, cmd any) error {

	c, ok := cmd.(CreateGreenhouseCmd)
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

	unit := greenhouse.NewGreenhouse(domain2.GrowingAreaID(c.ID), c.Name, dim)

	if err := uowObj.GrowingFacilities().Save(unit); err != nil {
		return err
	}
	uowObj.RegisterAggregate(unit)
	return uowObj.Commit()
}
