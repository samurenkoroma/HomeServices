package query

import (
	"context"
	"errors"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/inmemory"
)

// SearchVarietiesHandler запрос поиска сортов
type searchVarietiesHandler struct {
	uowFactory repository.Factory
}

func (h *searchVarietiesHandler) Name() string {
	return "SearchVarieties"
}

func NewSearchVarietiesHandler(factory repository.Factory) query.Handler {
	return &searchVarietiesHandler{
		uowFactory: factory,
	}
}

// SearchVarietiesQuery параметры поиска
type SearchVarietiesQuery struct {
	SpeciesKey        string `json:"speciesKey,omitempty"`   // tomato, eggplant, cucumber
	GrowingType       string `json:"growing_type,omitempty"` // open_ground, greenhouse
	Season            string `json:"season,omitempty"`       // spring, summer, autumn
	Query             string `json:"query,omitempty"`        // поиск по названию
	MaxDaysToMaturity int    `json:"max_days_to_maturity,omitempty"`
	Limit             int    `json:"limit,omitempty"`
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
func (h *searchVarietiesHandler) Handle(ctx context.Context, query any) (any, error) {
	q, ok := query.(*SearchVarietiesQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	uow, err := h.uowFactory.Begin(ctx)
	pr := inmemory.NewGrowingProvider(uow.Tx()).(*inmemory.GrowingProvider)

	// Строим фильтр
	filter := catalog.VarietyFilter{
		SpeciesKey:        q.SpeciesKey,
		GrowingType:       q.GrowingType,
		Season:            q.Season,
		Query:             q.Query,
		MaxDaysToMaturity: q.MaxDaysToMaturity,
	}

	// Ищем сорта
	varieties, err := pr.Catalogs().SearchVarieties(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Применяем лимит
	limit := q.Limit
	if limit <= 0 || limit > len(varieties) {
		limit = len(varieties)
	}

	result := varieties[:limit]

	// Конвертируем в DTO
	varietiesDTO := make([]VarietyDTO, len(result))
	for i, v := range result {
		varietiesDTO[i] = VarietyDTO{
			ID:                 v.ID,
			Name:               v.Name,
			SpeciesKey:         v.SpeciesKey,
			SpeciesName:        v.SpeciesName,
			DaysToMaturity:     v.DaysToMaturity,
			YieldPotential:     v.YieldPotential,
			PlantHeight:        v.PlantHeight,
			RecommendedSeasons: v.RecommendedSeasons,
			GrowingTypes:       v.GrowingTypes,
			Characteristics:    v.Characteristics,
			Description:        v.Description,
		}
	}

	return varietiesDTO, nil
}
