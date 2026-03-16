package commands

import (
	"context"
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

type createFieldBlockHandler struct {
	uowFactory uow.Factory
}

func NewCreateFieldBlockHandler(uowFactory uow.Factory) command.Handler {
	return &createFieldBlockHandler{uowFactory: uowFactory}
}

func (h *createFieldBlockHandler) Handle(ctx context.Context, cmd any) error {
	fmt.Print("create field block")
	return nil
}
