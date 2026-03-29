package cropplan

import (
	"encoding/json"
	"samurenkoroma/services/internal/modules/crop/domain/valueobject"
)

// GrowthStage - этап роста культуры
type GrowthStage struct {
	Order           int                  `json:"order"`    // Порядковый номер
	Name            string               `json:"name"`     // Название этапа (вегетация, цветение)
	Duration        valueobject.Duration `json:"duration"` // Длительность в днях
	Description     string               `json:"description"`
	Recommendations Attributes           `json:"recommendations"`
}

// Attributes - типизированные атрибуты для разных объектов
type Attributes struct {
	// Температурные требования
	MinTemp     float64 `json:"min_temp,omitempty"`
	MaxTemp     float64 `json:"max_temp,omitempty"`
	OptimalTemp float64 `json:"optimal_temp,omitempty"`

	// Влажность
	MinHumidity float64 `json:"min_humidity,omitempty"`
	MaxHumidity float64 `json:"max_humidity,omitempty"`

	// Полив
	WaterPerDay float64 `json:"water_per_day,omitempty"` // литров на м² в день
	// Питание
	Nitrogen   float64 `json:"nitrogen,omitempty"` // кг/га
	Phosphorus float64 `json:"phosphorus,omitempty"`
	Potassium  float64 `json:"potassium,omitempty"`
	// Общие
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

func (a *Attributes) ToMap() map[string]interface{} {
	mapped := a.Metadata
	mapped["min_temp"] = a.MinTemp
	mapped["max_temp"] = a.MaxTemp
	mapped["optimal_temp"] = a.OptimalTemp
	mapped["min_humidity"] = a.MinHumidity
	mapped["max_humidity"] = a.MaxHumidity
	mapped["water_per_day"] = a.WaterPerDay
	mapped["nitrogen"] = a.Nitrogen
	mapped["potassium"] = a.Potassium
	mapped["phosphorus"] = a.Phosphorus
	return mapped
}
func (a *Attributes) Marshal() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Attributes) Unmarshal(data []byte) error {
	return json.Unmarshal(data, a)
}

func NewGrowthStage(
	order int,
	name string,
	duration valueobject.Duration,
	minTemp, maxTemp, optimalTemp float64,
	waterPerDay float64,
) (GrowthStage, error) {
	if duration <= 0 {
		return GrowthStage{}, ErrInvalidDuration
	}
	if name == "" {
		return GrowthStage{}, ErrInvalidStageName
	}
	if minTemp > maxTemp {
		return GrowthStage{}, ErrInvalidTemperatureRange
	}

	return GrowthStage{
		Order:    order,
		Name:     name,
		Duration: duration,
		Recommendations: Attributes{
			MinTemp:     minTemp,
			MaxTemp:     maxTemp,
			OptimalTemp: optimalTemp,
			WaterPerDay: waterPerDay,
		},
	}, nil
}
