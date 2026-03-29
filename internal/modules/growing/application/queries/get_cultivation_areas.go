package queries

import (
	"context"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
)

type GetCultivationAreasQuery struct {
	Id     string `json:"id,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty"`
}

type getCultivationAreasHandler struct {
	projector cultivationarea.Projections
}

func (h *getCultivationAreasHandler) Name() string {
	return "GetCultivationAreas"
}

func NewGetCultivationAreasHandler(projector cultivationarea.Projections) query.Handler {
	return &getCultivationAreasHandler{
		projector: projector,
	}
}

func (h *getCultivationAreasHandler) Handle(ctx context.Context, payload any) (any, error) {
	q, ok := payload.(*GetCultivationAreasQuery)
	if !ok {
		return nil, query.ErrInvalidPayloadType
	}

	//if q.Id != "" {
	//	return h.projector.GetList(ctx, )
	//}

	filter := cultivationarea.Filter{
		Limit:  q.Limit,
		Offset: q.Offset,
	}
	if q.Limit == 0 {
		filter.Limit = 10
	}
	return h.projector.GetList(ctx, filter)
}
