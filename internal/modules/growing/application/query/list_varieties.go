package query

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/inmemory"
)

type listVarietiesHandler struct {
	uowFactory repository.Factory
}

func (h *listVarietiesHandler) Name() string {
	return "ListVarieties"
}

func NewListVarietiesHandler(factory repository.Factory) query.Handler {
	return &listVarietiesHandler{
		uowFactory: factory,
	}
}

type ListVarietiesQuery struct {
	SpeciesKey string `json:"speciesKey,omitempty"` // tomato, eggplant, cucumber
}

// VarietyDTO информация о сорте для ответа
type VarietyDTO struct {
	ID                 string            `json:"id"`
	Name               string            `json:"name"`
	SpeciesKey         string            `json:"speciesKey"`
	SpeciesName        string            `json:"speciesName"`
	DaysToMaturity     int               `json:"daysToMaturity"`
	YieldPotential     float64           `json:"yieldPotential"`
	PlantHeight        float64           `json:"plantHeight"`
	RecommendedSeasons []string          `json:"recommendedSeasons"`
	GrowingTypes       []string          `json:"growingTypes"`
	Characteristics    map[string]string `json:"characteristics"`
	Description        string            `json:"description"`
}

// Handle выполняет запрос
func (h *listVarietiesHandler) Handle(ctx context.Context, query any) (any, error) {
	q, ok := query.(*ListVarietiesQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	uow, _ := h.uowFactory.Begin(ctx)
	pr := inmemory.NewRedisGrowingProvider(uow.Tx()).(*inmemory.RedisGrowingProvider)

	// Ищем сорта
	return pr.Catalogs().ListVarieties(ctx, q.SpeciesKey)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// Конвертируем в DTO
	//varietiesDTO := make([]VarietyDTO, len(varieties))
	//for i, v := range varieties {
	//	varietiesDTO[i] = VarietyDTO{
	//		ID:                 v.ID,
	//		Name:               v.Name,
	//		SpeciesKey:         v.SpeciesKey,
	//		SpeciesName:        v.SpeciesName,
	//		DaysToMaturity:     v.DaysToMaturity,
	//		YieldPotential:     v.YieldPotential,
	//		PlantHeight:        v.PlantHeight,
	//		RecommendedSeasons: v.RecommendedSeasons,
	//		GrowingTypes:       v.GrowingTypes,
	//		Characteristics:    v.Characteristics,
	//		Description:        v.Description,
	//	}
	//}
	//
	//return varietiesDTO, nil
}
