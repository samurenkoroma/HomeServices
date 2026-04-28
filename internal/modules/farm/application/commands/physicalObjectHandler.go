package commands

import (
	"samurenkoroma/services/internal/core/domain/repository"
)

type FarmObjectHandler struct {
	uowFactory repository.Factory
}

func NewFarmObjectHandler(uowFactory repository.Factory) *FarmObjectHandler {
	return &FarmObjectHandler{uowFactory: uowFactory}
}
