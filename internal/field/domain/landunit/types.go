package landunit

type LandUnitID string
type SectionID string
type BedID string

type LandUnitType string

const (
	Field      LandUnitType = "field"
	Greenhouse LandUnitType = "greenhouse"
)
