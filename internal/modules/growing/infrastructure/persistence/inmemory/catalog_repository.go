package inmemory

import (
	"context"
	"encoding/json"
	"errors"
	catalog2 "samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"sync"
)

// ========== IN-MEMORY РЕАЛИЗАЦИЯ (для тестов и разработки) ==========

// InMemoryCatalogRepository реализация каталога в памяти
type InMemoryCatalogRepository struct {
	mu        sync.RWMutex
	species   map[string]catalog2.Species
	varieties map[string]map[string]catalog2.Variety
	templates map[string][]catalog2.StageTemplate
}

// NewInMemoryCatalogRepository создает новый in-memory репозиторий
// и инициализирует его данными из GlobalCatalog
func NewInMemoryCatalogRepository() *InMemoryCatalogRepository {
	repo := &InMemoryCatalogRepository{
		species:   make(map[string]catalog2.Species),
		varieties: make(map[string]map[string]catalog2.Variety),
		templates: make(map[string][]catalog2.StageTemplate),
	}

	// Загружаем данные из глобального каталога
	for key, species := range catalog2.GlobalCatalog.Species {
		repo.species[key] = species
	}

	for speciesKey, vars := range catalog2.GlobalCatalog.Varieties {
		if repo.varieties[speciesKey] == nil {
			repo.varieties[speciesKey] = make(map[string]catalog2.Variety)
		}
		for id, v := range vars {
			repo.varieties[speciesKey][id] = v
		}
	}

	for speciesKey, templates := range catalog2.StageTemplatesBySpecies {
		repo.templates[speciesKey] = templates
	}

	return repo
}

// GetSpecies возвращает вид по ключу
func (r *InMemoryCatalogRepository) GetSpecies(ctx context.Context, key string) (*catalog2.Species, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	species, ok := r.species[key]
	if !ok {
		return nil, catalog2.ErrSpeciesNotFound
	}
	return &species, nil
}

// ListSpecies возвращает все виды
func (r *InMemoryCatalogRepository) ListSpecies(ctx context.Context) ([]catalog2.Species, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]catalog2.Species, 0, len(r.species))
	for _, s := range r.species {
		result = append(result, s)
	}
	return result, nil
}

