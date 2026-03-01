package query

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type GetFacilityOverviewQuery struct {
	FacilityID string `json:"facility_id"`
}

type FacilityOverviewDTO struct {
	ID    string        `json:"id"`
	Name  string        `json:"name"`
	Type  string        `json:"type"` // field | greenhouse
	Area  float64       `json:"area"`
	Units []LandUnitDTO `json:"units"`
}

type LandUnitDTO struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Area float64 `json:"area"`
}
type FacilityReadRepository interface {
	GetOverview(ctx context.Context, id string) (*FacilityOverviewDTO, error)
}

type GetFacilityOverviewHandler struct {
}

func NewGetFacilityOverviewHandler() *GetFacilityOverviewHandler {
	return &GetFacilityOverviewHandler{}
}

func (h *GetFacilityOverviewHandler) AsHandler() func(context.Context, any) (any, error) {
	return func(ctx context.Context, payload any) (any, error) {
		q, ok := payload.(*GetFacilityOverviewQuery)
		if !ok {
			return nil, errors.New("invalid payload type")
		}
		return h.Handle(ctx, q)
	}
}

func (h *GetFacilityOverviewHandler) Handle(
	ctx context.Context,
	q *GetFacilityOverviewQuery,
) (*FacilityOverviewDTO, error) {

	if q.FacilityID == "" {
		return nil, errors.New("facility_id is required")
	}

	return &FacilityOverviewDTO{
		ID:    uuid.New().String(),
		Name:  "test",
		Type:  "test",
		Area:  0.6,
		Units: nil,
	}, nil
}
