package cropplan

import (
	"samurenkoroma/services/internal/core/domain/repository"
)

type CropPlanHandler struct {
	uowFactory repository.Factory
}

func NewCropPlanHandler(uowFactory repository.Factory) *CropPlanHandler {
	return &CropPlanHandler{uowFactory: uowFactory}
}
