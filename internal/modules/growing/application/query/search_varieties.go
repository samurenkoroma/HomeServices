package query

import (
	"context"
	"encoding/json"
	"errors"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
)

// SearchVarietiesHandler запрос поиска сортов
type SearchVarietiesHandler struct {
	CatalogRepo catalog.Repository
}

// SearchVarietiesQuery параметры поиска
type SearchVarietiesQuery struct {
	SpeciesKey        string `json:"species_key,omitempty"`  // tomato, eggplant, cucumber
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
	SpeciesKey         string            `json:"species_key"`
	SpeciesName        string            `json:"species_name"`
	DaysToMaturity     int               `json:"days_to_maturity"`
	YieldPotential     float64           `json:"yield_potential"`
	PlantHeight        float64           `json:"plant_height"`
	RecommendedSeasons []string          `json:"recommended_seasons"`
	GrowingTypes       []string          `json:"growing_types"`
	Characteristics    map[string]string `json:"characteristics"`
	Description        string            `json:"description"`
}

// SearchVarietiesResponse ответ с результатами поиска
type SearchVarietiesResponse struct {
	Total     int          `json:"total"`
	Varieties []VarietyDTO `json:"varieties"`
}

// DecodeSearchVarieties декодирует JSON в запрос
func DecodeSearchVarieties(data []byte) (any, error) {
	var q SearchVarietiesQuery
	if err := json.Unmarshal(data, &q); err != nil {
		return nil, err
	}
	return q, nil
}

// Handle выполняет запрос
func (h *SearchVarietiesHandler) Handle(ctx context.Context, query any) (any, error) {
	q, ok := query.(SearchVarietiesQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	// Строим фильтр
	filter := catalog.VarietyFilter{
		SpeciesKey:        q.SpeciesKey,
		GrowingType:       q.GrowingType,
		Season:            q.Season,
		Query:             q.Query,
		MaxDaysToMaturity: q.MaxDaysToMaturity,
	}

	// Ищем сорта
	varieties, err := h.CatalogRepo.SearchVarieties(ctx, filter)
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

	return &SearchVarietiesResponse{
		Total:     len(varieties),
		Varieties: varietiesDTO,
	}, nil
}
