package queries

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
)

type GetPhysicalObjectsQuery struct {
	Id      string `json:"id"`
	TypeObj string `json:"type"`
	Status  string `json:"status"`
	OwnerId string `json:"owner_id"`
	Search  string `json:"search"`
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
}
type getPhysicalObjectsHandler struct {
	projector physicalobject.POProjection
}

func (h *getPhysicalObjectsHandler) Name() string {
	return "GetPhysicalObjects"
}

func NewGetPhysicalObjectsHandler(projector physicalobject.POProjection) query.QueryHandler {
	return &getPhysicalObjectsHandler{projector: projector}
}
func (h *getPhysicalObjectsHandler) Handle(ctx context.Context, payload any) (any, error) {
	q, ok := payload.(*GetPhysicalObjectsQuery)
	if !ok {
		return nil, errors.New("invalid payload type")
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
