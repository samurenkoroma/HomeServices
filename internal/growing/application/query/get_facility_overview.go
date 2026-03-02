package query

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/growing/application"
)

type GetFacilityOverviewQuery struct {
	FacilityID string `json:"facility_id"`
}

type GetFacilityOverviewHandler struct {
	repo application.FacilityReadRepository
}

func NewGetFacilityOverviewHandler(
	repo application.FacilityReadRepository,
) *GetFacilityOverviewHandler {
	return &GetFacilityOverviewHandler{
		repo: repo,
	}
}
func (h *GetFacilityOverviewHandler) AsHandler() func(context.Context, any) (any, error) {
	return func(ctx context.Context, payload any) (any, error) {
		q, ok := payload.(*GetFacilityOverviewQuery)
		if !ok {
			return nil, errors.New("invalid payload type")
		}
		return h.Handle(ctx, *q)
	}
}

func (h *GetFacilityOverviewHandler) Handle(
	ctx context.Context,
	q GetFacilityOverviewQuery,
) (any, error) {

	return h.repo.GetOverview(ctx, q.FacilityID)
}
