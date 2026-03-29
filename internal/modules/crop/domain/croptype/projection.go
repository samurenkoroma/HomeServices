package croptype

import (
	"context"
	"time"
)

type Filter struct {
	Category string
	IsActive bool
	Search   string
}

type VarietySimpleDTO struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	VegetationDays string `json:"vegetation_days"`
	IsActive       bool   `json:"is_active"`
}

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
	Description    string             `json:"description"`
	IsActive       bool               `json:"is_active"`
}

// CropTypeSimpleDTO — упрощённый DTO для списка
type CropTypeSimpleDTO struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	CategoryName string `json:"category_name"`
	IsPerennial  bool   `json:"is_perennial"`
	IsActive     bool   `json:"is_active"`
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
