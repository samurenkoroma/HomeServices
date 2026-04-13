package query

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
)

type listSpeciesHandler struct {
	catalog catalog.Repository
}

func (h *listSpeciesHandler) Name() string {
	return "ListSpecies"
}

func NewListSpeciesHandler(repo catalog.Repository) query.Handler {
	return &listSpeciesHandler{
		catalog: repo,
	}
}

type ListSpeciesQuery struct {
}

type SpeciesDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ListSpeciesResponse struct {
	Total   int          `json:"total"`
	Species []SpeciesDTO `json:"varieties"`
}

// Handle выполняет запрос
func (h *listSpeciesHandler) Handle(ctx context.Context, query any) (any, error) {
	_, ok := query.(*ListSpeciesQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	return h.catalog.ListSpecies(ctx)
}
