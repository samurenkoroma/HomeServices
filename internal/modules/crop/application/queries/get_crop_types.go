package queries

import (
	"context"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/modules/crop/domain/croptype"

	"github.com/olekukonko/errors"
)

// GetCropTypesQuery — параметры запроса списка типов культур
type GetCropTypesQuery struct {
	ID            string `json:"id"`
	Category      string `json:"category,omitempty"`
	Family        string `json:"family,omitempty"`
	ActiveOnly    bool   `json:"active_only"`
	Search        string `json:"search"`
	WithVarieties bool   `json:"with_varieties"`
}

// GetCropTypesHandler — обработчик запроса
type getCropTypesHandler struct {
	projector croptype.Projections
}

func (h *getCropTypesHandler) Name() string {
	return "GetCropTypes"
}

func NewGetCropTypesHandler(projector croptype.Projections) query.Handler {
	return &getCropTypesHandler{projector: projector}
}

func (h *getCropTypesHandler) Handle(ctx context.Context, payload any) (any, error) {
	q, ok := payload.(*GetCropTypesQuery)
	if !ok {
		return nil, errors.New("invalid payload type")
	}

	if q.ID != "" {
		return h.projector.GetCropTypeWithVarieties(ctx, q.ID)
	}

	filter := croptype.Filter{
		Search: q.Search,
	}

	if q.Category != "" {
		filter.Category = q.Category
	}
	if q.ActiveOnly {
		filter.IsActive = true
	}

	return h.projector.GetList(ctx, filter)
}
