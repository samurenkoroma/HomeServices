package croptype

import (
	"context"
	"time"
)

type Filter struct {
	Category      string
	IsActive      bool
	Search        string
	Family        string
	ActiveOnly    bool
	WithVarieties bool
}

type VarietySimpleDTO struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Image          string `json:"image"`
	VegetationDays string `json:"growingTime"`
	YieldPotential string `json:"yieldEstimate"`
	PlantingTime   string `json:"plantingTime"`
}

// CropTypeWithVarietiesDTO — денормализованный объект
type CropTypeWithVarietiesDTO struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Category    string             `json:"category"`
	Family      string             `json:"family"`
	Description string             `json:"description"`
	IsPerennial bool               `json:"isPerennial"`
	ImageUrl    string             `json:"imageURL"`
	Varieties   []VarietySimpleDTO `json:"varieties,omitempty"`
}

// CropTypeSimpleDTO — упрощённый DTO для списка
type CropTypeSimpleDTO struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Category         string `json:"category"`
	Family           string `json:"family"`
	Icon             string `json:"icon"`
	ImageUrl         string `json:"imageURL"`
	CountVarieties   int    `json:"countVarieties"`
	YieldEstimateMin *int   `json:"yieldEstimateMin"`
	YieldEstimateMax *int   `json:"yieldEstimateMax"`
	GrowingDaysMin   *int   `json:"growingDaysMin"`
	GrowingDaysMax   *int   `json:"growingDaysMax"`
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

// CategoryDTO — DTO для категории
type CategoryDTO struct {
	Code          string   `json:"code"`
	Name          string   `json:"name"`
	NameEn        string   `json:"nameEn"`
	Description   string   `json:"description"`
	Parent        *string  `json:"parent,omitempty"`
	Subcategories []string `json:"subcategories"`
}

type Projections interface {
	GetList(ctx context.Context, filter Filter) ([]CropTypeSimpleDTO, error)
	GetByID(ctx context.Context, id string) (*CropTypeDetailDTO, error)
	GetCropTypeWithVarieties(ctx context.Context, id string) (*CropTypeWithVarietiesDTO, error)
	GetAllCropTypesSimple(ctx context.Context) ([]CropTypeSimpleDTO, error)
}
