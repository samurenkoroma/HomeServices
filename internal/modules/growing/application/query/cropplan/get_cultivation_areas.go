package cropplan

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
)

type CultivationAreasQuery struct {
	ObjectId string `json:"objectId"`
}

func (h *QueryHandler) GetCultivationAreas(ctx context.Context, query any) (any, error) {
	c, ok := query.(*CultivationAreasQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	filter := cultivationarea.AreaFilter{}

	if c.ObjectId != "" {
		filter.ObjectId = c.ObjectId
	}

	return h.projector.Areas().GetList(ctx, filter)
}
