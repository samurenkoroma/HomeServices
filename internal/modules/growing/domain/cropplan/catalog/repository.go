package catalog

import (
	"context"
)

// ========== ИНТЕРФЕЙСЫ ==========

// Repository интерфейс для работы с каталогом (может быть реализован для PostgreSQL)
type Repository interface {
	// Species
	GetSpecies(ctx context.Context, key string) (*Species, error)
	ListSpecies(ctx context.Context) ([]Species, error)
	SaveSpecies(ctx context.Context, species *Species) error

	// Varieties
	GetVariety(ctx context.Context, speciesKey, varietyID string) (*Variety, error)
	ListVarieties(ctx context.Context, speciesKey string) ([]Variety, error)
	SearchVarieties(ctx context.Context, filter VarietyFilter) ([]Variety, error)
	SaveVariety(ctx context.Context, speciesKey string, variety *Variety) error
	DeleteVariety(ctx context.Context, speciesKey, varietyID string) error

	// Stage templates
	GetStageTemplates(ctx context.Context, speciesKey string) ([]StageTemplate, error)
	SaveStageTemplate(ctx context.Context, speciesKey string, template *StageTemplate) error
}

// VarietyFilter фильтр для поиска сортов
type VarietyFilter struct {
	SpeciesKey        string `json:"species_key,omitempty"`
	GrowingType       string `json:"growing_type,omitempty"` // "open_ground", "greenhouse"
	Season            string `json:"season,omitempty"`       // "spring", "summer", "autumn"
	MaxDaysToMaturity int    `json:"max_days_to_maturity,omitempty"`
	Query             string `json:"query,omitempty"` // поиск по названию
}
