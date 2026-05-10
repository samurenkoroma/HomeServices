package catalog

import (
	"context"
	"errors"
)

type VarietiesQuery struct {
	CropKey string `json:"cropKey,omitempty"` // tomato, eggplant, cucumber
	Id      string `form:"id,omitempty"`
}

func (h *QueryHandler) GetVarieties(ctx context.Context, query any) (any, error) {
	q, ok := query.(*VarietiesQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	if q.Id != "" {
		return h.catalog.GetVariety(ctx, q.Id)
	}

	return h.catalog.GetVarieties(ctx, q.CropKey)

}
