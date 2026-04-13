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

	// GetStageTemplates возвращает шаблоны этапов для вида
	GetStageTemplates(ctx context.Context, speciesKey string) ([]StageTemplate, error)

	// GetStageTemplatesByBBCH возвращает шаблоны этапов для вида в определенном BBCH диапазоне
	GetStageTemplatesByBBCH(ctx context.Context, speciesKey string, bbchCode int) ([]StageTemplate, error)

	// SaveStageTemplate сохраняет шаблон этапа
	SaveStageTemplate(ctx context.Context, speciesKey string, template *StageTemplate) error

	// SaveStageTemplatesBatch массово сохраняет шаблоны этапов
	SaveStageTemplatesBatch(ctx context.Context, speciesKey string, templates []StageTemplate) error

	// DeleteStageTemplate удаляет шаблон этапа
	DeleteStageTemplate(ctx context.Context, speciesKey string, templateType string, displayOrder int) error

	// DeleteAllStageTemplates удаляет все шаблоны этапов для вида
	DeleteAllStageTemplates(ctx context.Context, speciesKey string) error
}

// VarietyFilter фильтр для поиска сортов
type VarietyFilter struct {
	SpeciesKey        string `json:"species_key,omitempty"`
	GrowingType       string `json:"growing_type,omitempty"` // "open_ground", "greenhouse"
	Season            string `json:"season,omitempty"`       // "spring", "summer", "autumn"
	MaxDaysToMaturity int    `json:"max_days_to_maturity,omitempty"`
	Query             string `json:"query,omitempty"` // поиск по названию
}
