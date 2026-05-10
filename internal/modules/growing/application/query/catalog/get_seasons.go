package catalog

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
)

type SeasonsQuery struct {
}

func (h *QueryHandler) GetSeasons(ctx context.Context, query any) (any, error) {
	_, ok := query.(*SeasonsQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}
	orgId, ok := ctx.Value("organization_id").(string)
	if !ok {
		return nil, errors.New("organization_id is required")
	}

	return h.catalog.GetSeasons(ctx, catalog.SeasonFilter{
		OwnerId: orgId,
	})
}
