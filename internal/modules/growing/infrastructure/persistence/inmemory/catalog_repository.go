package inmemory

import (
	"context"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"strings"
	"sync"
)

// InMemoryCatalogRepository реализация репозитория каталога в памяти
type InMemoryCatalogRepository struct {
	mu        sync.RWMutex
	species   map[string]catalog.Species
	varieties map[string]map[string]catalog.Variety
}

// NewInMemoryCatalogRepository создает новый in-memory репозиторий
func NewInMemoryCatalogRepository() *InMemoryCatalogRepository {
	repo := &InMemoryCatalogRepository{
		species:   make(map[string]catalog.Species),
		varieties: make(map[string]map[string]catalog.Variety),
	}

	// Инициализируем данными из глобального каталога
	for key, species := range catalog.GlobalCatalog.Species {
		repo.species[key] = species
	}

	for speciesKey, vars := range catalog.GlobalCatalog.Varieties {
		if repo.varieties[speciesKey] == nil {
			repo.varieties[speciesKey] = make(map[string]catalog.Variety)
		}
		for id, v := range vars {
			repo.varieties[speciesKey][id] = v
		}
	}

	return repo
}

// GetSpecies возвращает вид по ключу
func (r *InMemoryCatalogRepository) GetSpecies(ctx context.Context, key string) (*catalog.Species, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	species, ok := r.species[key]
	if !ok {
		return nil, catalog.ErrSpeciesNotFound
	}
	return &species, nil
}

// ListSpecies возвращает все виды
func (r *InMemoryCatalogRepository) ListSpecies(ctx context.Context) ([]catalog.Species, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]catalog.Species, 0, len(r.species))
	for _, s := range r.species {
		result = append(result, s)
	}
	return result, nil
}

// ListSpeciesByCategory возвращает виды по категории
func (r *InMemoryCatalogRepository) ListSpeciesByCategory(ctx context.Context, category string) ([]catalog.Species, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []catalog.Species
	for _, s := range r.species {
		if s.Category == category {
			result = append(result, s)
		}
	}
	return result, nil
}

// SaveSpecies сохраняет вид
func (r *InMemoryCatalogRepository) SaveSpecies(ctx context.Context, species *catalog.Species) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.species[species.Key] = *species
	return nil
}

// DeleteSpecies удаляет вид
func (r *InMemoryCatalogRepository) DeleteSpecies(ctx context.Context, key string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.species[key]; !ok {
		return catalog.ErrSpeciesNotFound
	}
	delete(r.species, key)
	return nil
}

// GetVariety возвращает сорт
func (r *InMemoryCatalogRepository) GetVariety(ctx context.Context, speciesKey, varietyID string) (*catalog.Variety, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Если speciesKey не указан, ищем во всех
	if speciesKey == "" {
		for _, speciesVarieties := range r.varieties {
			if v, ok := speciesVarieties[varietyID]; ok {
				return &v, nil
			}
		}
		return nil, catalog.ErrVarietyNotFound
	}

	speciesVarieties, ok := r.varieties[speciesKey]
	if !ok {
		return nil, catalog.ErrSpeciesNotFound
	}

	variety, ok := speciesVarieties[varietyID]
	if !ok {
		return nil, catalog.ErrVarietyNotFound
	}
	return &variety, nil
}

// ListVarieties возвращает все сорта вида
func (r *InMemoryCatalogRepository) ListVarieties(ctx context.Context, speciesKey string) ([]catalog.Variety, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	speciesVarieties, ok := r.varieties[speciesKey]
	if !ok {
		return nil, catalog.ErrSpeciesNotFound
	}

	result := make([]catalog.Variety, 0, len(speciesVarieties))
	for _, v := range speciesVarieties {
		result = append(result, v)
	}
	return result, nil
}

// SearchVarieties ищет сорта по фильтру
func (r *InMemoryCatalogRepository) SearchVarieties(ctx context.Context, filter catalog.VarietyFilter) ([]catalog.Variety, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []catalog.Variety

	for speciesKey, speciesVarieties := range r.varieties {
		if filter.SpeciesKey != "" && filter.SpeciesKey != speciesKey {
			continue
		}

		for _, variety := range speciesVarieties {
			if filter.GrowingType != "" && !contains(variety.GrowingTypes, filter.GrowingType) {
				continue
			}
			if filter.Season != "" && !contains(variety.RecommendedSeasons, filter.Season) {
				continue
			}
			if filter.MaxDaysToMaturity > 0 && variety.DaysToMaturity > filter.MaxDaysToMaturity {
				continue
			}
			if filter.Query != "" && !strings.Contains(strings.ToLower(variety.Name), strings.ToLower(filter.Query)) {
				continue
			}
			result = append(result, variety)
		}
	}

	return result, nil
}

// SaveVariety сохраняет сорт
func (r *InMemoryCatalogRepository) SaveVariety(ctx context.Context, speciesKey string, variety *catalog.Variety) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.varieties[speciesKey] == nil {
		r.varieties[speciesKey] = make(map[string]catalog.Variety)
	}

	variety.SpeciesKey = speciesKey
	r.varieties[speciesKey][variety.ID] = *variety
	return nil
}

// DeleteVariety удаляет сорт
func (r *InMemoryCatalogRepository) DeleteVariety(ctx context.Context, speciesKey, varietyID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	speciesVarieties, ok := r.varieties[speciesKey]
	if !ok {
		return catalog.ErrSpeciesNotFound
	}

	if _, ok := speciesVarieties[varietyID]; !ok {
		return catalog.ErrVarietyNotFound
	}

	delete(r.varieties[speciesKey], varietyID)
	return nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
