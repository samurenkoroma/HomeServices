package catalog

import (
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
)

type QueryHandler struct {
	catalog catalog.Repository
}

func NewCatalogHandler(catalog catalog.Repository) *QueryHandler {
	return &QueryHandler{
		catalog: catalog,
	}
}
