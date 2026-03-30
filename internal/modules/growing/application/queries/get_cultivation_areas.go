package queries

import (
	"context"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
)

type GetCultivationAreasQuery struct {
	Id       string `json:"id,omitempty"`
	Type     string `json:"type"`      // field, block, greenhouse, bed
	SeasonID string `json:"season_id"` // фильтр по сезону
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

	if q.Id != "" {
		return h.projector.GetByID(ctx, q.Id)
	}

	filter := cultivationarea.AreaFilter{
		Type:     q.Type,
		SeasonID: q.SeasonID,
	}

	return h.projector.GetList(ctx, filter)
}
