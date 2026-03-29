package queries

import (
	"context"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/modules/crop/domain/cropplan"
)

type GetCropPlanQuery struct {
	Id     string `json:"id,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty"`
}
type getCropPlanHandler struct {
	projector cropplan.Projections
}

func (h *getCropPlanHandler) Handle(ctx context.Context, payload any) (any, error) {
	q, ok := payload.(*GetCropPlanQuery)
	if !ok {
		return nil, query.ErrInvalidPayloadType
	}

	if q.Id != "" {
		return h.projector.GetByID(ctx, q.Id)
	}

	filter := cropplan.Filter{
		Limit:  q.Limit,
		Offset: q.Offset,
	}
	if q.Limit == 0 {
		filter.Limit = 10
	}
	return h.projector.GetList(ctx, filter)
}

func (h *getCropPlanHandler) Name() string {
	return "GetCropPlans"
}

func NewGetCropPlanHandler(projector cropplan.Projections) query.Handler {
	return &getCropPlanHandler{
		projector: projector,
	}
}
