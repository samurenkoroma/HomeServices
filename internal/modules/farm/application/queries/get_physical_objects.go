package queries

import (
	"context"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
)

type GetPhysicalObjectsQuery struct {
	Id      string `json:"id,omitempty"`
	TypeObj string `json:"type,omitempty"`
	Status  string `json:"status,omitempty"`
	OwnerId string `json:"owner_id,omitempty"`
	Search  string `json:"search,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Offset  int    `json:"offset,omitempty"`
}
type getPhysicalObjectsHandler struct {
	projector physicalobject.ObjectProjections
}

func (h *getPhysicalObjectsHandler) Name() string {
	return "GetPhysicalObjects"
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
		Limit:   q.Limit,
		Offset:  q.Offset,
	}

	if q.TypeObj != "" {
		filter.Type = physicalobject.ObjectType(q.TypeObj)
	}
	if q.Status != "" {
		filter.Status = valueobject.Status(q.Status)
	}
	if q.Limit == 0 {
		filter.Limit = 10
	}

	return h.projector.GetList(ctx, filter)
}
