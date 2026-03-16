package commands

import (
	"context"
	"fmt"
	"samurenkoroma/services/internal/application/command"
	uow "samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/core/spatial"
)

type CreateBedCmd struct {
	Name     string          `json:"name"`
	Length   float64         `json:"length"`
	Width    float64         `json:"width"`
	Geom     spatial.GeoJSON `json:"geom"`
	ParentId string          `json:"parent_id"`
}

type createBedHandler struct {
	uowFactory uow.Factory
}

func NewCreateBedHandler(uowFactory uow.Factory) command.Handler {
	return &createBedHandler{uowFactory: uowFactory}
}

func (h *createBedHandler) Handle(ctx context.Context, cmd any) error {
	fmt.Print("create bed")
	return nil
}
