package cropplan

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
)

type CultivationPlansQuery struct {
	CropKey string `json:"cropKey"`
}

func (h *QueryHandler) GetCultivationPlan(ctx context.Context, query any) (any, error) {
	q, ok := query.(*CultivationPlansQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}
	return h.projector.Catalog().GetCultivationPlans(ctx, catalog.CultivationPlansFilter{CropKey: q.CropKey})
}
