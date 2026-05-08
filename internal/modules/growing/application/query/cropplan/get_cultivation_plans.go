package cropplan

import (
	"context"
	"errors"
)

type CultivationPlansQuery struct {
	CropKey string `json:"cropKey"`
}

func (h *QueryHandler) GetCultivationPlan(ctx context.Context, query any) (any, error) {
	q, ok := query.(*CultivationPlansQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}
	//if q.CropKey != "" {
	return h.provider.Cultivation().List(ctx, q.CropKey)
	//}
	//return h.provider.All(ctx, cropplan.Filter{
	//	OwnerID: &orgId,
	//})
}
