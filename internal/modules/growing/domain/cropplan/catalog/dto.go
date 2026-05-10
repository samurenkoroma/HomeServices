package catalog

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

func (r SeedingRate) CalculateSeedsNeeded(areaM2 float64) (seeds int, weightKg float64) {
	// Количество растений на 1 м²
	plantsPerM2 := (1 / r.RowSpacing) * (1 / r.PlantSpacing)

	// Количество семян с учетом всхожести и страховки
	seedsNeeded := plantsPerM2 * areaM2 / (r.GerminationRate / 100) * r.SafetyFactor

	// Примерный вес семян (1 семя ≈ 0.005 г)
	weightG := float64(int(seedsNeeded)) * 0.005

	return int(seedsNeeded), weightG / 1000
}

type WaterRequirement struct {
	DailyNeedMin   float64  `json:"daily_need_min"`  // л/м² в день (минимально)
	DailyNeedOpt   float64  `json:"daily_need_opt"`  // л/м² в день (оптимально)
	CriticalPhases []string `json:"critical_phases"` // критические BBCH коды
}

// LightRequirement потребность в освещении
type LightRequirement struct {
	PPFDMin         int      `json:"ppfd_min"`         // μmol/m²/s (минимальный фотосинтетический поток)
	PPFDOpt         int      `json:"ppfd_opt"`         // μmol/m²/s (оптимальный)
	DayLengthMin    float64  `json:"day_length_min"`   // часов (минимальный световой день)
	DayLengthOpt    float64  `json:"day_length_opt"`   // часов (оптимальный световой день)
	PhotoperiodType string   `json:"photoperiod_type"` // "short_day", "long_day", "day_neutral"
	CriticalPhases  []string `json:"critical_phases"`  // критические BBCH коды для света
}
