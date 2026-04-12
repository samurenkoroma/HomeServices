package phenology

import (
	"time"
)

// DailyTemp температура за один день
type DailyTemp struct {
	Date time.Time `json:"date"`
	Min  float64   `json:"min"` // минимальная температура (°C)
	Max  float64   `json:"max"` // максимальная температура (°C)
}

// CurrentPhenology текущее фенологическое состояние
type CurrentPhenology struct {
	PlanID      string `json:"plan_id"`
	VarietyID   string `json:"variety_id"`
	VarietyName string `json:"variety_name"`

	// GDD
	AccumulatedGDD     float64 `json:"accumulated_gdd"`
	RequiredGDDForNext float64 `json:"required_gdd_for_next"`

	// Текущая фаза
	CurrentPhaseCode string  `json:"current_phase_code"`
	CurrentPhaseName string  `json:"current_phase_name"`
	ProgressPercent  float64 `json:"progress_percent"`

	// Прогноз
	EstimatedDaysToNextPhase int        `json:"estimated_days_to_next_phase"`
	EstimatedHarvestDate     *time.Time `json:"estimated_harvest_date,omitempty"`

	// Отклонения
	DeviationDays   int    `json:"deviation_days"`   // +/- дней от нормы
	DeviationReason string `json:"deviation_reason"` // "heat_wave", "cold_spell", "drought"

	// Критичность
	IsCritical bool `json:"is_critical"`

	// Рекомендации
	RecommendedActions []RecommendedAction `json:"recommended_actions"`
}

// RecommendedAction рекомендация по действию
type RecommendedAction struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"` // "low", "medium", "high", "urgent"
	DueDays     int    `json:"due_days"` // через сколько дней нужно выполнить
}

// PhenologyForecast прогноз развития
type PhenologyForecast struct {
	PlanID             string              `json:"plan_id"`
	PlantingDate       time.Time           `json:"planting_date"`
	ForecastDate       time.Time           `json:"forecast_date"`
	Phases             []ForecastPhase     `json:"phases"`
	RecommendedActions []RecommendedAction `json:"recommended_actions"`
}

// ForecastPhase прогнозируемая фаза
type ForecastPhase struct {
	PhaseCode    string    `json:"phase_code"`
	PhaseName    string    `json:"phase_name"`
	ExpectedDate time.Time `json:"expected_date"`
	GDDRequired  float64   `json:"gdd_required"`
	IsCritical   bool      `json:"is_critical"`
}
