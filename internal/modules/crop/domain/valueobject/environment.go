package valueobject

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
