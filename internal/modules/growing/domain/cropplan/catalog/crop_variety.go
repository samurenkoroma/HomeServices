package catalog

import (
	"errors"
)

// ========== ОШИБКИ ==========

var (
	ErrSpeciesNotFound = errors.New("species not found")
	ErrVarietyNotFound = errors.New("variety not found")
	ErrInvalidGDD      = errors.New("invalid GDD value")
)

// ========== ТИПЫ ДЛЯ КУЛЬТУРЫ (SPECIES) ==========

// Species - вид культуры (обобщенная информация)
// Например: "Томат", "Огурец", "Баклажан"
type Species struct {
	Key         string `json:"key"`      // "tomato"
	Name        string `json:"name"`     // "Томат"
	Family      string `json:"family"`   // "nightshade"
	Category    string `json:"category"` // "Овощные"
	ImageUrl    string `json:"imageUrl"`
	Description string `json:"description"` // описание культуры
}

// ========== ТИПЫ ДЛЯ ФЕНОЛОГИИ (BBCH + GDD) ==========

// PhenophaseGDD - фаза развития с требованиями по GDD
type PhenophaseGDD struct {
	Code        string  `json:"code"`         // "BBCH-10"
	Name        string  `json:"name"`         // "Первый настоящий лист"
	GDDRequired float64 `json:"gdd_required"` // накопленное GDD для достижения
	Description string  `json:"description"`  // описание фазы
	IsCritical  bool    `json:"is_critical"`  // критическая фаза?
}

// ========== ТИПЫ ДЛЯ НОРМ ВЫСЕВА ==========

// SeedingRate - норма высева для одного способа выращивания
type SeedingRate struct {
	GrowingType     string  `json:"growing_type"`     // "open_ground", "greenhouse"
	RowSpacing      float64 `json:"row_spacing"`      // расстояние между рядами (м)
	PlantSpacing    float64 `json:"plant_spacing"`    // расстояние между растениями (м)
	SowingDepth     float64 `json:"sowing_depth"`     // глубина посева (см)
	GerminationRate float64 `json:"germination_rate"` // всхожесть (%)
	SafetyFactor    float64 `json:"safety_factor"`    // страховой коэффициент (1.1-1.3)
}

// CalculateSeedsNeeded рассчитывает количество семян на площадь
func (r SeedingRate) CalculateSeedsNeeded(areaM2 float64) (seeds int, weightKg float64) {
	// Количество растений на 1 м²
	plantsPerM2 := (1 / r.RowSpacing) * (1 / r.PlantSpacing)

	// Количество семян с учетом всхожести и страховки
	seedsNeeded := plantsPerM2 * areaM2 / (r.GerminationRate / 100) * r.SafetyFactor

	// Примерный вес семян (1 семя ≈ 0.005 г)
	weightG := float64(int(seedsNeeded)) * 0.005

	return int(seedsNeeded), weightG / 1000
}

// ========== ГЛАВНЫЙ ТИП: СОРТ (VARIETY) ==========

// Variety - сорт культуры
// Например: "Бычье сердце" (томат), "Алмаз" (баклажан)
type Variety struct {
	// Идентификация
	ID          string `json:"id"`          // "bull_heart"
	Name        string `json:"name"`        // "Бычье сердце"
	SpeciesKey  string `json:"speciesKey"`  // ссылка на вид "tomato"
	SpeciesName string `json:"speciesName"` // денормализовано: "Томат"

	// Температурные параметры для GDD расчета
	BaseTemperature float64 `json:"baseTemperature"` // Tbase (ниже которой рост останавливается)
	MaxTemperature  float64 `json:"maxTemperature"`  // Tmax (выше которой рост не ускоряется)

	// Период вегетации
	DaysToMaturity int `json:"daysToMaturity"` // дней от посадки до сбора

	// Фенология (GDD требования)
	PhenophaseGDD []PhenophaseGDD `json:"phenophaseGDD"`

	// Нормы высева (по способам выращивания)
	SeedingRates map[string]SeedingRate `json:"seedingRates"` // key: "open_ground", "greenhouse"

	// Характеристики
	YieldPotential     float64           `json:"yieldPotential"`     // кг/м²
	PlantHeight        float64           `json:"plantHeight"`        // м
	RecommendedSeasons []string          `json:"recommendedSeasons"` // "spring", "summer"
	GrowingTypes       []string          `json:"growingTypes"`       // "open_ground", "greenhouse"
	Characteristics    map[string]string `json:"characteristics"`
	Description        string            `json:"description"`
	Image              string            `json:"image"`
}

// ========== МЕТОДЫ СОРТА ==========

