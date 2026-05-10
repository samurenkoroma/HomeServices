package cropplan

import (
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
	"samurenkoroma/services/internal/modules/growing/infrastructure/projections"
)

type QueryHandler struct {
	projector *projections.GrowingProjectionsProvider
	provider  *postgres.PostgresGrowingProvider
}

func NewCropPlanQueryHandler(projector *projections.GrowingProjectionsProvider, provider *postgres.PostgresGrowingProvider) *QueryHandler {
	return &QueryHandler{projector: projector, provider: provider}
}
