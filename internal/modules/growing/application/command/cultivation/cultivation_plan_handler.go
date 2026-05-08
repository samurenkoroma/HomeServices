package cultivation

import (
	"samurenkoroma/services/internal/core/domain/repository"
)

type PlanHandler struct {
	uowFactory repository.Factory
}

func NewCultivationPlanHandler(uowFactory repository.Factory) *PlanHandler {
	return &PlanHandler{uowFactory: uowFactory}
}