// SaveSpecies сохраняет вид
func (r *InMemoryCatalogRepository) SaveSpecies(ctx context.Context, species *catalog2.Species) error {
	if species.Key == "" {
		return errors.New("species key is required")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.species[species.Key] = *species
	return nil
}

// GetVariety возвращает сорт
func (r *InMemoryCatalogRepository) GetVariety(ctx context.Context, speciesKey, varietyID string) (*catalog2.Variety, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	speciesVars, ok := r.varieties[speciesKey]
	if !ok {
		return nil, catalog2.ErrSpeciesNotFound
	}

	variety, ok := speciesVars[varietyID]
	if !ok {
		return nil, catalog2.ErrVarietyNotFound
	}
	return &variety, nil
}

// ListVarieties возвращает все сорта вида
func (r *InMemoryCatalogRepository) ListVarieties(ctx context.Context, speciesKey string) ([]catalog2.Variety, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	speciesVars, ok := r.varieties[speciesKey]
	if !ok {
		return nil, catalog2.ErrSpeciesNotFound
	}

	result := make([]catalog2.Variety, 0, len(speciesVars))
	for _, v := range speciesVars {
		result = append(result, v)
	}
	return result, nil
}

// SearchVarieties ищет сорта по фильтру
func (r *InMemoryCatalogRepository) SearchVarieties(ctx context.Context, filter catalog2.VarietyFilter) ([]catalog2.Variety, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []catalog2.Variety

	for speciesKey, speciesVars := range r.varieties {
		// Фильтр по виду
		if filter.SpeciesKey != "" && filter.SpeciesKey != speciesKey {
			continue
		}

		for _, variety := range speciesVars {
			// Фильтр по типу выращивания
			if filter.GrowingType != "" {
				if !variety.IsSuitableForGrowingType(filter.GrowingType) {
					continue
				}
			}

			// Фильтр по сезону
			if filter.Season != "" {
				if !variety.IsSuitableForSeason(filter.Season) {
					continue
				}
			}

			// Фильтр по сроку созревания
			if filter.MaxDaysToMaturity > 0 {
				if variety.DaysToMaturity > filter.MaxDaysToMaturity {
					continue
				}
			}

			// Поиск по названию
			if filter.Query != "" {
				// простая проверка вхождения (можно заменить на contains)
				if !contains(variety.Name, filter.Query) && !contains(variety.ID, filter.Query) {
					continue
				}
			}

			result = append(result, variety)
		}
	}

	return result, nil
}

// SaveVariety сохраняет сорт
func (r *InMemoryCatalogRepository) SaveVariety(ctx context.Context, speciesKey string, variety *catalog2.Variety) error {
	if speciesKey == "" {
		return errors.New("species key is required")
	}
	if variety.ID == "" {
		return errors.New("variety ID is required")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.varieties[speciesKey] == nil {
		r.varieties[speciesKey] = make(map[string]catalog2.Variety)
	}

	variety.SpeciesKey = speciesKey
	r.varieties[speciesKey][variety.ID] = *variety
	return nil
}

// DeleteVariety удаляет сорт
func (r *InMemoryCatalogRepository) DeleteVariety(ctx context.Context, speciesKey, varietyID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	speciesVars, ok := r.varieties[speciesKey]
	if !ok {
		return catalog2.ErrSpeciesNotFound
	}

	if _, ok := speciesVars[varietyID]; !ok {
		return catalog2.ErrVarietyNotFound
	}

	delete(r.varieties[speciesKey], varietyID)
	return nil
}

// GetStageTemplates возвращает шаблоны этапов для вида
func (r *InMemoryCatalogRepository) GetStageTemplates(ctx context.Context, speciesKey string) ([]catalog2.StageTemplate, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	templates, ok := r.templates[speciesKey]
	if !ok {
		return []catalog2.StageTemplate{}, nil
	}
	return templates, nil
}

// SaveStageTemplate сохраняет шаблон этапа
func (r *InMemoryCatalogRepository) SaveStageTemplate(ctx context.Context, speciesKey string, template *catalog2.StageTemplate) error {
	if speciesKey == "" {
		return errors.New("species key is required")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	// Проверяем, есть ли уже шаблоны для этого вида
	existing := r.templates[speciesKey]

	// Обновляем или добавляем
	found := false
	for i, t := range existing {
		if t.Name == template.Name {
			existing[i] = *template
			found = true
			break
		}
	}

	if !found {
		r.templates[speciesKey] = append(r.templates[speciesKey], *template)
	} else {
		r.templates[speciesKey] = existing
	}

	return nil
}

// ========== ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ ==========

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0)
}

// ========== JSON СЕРИАЛИЗАЦИЯ (для экспорта/импорта) ==========

// Export экспортирует каталог в JSON
func (r *InMemoryCatalogRepository) Export() ([]byte, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	data := struct {
		Species   map[string]catalog2.Species            `json:"species"`
		Varieties map[string]map[string]catalog2.Variety `json:"varieties"`
		Templates map[string][]catalog2.StageTemplate    `json:"templates"`
	}{
		Species:   r.species,
		Varieties: r.varieties,
		Templates: r.templates,
	}

	return json.MarshalIndent(data, "", "  ")
}

// Import импортирует каталог из JSON
func (r *InMemoryCatalogRepository) Import(data []byte) error {
	var input struct {
		Species   map[string]catalog2.Species            `json:"species"`
		Varieties map[string]map[string]catalog2.Variety `json:"varieties"`
		Templates map[string][]catalog2.StageTemplate    `json:"templates"`
	}

	if err := json.Unmarshal(data, &input); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.species = input.Species
	r.varieties = input.Varieties
	r.templates = input.Templates

	return nil
}
