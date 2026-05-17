package cropplan

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
)

type CropPlansQuery struct {
	ObjectID string `json:"objectID,omitempty"`
}

func (h *QueryHandler) GetCropPlans(ctx context.Context, query any) (any, error) {
	q, ok := query.(*CropPlansQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}
	orgId, ok := ctx.Value("organization_id").(string)
	filter := catalog.CropPlanFilter{
		OwnerId: orgId,
	}

	if q.ObjectID != "" {
		filter.ProductionUnitId = q.ObjectID
	}

	return h.projector.Catalog().GetCropPlans(ctx, filter)
}
