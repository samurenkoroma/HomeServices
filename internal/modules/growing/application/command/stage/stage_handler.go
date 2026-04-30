package stage

import (
	"samurenkoroma/services/internal/core/domain/repository"
)

type StageHandler struct {
	uowFactory repository.Factory
}

func NewStageHandler(uowFactory repository.Factory) *StageHandler {
	return &StageHandler{
		uowFactory: uowFactory,
	}
}
