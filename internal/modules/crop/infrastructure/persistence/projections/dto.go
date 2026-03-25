package projections

import "time"

// CropTypeWithVarietiesDTO — денормализованный объект
type CropTypeWithVarietiesDTO struct {
	ID             string             `json:"id"`
	Name           string             `json:"name"`
	Category       string             `json:"category"`
	CategoryName   string             `json:"category_name"`
	IsPerennial    bool               `json:"is_perennial"`
	VarietiesCount int                `json:"varieties_count"`
	Varieties      []VarietySimpleDTO `json:"varieties,omitempty"`
	CreatedAt      time.Time          `json:"created_at"`
}

type VarietySimpleDTO struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	VegetationDays string `json:"vegetation_days"`
	IsActive       bool   `json:"is_active"`
}

// CropTypeSimpleDTO — упрощённый DTO для списка
type CropTypeSimpleDTO struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	CategoryName string `json:"category_name"`
	IsPerennial  bool   `json:"is_perennial"`
}

// CropTypeDetailDTO — детальный DTO
type CropTypeDetailDTO struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	ScientificName string    `json:"scientific_name"`
	Category       string    `json:"category"`
	CategoryName   string    `json:"category_name"`
	Description    string    `json:"description"`
	IsPerennial    bool      `json:"is_perennial"`
	IsActive       bool      `json:"is_active"`
	VarietiesCount int       `json:"varieties_count"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
type MinMax struct {
	Min     int `json:"min"`
	Max     int `json:"max"`
	Optimal int `json:"optimal"`
}

// VarietyDTO — DTO для ответа
type VarietyDTO struct {
	ID                 string    `json:"id"`
	CropTypeID         string    `json:"crop_type_id"`
	CropTypeName       string    `json:"crop_type_name"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	VegetationDays     MinMax    `json:"vegetation_days"`
	YieldPotential     MinMax    `json:"yield_potential"`
	DiseaseResistance  []string  `json:"disease_resistance"`
	RecommendedRegions []string  `json:"recommended_regions"`
	PlantingDensity    int       `json:"planting_density"`
	SeedRate           float64   `json:"seed_rate"`
	Breeder            string    `json:"breeder"`
	YearReleased       int       `json:"year_released"`
	IsActive           bool      `json:"is_active"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
