package query

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/modules/growing/domain"
)

type GetFacilitiesListQuery struct {
	Limit string `json:"limit"`
}

type GetFacilitiesListHandler struct {
	repo domain.FacilityReadRepository
}

func NewGetFacilitiesListHandler(
	repo domain.FacilityReadRepository,
) *GetFacilitiesListHandler {
	return &GetFacilitiesListHandler{
		repo: repo,
	}
}
func (h *GetFacilitiesListHandler) AsHandler() func(context.Context, any) (any, error) {
	return func(ctx context.Context, payload any) (any, error) {
		q, ok := payload.(*GetFacilitiesListQuery)
		if !ok {
			return nil, errors.New("invalid payload type")
		}
		return h.Handle(ctx, *q)
	}
}

func (h *GetFacilitiesListHandler) Handle(ctx context.Context, q GetFacilitiesListQuery) (any, error) {

	return h.repo.GetList(ctx, domain.FacilitiesListParams{})
}
