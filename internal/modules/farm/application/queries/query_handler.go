package queries

import (
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
)

type QueryHandler struct {
	projector physicalobject.ObjectProjections
}

func NewFarmHandler(projector physicalobject.ObjectProjections) *QueryHandler {
	return &QueryHandler{projector: projector}
}
