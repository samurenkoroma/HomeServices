package catalog

import (
	"errors"
)

// ========== ОШИБКИ ==========

var (
	ErrCropNotFound    = errors.New("crop not found")
	ErrVarietyNotFound = errors.New("variety not found")
	ErrInvalidGDD      = errors.New("invalid GDD value")
)

// ========== ТИПЫ ДЛЯ КУЛЬТУРЫ (SPECIES) ==========

// Crop - вид культуры (обобщенная информация)
// Например: "Томат", "Огурец", "Баклажан"
type Crop struct {
	Key         string // "tomato"
	Name        string // "Томат"
	Family      string // "nightshade"
	Category    string // "Овощные"
	ImageUrl    string
	Description string // описание культуры
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
	// Водные требования
	WaterRequirement WaterRequirement `json:"water_requirement"`
	// Световые требования
	LightRequirement LightRequirement `json:"light_requirement"`
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

	phaseIndex map[string]int // не сериализуется
}

func (v *Variety) initIndex() {
	if v.phaseIndex != nil {
		return
	}

	v.phaseIndex = make(map[string]int)
	for i, p := range v.PhenophaseGDD {
		v.phaseIndex[p.Code] = i
	}
}
func (v *Variety) GetId() string {
	return v.ID
}

func (v *Variety) GetName() string {
	return v.Name
}

func (v *Variety) GetSpeciesName() string {
	return v.SpeciesName
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
	for i := range v.PhenophaseGDD {
		if currentGDD < v.PhenophaseGDD[i].GDDRequired {
			return &v.PhenophaseGDD[i]
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
func (v *Variety) GetPhaseByCode(code string) *PhenophaseGDD {
	v.initIndex()

	if i, ok := v.phaseIndex[code]; ok {
		return &v.PhenophaseGDD[i]
	}
	return nil
}
