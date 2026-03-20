package cropplan

// GrowthStage - этап роста культуры
type GrowthStage struct {
	Order       int    `json:"order"`    // Порядковый номер
	Name        string `json:"name"`     // Название этапа (вегетация, цветение)
	Duration    int    `json:"duration"` // Длительность в днях
	Description string `json:"description"`

	// Температурные требования
	MinTemp     float64 `json:"min_temp"`
	MaxTemp     float64 `json:"max_temp"`
	OptimalTemp float64 `json:"optimal_temp"`

	// Влажность
	MinHumidity float64 `json:"min_humidity"`
	MaxHumidity float64 `json:"max_humidity"`

	// Полив
	WaterPerDay float64 `json:"water_per_day"` // литров на м² в день

	// Питание
	NitrogenReq   float64 `json:"nitrogen_req"` // кг/га
	PhosphorusReq float64 `json:"phosphorus_req"`
	PotassiumReq  float64 `json:"potassium_req"`
}

func NewGrowthStage(
	order int,
	name string,
	duration int,
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
		Order:       order,
		Name:        name,
		Duration:    duration,
		MinTemp:     minTemp,
		MaxTemp:     maxTemp,
		OptimalTemp: optimalTemp,
		WaterPerDay: waterPerDay,
	}, nil
}
