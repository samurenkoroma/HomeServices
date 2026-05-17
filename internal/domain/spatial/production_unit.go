package spatial

import (
	"time"

	"github.com/samurenkoroma/HomeServices/internal/domain/shared/valueobjects"
)

// ProductionUnit - Aggregate root for Spatial Context
type ProductionUnit struct {
	ID          valueobjects.ProductionUnitID `json:"id"`
	Name        string                        `json:"name"`
	Type        string                        `json:"type"` // field, greenhouse, etc.
	Geometry    valueobjects.Geometry         `json:"geometry"`
	ParentID    *valueobjects.ProductionUnitID `json:"parent_id,omitempty"`
	CreatedAt   time.Time                     `json:"created_at"`
	UpdatedAt   time.Time                     `json:"updated_at"`
}

// NewProductionUnit creates a new Production Unit
func NewProductionUnit(id valueobjects.ProductionUnitID, name, unitType string, geometry valueobjects.Geometry) *ProductionUnit {
	return &ProductionUnit{
		ID:        id,
		Name:      name,
		Type:      unitType,
		Geometry:  geometry,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
