package seasons

import (
	"samurenkoroma/services/internal/modules/growing/domain/season"
)

type QueryHandler struct {
	seasons season.Repository
}

func NewSeasonsHandler(seasons season.Repository) *QueryHandler {
	return &QueryHandler{
		seasons: seasons,
	}
}
