package queries

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
)

type ListObjectsOnMapQuery struct {
	Bounds spatial.BoundingBox `json:"bounds"`
}

type MapObjectDTO struct {
	ID       string          `json:"id"`
	Type     string          `json:"type"`
	Name     string          `json:"name"`
	Geometry spatial.GeoJSON `json:"geometry"`
	//Center   [2]float64      `json:"center"`
}

type ListObjectsOnMapHandler struct {
	repo physicalobject.Repository
}

func NewListObjectsOnMapHandler(repo physicalobject.Repository) *ListObjectsOnMapHandler {
	return &ListObjectsOnMapHandler{repo: repo}
}

func (h *ListObjectsOnMapHandler) AsHandler() func(context.Context, any) (any, error) {
	return func(ctx context.Context, payload any) (any, error) {
		q, ok := payload.(*ListObjectsOnMapQuery)
		if !ok {
			return nil, errors.New("invalid payload type")
		}
		return h.Handle(ctx, *q)
	}
}

func (h *ListObjectsOnMapHandler) Handle(ctx context.Context, q ListObjectsOnMapQuery) ([]MapObjectDTO, error) {
	objects, err := h.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	dtos := make([]MapObjectDTO, len(objects))
	for i, obj := range objects {
		dtos[i] = MapObjectDTO{
			ID:       string(obj.Id),
			Type:     string(obj.Type),
			Name:     obj.Name,
			Geometry: obj.Geometry,
			//Center:   spatial.CalculateCenter(obj.Geometry),
		}
	}

	return dtos, nil
}