// GetPhaseByGDD возвращает текущую фазу по накопленному GDD
func (v *Variety) GetPhaseByGDD(accumulatedGDD float64) *PhenophaseGDD {
	if len(v.PhenophaseGDD) == 0 {
		return nil
	}

	// Идем от начала к концу
	for i, phase := range v.PhenophaseGDD {
		if accumulatedGDD < phase.GDDRequired {
			if i == 0 {
				return &v.PhenophaseGDD[0]
			}
			return &v.PhenophaseGDD[i-1]
		}
	}
	// Накоплено больше всех требований → последняя фаза
	return &v.PhenophaseGDD[len(v.PhenophaseGDD)-1]
}

// GetNextPhase возвращает следующую фазу (к которой нужно стремиться)
func (v *Variety) GetNextPhase(currentGDD float64) *PhenophaseGDD {
	for _, phase := range v.PhenophaseGDD {
		if currentGDD < phase.GDDRequired {
			return &phase
		}
	}
	return nil
}

// GetGDDRequiredForPhase возвращает GDD для достижения конкретной фазы
func (v *Variety) GetGDDRequiredForPhase(phaseCode string) (float64, bool) {
	for _, phase := range v.PhenophaseGDD {
		if phase.Code == phaseCode {
			return phase.GDDRequired, true
		}
	}
	return 0, false
}

// GetSeedingRateForGrowingType возвращает норму высева для типа выращивания
func (v *Variety) GetSeedingRateForGrowingType(growingType string) (*SeedingRate, error) {
	rate, ok := v.SeedingRates[growingType]
	if !ok {
		return nil, errors.New("seeding rate not found for " + growingType)
	}
	return &rate, nil
}

// IsSuitableForSeason проверяет, подходит ли сорт для сезона
func (v *Variety) IsSuitableForSeason(seasonName string) bool {
	for _, s := range v.RecommendedSeasons {
		if s == seasonName {
			return true
		}
	}
	return false
}

// IsSuitableForGrowingType проверяет, подходит ли сорт для типа выращивания
func (v *Variety) IsSuitableForGrowingType(growingType string) bool {
	for _, gt := range v.GrowingTypes {
		if gt == growingType {
			return true
		}
	}
	return false
}

// ========== ГЛОБАЛЬНЫЙ КАТАЛОГ ==========

// CropCatalog - глобальный каталог всех видов и сортов
type CropCatalog struct {
	Species   map[string]Species            `json:"species"`   // key: speciesKey
	Varieties map[string]map[string]Variety `json:"varieties"` // key: speciesKey -> varietyID -> Variety
}

// Глобальный экземпляр каталога (заполняется в init() файлов с данными)
var GlobalCatalog = CropCatalog{
	Species:   make(map[string]Species),
	Varieties: make(map[string]map[string]Variety),
}

// ========== ФУНКЦИИ ДЛЯ РАБОТЫ С КАТАЛОГОМ ==========

// GetSpecies возвращает вид по ключу
func GetSpecies(speciesKey string) (*Species, error) {
	species, ok := GlobalCatalog.Species[speciesKey]
	if !ok {
		return nil, ErrSpeciesNotFound
	}
	return &species, nil
}

// GetVariety возвращает сорт по ключу вида и ID сорта
func GetVariety(speciesKey, varietyID string) (*Variety, error) {
	speciesVarieties, ok := GlobalCatalog.Varieties[speciesKey]
	if !ok {
		return nil, ErrSpeciesNotFound
	}

	variety, ok := speciesVarieties[varietyID]
	if !ok {
		return nil, ErrVarietyNotFound
	}
	return &variety, nil
}

// GetVarietiesBySpecies возвращает все сорта вида
func GetVarietiesBySpecies(speciesKey string) ([]Variety, error) {
	speciesVarieties, ok := GlobalCatalog.Varieties[speciesKey]
	if !ok {
		return nil, ErrSpeciesNotFound
	}

	varieties := make([]Variety, 0, len(speciesVarieties))
	for _, v := range speciesVarieties {
		varieties = append(varieties, v)
	}
	return varieties, nil
}

// GetAllVarieties возвращает все сорта из каталога
func GetAllVarieties() []Variety {
	var all []Variety
	for _, speciesVarieties := range GlobalCatalog.Varieties {
		for _, v := range speciesVarieties {
			all = append(all, v)
		}
	}
	return all
}

// RegisterVariety регистрирует новый сорт в каталоге
func RegisterVariety(speciesKey string, variety Variety) error {
	// Проверяем, существует ли вид
	if _, ok := GlobalCatalog.Species[speciesKey]; !ok {
		return ErrSpeciesNotFound
	}

	// Инициализируем map для вида, если еще нет
	if GlobalCatalog.Varieties[speciesKey] == nil {
		GlobalCatalog.Varieties[speciesKey] = make(map[string]Variety)
	}

	GlobalCatalog.Varieties[speciesKey][variety.ID] = variety
	return nil
}

// RegisterSpecies регистрирует новый вид
func RegisterSpecies(species Species) {
	GlobalCatalog.Species[species.Key] = species
}
