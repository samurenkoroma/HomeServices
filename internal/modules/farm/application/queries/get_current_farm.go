package queries

import (
	"context"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/auth/domain"
)

type GetCurrentFarmQuery struct {
	Id string `json:"id,omitempty"`
}

func (h *QueryHandler) CurrentFarm(ctx context.Context, payload any) (any, error) {
	q, ok := payload.(*GetCurrentFarmQuery)
	if !ok {
		return nil, query.ErrInvalidPayloadType
	}

	if q.Id != "" {
		return h.projector.GetByID(ctx, q.Id)
	}

	farm, err := domain.NewOrganization("polevod", types.NewUUID(), "", "", "")
	return farm, err
}
