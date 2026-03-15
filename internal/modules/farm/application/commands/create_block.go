package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	uow "samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/spatial"
)

type CreateFieldBlockCmd struct {
	Name     string          `json:"name"`
	Geom     spatial.GeoJSON `json:"geom"`
	ParentId string          `json:"parent_id"`
}

func NewCreateFieldBlockHandler(uowFactory uow.Factory) command.Handler {
	return &CreateFieldHandler{UowFactory: uowFactory}
}

type createFieldBlockHandler struct {
	uowFactory uow.Factory
}

func DecodeCreateFieldBlock(data []byte) (any, error) {

	var cmd CreateFieldBlockCmd

	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (h *createFieldBlockHandler) Handle(ctx context.Context, cmd any) error {
	fmt.Print("create field block")
	return nil
}
