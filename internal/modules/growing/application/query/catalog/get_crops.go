package catalog

import (
	"context"
	"errors"
)

type CropsQuery struct {
	Key string `form:"key,omitempty"`
}

func (h *QueryHandler) GetCrops(ctx context.Context, query any) (any, error) {
	q, ok := query.(*CropsQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	if q.Key != "" {
		return h.catalog.GetCrop(ctx, q.Key)
	}
	return h.catalog.ListCrops(ctx)
}
