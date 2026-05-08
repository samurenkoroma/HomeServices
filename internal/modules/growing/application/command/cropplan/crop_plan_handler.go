package cropplan

import (
	"samurenkoroma/services/internal/core/domain/repository"
)

type PlanHandler struct {
	uowFactory repository.Factory
}

func NewCropPlanHandler(uowFactory repository.Factory) *PlanHandler {
	return &PlanHandler{uowFactory: uowFactory}
}
