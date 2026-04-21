package queries

import (
	"context"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
	"time"
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

	fields, _ := h.projector.GetList(ctx, physicalobject.POFilter{Type: physicalobject.ObjectTypeField})
	plots, _ := h.projector.GetList(ctx, physicalobject.POFilter{Type: physicalobject.ObjectTypePlot})
	greenhouse, _ := h.projector.GetList(ctx, physicalobject.POFilter{Type: physicalobject.ObjectTypeGreenhouse})
	farm := Farm{
		Id:        types.NewUUID(),
		Name:      "polevod",
		TotalArea: 12.5,
		Location: struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		}{Lat: 45.9098828805537, Lng: 30.042502443243485},
		Fields:      fields,
		Plots:       plots,
		Greenhouses: greenhouse,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	return farm, nil
}

type Farm struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	TotalArea float64 `json:"totalArea"`
	Location  struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"location"`
	Fields      []*physicalobject.POListItem `json:"fields"`
	Plots       []*physicalobject.POListItem `json:"plots"`
	Greenhouses []*physicalobject.POListItem `json:"greenhouses"`
	CreatedAt   time.Time                    `json:"createdAt"`
	UpdatedAt   time.Time                    `json:"updatedAt"`
}
