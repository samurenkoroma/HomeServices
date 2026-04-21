package queries

import (
	"context"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
)

type GetPhysicalObjectsQuery struct {
	Id      string `json:"id,omitempty"`
	FarmId  string `json:"farmId,omitempty"`
	TypeObj string `json:"type,omitempty"`
	Status  string `json:"status,omitempty"`
	OwnerId string `json:"owner_id,omitempty"`
	Search  string `json:"search,omitempty"`
}
type getPhysicalObjectsHandler struct {
	projector physicalobject.ObjectProjections
}

func NewGetPhysicalObjectsHandler(projector physicalobject.ObjectProjections) query.Handler {
	return &getPhysicalObjectsHandler{projector: projector}
}
func (h *getPhysicalObjectsHandler) Handle(ctx context.Context, payload any) (any, error) {
	q, ok := payload.(*GetPhysicalObjectsQuery)
	if !ok {
		return nil, query.ErrInvalidPayloadType
	}

	if q.Id != "" {
		return h.projector.GetByID(ctx, q.Id)
	}

	filter := physicalobject.POFilter{
		OwnerId: q.OwnerId,
		Search:  q.Search,
	}

	if q.TypeObj != "" {
		filter.Type = physicalobject.ObjectType(q.TypeObj)
	}
	if q.Status != "" {
		filter.Status = valueobject.Status(q.Status)
	} else {
		filter.Status = valueobject.Active
	}

	return h.projector.GetList(ctx, filter)
}
