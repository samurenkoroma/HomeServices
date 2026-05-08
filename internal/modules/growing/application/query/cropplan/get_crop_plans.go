package cropplan

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
)

type CropPlansQuery struct {
	PlanID   string `json:"planId"`
	ObjectID string `json:"objectID,omitempty"`
}

func (h *QueryHandler) GetCropPlans(ctx context.Context, query any) (any, error) {
	q, ok := query.(*CropPlansQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}
	orgId, ok := ctx.Value("organization_id").(string)
	if q.PlanID != "" {
		return h.provider.CropPlans().GetByID(ctx, q.PlanID)
	}
	return h.provider.CropPlans().All(ctx, cropplan.Filter{
		OwnerID: &orgId,
	})
}
