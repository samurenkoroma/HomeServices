package cropplan

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
)

type CultivationAreasQuery struct {
}

func (h *QueryHandler) GetCultivationAreas(ctx context.Context, query any) (any, error) {
	_, ok := query.(*CultivationAreasQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}
	return h.provider.CultivationAreas().FindAllBy(ctx, cultivationarea.Filter{})
}
