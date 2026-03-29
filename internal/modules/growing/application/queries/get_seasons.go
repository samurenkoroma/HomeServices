package queries

import (
	"context"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/modules/growing/domain/season"
)

type GetSeasonsQuery struct {
	Id     string `json:"id,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty"`
}

type getSeasonsHandler struct {
	projector season.Projections
}

func (h *getSeasonsHandler) Handle(ctx context.Context, payload any) (any, error) {
	q, ok := payload.(*GetSeasonsQuery)
	if !ok {
		return nil, query.ErrInvalidPayloadType
	}

	if q.Id != "" {
		return h.projector.GetByID(ctx, q.Id)
	}

	filter := season.Filter{
		Limit:  q.Limit,
		Offset: q.Offset,
	}
	if q.Limit == 0 {
		filter.Limit = 10
	}
	return h.projector.GetList(ctx, filter)
}

func (h *getSeasonsHandler) Name() string { return "GetSeasons" }

func NewGetSeasons(projector season.Projections) query.Handler {
	return &getSeasonsHandler{projector: projector}
}
