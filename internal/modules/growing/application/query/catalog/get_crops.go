package catalog

import (
	"context"
)

type CropsQuery struct {
	Key string `form:"key,omitempty"`
}

func (h *QueryHandler) GetCrops(ctx context.Context, query any) (any, error) {
	return h.catalog.GetCrops(ctx)
}
