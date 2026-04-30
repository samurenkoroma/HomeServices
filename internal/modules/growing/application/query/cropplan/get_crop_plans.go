package cropplan

import (
	"context"
	"errors"
)

type GetCropPlanQuery struct {
	PlanID   string `json:"planId"`
	ObjectID string `json:"objectID,omitempty"`
}

func (h *QueryHandler) GetCropPlans(ctx context.Context, query any) (any, error) {
	q, ok := query.(*GetCropPlanQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	if q.PlanID != "" {
		return h.PlanRepo.FindByID(ctx, q.PlanID)
	}
	return h.PlanRepo.FindByObject(ctx, q.ObjectID)
}
