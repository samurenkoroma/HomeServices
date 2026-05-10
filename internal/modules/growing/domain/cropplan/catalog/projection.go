package catalog

import "context"

type CropDto struct {
	Key         string `json:"key"`      // "tomato"
	Name        string `json:"name"`     // "Томат"
	Family      string `json:"family"`   // "nightshade"
	Category    string `json:"category"` // "Овощные"
	ImageUrl    string `json:"imageUrl"`
	Description string `json:"description"` // описание культуры
}
type VarietyItem struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Desc           string   `json:"desc"`
	YieldPotential float64  `json:"yieldPotential"`
	PlantHeight    float64  `json:"plantHeight"`
	GrowingTypes   []string `json:"growingTypes"`
}
type VarietyDetail struct {
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
	WaterRequirement WaterRequirement `json:"waterRequirement"`
	// Световые требования
	LightRequirement LightRequirement `json:"lightRequirement"`
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

type SeasonItem struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type SeasonFilter struct {
	OwnerId string `json:"owner_id,omitempty"`
}

type CatalogProjections interface {
	GetSeasons(context.Context, SeasonFilter) ([]SeasonItem, error)
	GetCrops(context.Context) ([]CropDto, error)
	GetVarieties(context.Context, string) ([]VarietyItem, error)
	GetVariety(context.Context, string) (*VarietyDetail, error)
}
