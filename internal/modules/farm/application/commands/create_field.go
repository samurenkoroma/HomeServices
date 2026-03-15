package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"samurenkoroma/services/internal/core/domain/repository"
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
	fmt.Print("create field")
	return nil
}
