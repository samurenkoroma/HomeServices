package variety

import (
	"context"
	"samurenkoroma/services/internal/modules/crop/domain/valueobject"
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
	CropTypeID         string             `json:"cropTypeId"`
	CropTypeName       string             `json:"cropTypeName"`
	Name               string             `json:"name"`
	Description        string             `json:"description"`
	VegetationDays     valueobject.MinMax `json:"vegetationDays"`
	YieldPotential     valueobject.MinMax `json:"yieldPotential"`
	DiseaseResistance  []string           `json:"disease_resistance"`
	RecommendedRegions []string           `json:"recommended_regions"`
	PlantingDensity    int                `json:"planting_density"`
	SeedRate           float64            `json:"seed_rate"`
	Image              string             `json:"image"`
}

type Projections interface {
	GetList(context.Context, Filter) ([]*VarietyDTO, error)
	GetByID(context.Context, string) (*VarietyDTO, error)
}
