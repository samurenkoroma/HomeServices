package valueobjects

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// ProductionUnitID is a value object for Production Unit identifier
type ProductionUnitID string

func (id ProductionUnitID) String() string {
	return string(id)
}
