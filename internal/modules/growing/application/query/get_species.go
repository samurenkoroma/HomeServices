package query

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
)

type getSpeciesHandler struct {
	catalog catalog.Repository
}

func (h *getSpeciesHandler) Name() string {
	return "GetSpecies"
}

func NewGetSpeciesHandler(repo catalog.Repository) query.Handler {
	return &getSpeciesHandler{
		catalog: repo,
	}
}

type GetSpeciesQuery struct {
	Key string `form:"key"`
}

// Handle выполняет запрос
func (h *getSpeciesHandler) Handle(ctx context.Context, query any) (any, error) {
	cmd, ok := query.(*GetSpeciesQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	return h.catalog.GetSpecies(ctx, cmd.Key)
}
