package catalog

import (
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
)

type QueryHandler struct {
	catalog catalog.CatalogProjections
}

func NewCatalogHandler(catalog catalog.CatalogProjections) *QueryHandler {
	return &QueryHandler{
		catalog: catalog,
	}
}
