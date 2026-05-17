package catalog

import (
	"context"
)

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

type CultivationPlansFilter struct {
	CropKey string `json:"cropKey,omitempty"`
}

type CultivationPlanItem struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	CropKey   string  `json:"cropKey"`
	VarietyID *string `json:"varietyID"`
	Version   int     `json:"version"`
	Steps     []Step  `json:"steps"`
}
type Step struct {
	ID      uint8   `json:"id"`
	Trigger Trigger `json:"trigger"`
	Type    string  `json:"type"`
}

type Trigger struct {
	Type  string         `json:"type"`
	Value map[string]any `json:"value"`
}

// CropPlanListItemDTO для списка (компактный)
type CropPlanListItemDTO struct {
	ID                  string     `json:"id"`
	Crop                CropDTO    `json:"crop"`
	Variety             VarietyDTO `json:"variety"`
	ProductionUnit      UnitDTO    `json:"productionUnit"`
	CultivationPlan     PlanRefDTO `json:"cultivationPlan"`
	PlantingDate        string     `json:"plantingDate"`
	Status              string     `json:"status"`
	ExpectedHarvestDate string     `json:"expectedHarvestDate"`
	Progress            int        `json:"progress"`
}

// CropDTO информация о культуре
type CropDTO struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

// VarietyDTO информация о сорте
type VarietyDTO struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	DaysToMaturity int16  `json:"daysToMaturity"`
}

// UnitDTO информация о производственной единице (грядка/поле)
type UnitDTO struct {
	ID   string  `json:"id"`
	Area float64 `json:"area"`
	Name string  `json:"name"`
}

// PlanRefDTO ссылка на шаблон плана
type PlanRefDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CropPlanFilter struct {
	OwnerId          string `json:"owner_id,omitempty"`
	ProductionUnitId string `json:"puid,omitempty"`
}

type CatalogProjections interface {
	GetSeasons(context.Context, SeasonFilter) ([]SeasonItem, error)
	GetCropPlans(context.Context, CropPlanFilter) ([]CropPlanListItemDTO, error)
	GetCultivationPlans(context.Context, CultivationPlansFilter) ([]CultivationPlanItem, error)
	GetCrops(context.Context) ([]CropDto, error)
	GetVarieties(context.Context, string) ([]VarietyItem, error)
	GetVariety(context.Context, string) (*VarietyDetail, error)
}
