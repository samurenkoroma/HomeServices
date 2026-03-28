package queries

import (
	"context"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/modules/crop/domain/croptype"

	"github.com/olekukonko/errors"
)

// GetCropTypesQuery — параметры запроса списка типов культур
type GetCropTypesQuery struct {
	ID         string `json:"id"`
	Category   string `json:"category"`
	ActiveOnly bool   `json:"active_only"`
	Search     string `json:"search"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
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
		return h.projector.GetByID(ctx, q.ID)
	}

	filter := croptype.Filter{
		Search: q.Search,
		Limit:  q.Limit,
		Offset: q.Offset,
	}
	return h.projector.GetList(ctx, filter)
}
