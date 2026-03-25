package queries

import (
	"context"
	"samurenkoroma/services/internal/modules/crop/infrastructure/persistence/projections"
)

// GetVarietyQuery — параметры запроса сорта
type GetVarietyQuery struct {
	ID string `json:"id" validate:"required"`
}

// GetVarietyHandler — обработчик запроса
type GetVarietyHandler struct {
	projector *projections.CropProjection
}

func (h *GetVarietyHandler) Name() string {
	return "GetVarieties"
}

func NewGetVarietyHandler(projector *projections.CropProjection) *GetVarietyHandler {
	return &GetVarietyHandler{
		projector: projector,
	}
}

func (h *GetVarietyHandler) Handle(ctx context.Context, q any) (any, error) {
	// Получаем сорт
	varieties, err := h.projector.GetVarieties(ctx)
	if err != nil {
		return nil, err
	}

	return varieties, nil
}
