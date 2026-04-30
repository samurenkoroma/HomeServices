package season

import (
	"samurenkoroma/services/internal/core/domain/repository"
)

type SeasonHandler struct {
	uowFactory repository.Factory
}

func NewSeasonHandler(uowFactory repository.Factory) *SeasonHandler {
	return &SeasonHandler{uowFactory: uowFactory}
}
