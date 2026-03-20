package cropplan

// EnvironmentalRequirements - требования к окружающей среде
type EnvironmentalRequirements struct {
	// Температура
	MinTemp float64 `json:"min_temp"`
	MaxTemp float64 `json:"max_temp"`

	// Влажность
	MinHumidity float64 `json:"min_humidity"`
	MaxHumidity float64 `json:"max_humidity"`

	// Почва
	MinPH     float64  `json:"min_ph"`
	MaxPH     float64  `json:"max_ph"`
	SoilTypes []string `json:"soil_types"` // Подходящие типы почв

	// Освещение
	MinLightHours int `json:"min_light_hours"`

	// Вода
	TotalWaterRequirement float64 `json:"total_water_requirement"` // м³/га за сезон
}

// NutrientRequirements - требования к питательным веществам
type NutrientRequirements struct {
	Nitrogen   float64 `json:"nitrogen"`   // кг/га
	Phosphorus float64 `json:"phosphorus"` // кг/га
	Potassium  float64 `json:"potassium"`  // кг/га
	Calcium    float64 `json:"calcium"`
	Magnesium  float64 `json:"magnesium"`
	Sulfur     float64 `json:"sulfur"`
}

// Validate проверяет корректность требований
func (e EnvironmentalRequirements) Validate() error {
	if e.MinTemp > e.MaxTemp {
		return ErrInvalidTemperatureRange
	}
	if e.MinPH > e.MaxPH {
		return ErrInvalidPHRange
	}
	if e.MinHumidity > e.MaxHumidity {
		return ErrInvalidHumidityRange
	}
	return nil
}
