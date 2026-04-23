package queries

import (
	"context"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/auth/domain"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
)

type GetCurrentFarmQuery struct {
	Id string `json:"id,omitempty"`
}
type getCurrentFarmHandler struct {
	projector physicalobject.ObjectProjections
}

func NewGetCurrentFarmHandler(projector physicalobject.ObjectProjections) query.Handler {
	return &getCurrentFarmHandler{projector: projector}
}
func (h *getCurrentFarmHandler) Handle(ctx context.Context, payload any) (any, error) {
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
