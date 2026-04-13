package query

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
)

type getVarietyHandler struct {
	catalog catalog.Repository
}

func (h *getVarietyHandler) Name() string {
	return "GetVariety"
}

func NewGetVarietyHandler(repo catalog.Repository) query.Handler {
	return &getVarietyHandler{
		catalog: repo,
	}
}

type GetVarietyQuery struct {
	Key string `form:"key"`
	Id  string `form:"id"`
}

// Handle выполняет запрос
func (h *getVarietyHandler) Handle(ctx context.Context, query any) (any, error) {
	cmd, ok := query.(*GetVarietyQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	return h.catalog.GetVariety(ctx, cmd.Key, cmd.Id)
}
