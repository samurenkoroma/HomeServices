package queries

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/modules/crop/domain/variety"
)

// GetVarietyQuery — параметры запроса сорта
type GetVarietyQuery struct {
	ID         string `json:"id" validate:"required"`
	CropTypeId string `json:"crop_type_id,omitempty"`
	IsActive   bool   `json:"is_active,omitempty"`
	Search     string `json:"search,omitempty"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
}

// GetVarietyHandler — обработчик запроса
type GetVarietyHandler struct {
	projector variety.Projections
}

func (h *GetVarietyHandler) Name() string {
	return "GetVarieties"
}

func NewGetVarietyHandler(projector variety.Projections) query.Handler {
	return &GetVarietyHandler{
		projector: projector,
	}
}

func (h *GetVarietyHandler) Handle(ctx context.Context, query any) (any, error) {
	// Получаем сорт
	q, ok := query.(*GetVarietyQuery)
	if !ok {
		return nil, errors.New("invalid query")
	}
	if q.ID != "" {
		return h.projector.GetByID(ctx, q.ID)
	}
	filter := variety.Filter{
		Search: q.Search,
		Limit:  q.Limit,
		Offset: q.Offset,
	}
	if q.IsActive {
		filter.IsActive = true
	}
	if q.CropTypeId != "" {
		filter.CropTypeId = q.CropTypeId
	}

	if q.Limit == 0 {
		filter.Limit = 10
	}
	return h.projector.GetList(ctx, filter)
}
