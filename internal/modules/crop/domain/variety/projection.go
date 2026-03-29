package variety

import (
	"context"
	"samurenkoroma/services/internal/modules/crop/domain/valueobject"
	"time"
)

type Filter struct {
	Search     string
	Limit      int
	Offset     int
	CropTypeId string
	IsActive   bool
}

type VarietyDTO struct {
	ID                 string             `json:"id"`
	CropTypeID         string             `json:"crop_type_id"`
	CropTypeName       string             `json:"crop_type_name"`
	Name               string             `json:"name"`
	Description        string             `json:"description"`
	VegetationDays     valueobject.MinMax `json:"vegetation_days"`
	YieldPotential     valueobject.MinMax `json:"yield_potential"`
	DiseaseResistance  []string           `json:"disease_resistance"`
	RecommendedRegions []string           `json:"recommended_regions"`
	PlantingDensity    int                `json:"planting_density"`
	SeedRate           float64            `json:"seed_rate"`
	Breeder            string             `json:"breeder"`
	YearReleased       int                `json:"year_released"`
	IsActive           bool               `json:"is_active"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
}

type Projections interface {
	GetList(context.Context, Filter) ([]*VarietyDTO, error)
	GetByID(context.Context, string) (*VarietyDTO, error)
}
