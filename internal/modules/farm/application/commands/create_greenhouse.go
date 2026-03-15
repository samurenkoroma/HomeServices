package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/core/domain/repository"
)

type CreateGreenhouseCmd struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
}

type createGreenhouseHandler struct {
	UowFactory repository.Factory
}

func NewCreateGreenhouseHandler(uowFactory repository.Factory) command.Handler {
	return &createGreenhouseHandler{UowFactory: uowFactory}
}

func DecodeCreateGreenhouse(data []byte) (any, error) {

	var cmd CreateGreenhouseCmd

	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (h *createGreenhouseHandler) Handle(ctx context.Context, cmd any) error {
	fmt.Print("create greenhouse")
	return nil
}
