package cropplan

import (
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"
)

type QueryHandler struct {
	provider *postgres.PostgresGrowingProvider
}

func NewCropPlanQueryHandler(provider *postgres.PostgresGrowingProvider) *QueryHandler {
	return &QueryHandler{provider: provider}
}
