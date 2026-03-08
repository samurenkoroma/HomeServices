package facility

type FacilityID string
type GrowingAreaID string
type FacilityType string

const (
	FieldFacility      FacilityType = "FIELD"
	GreenhouseFacility FacilityType = "GREENHOUSE"
	BlockFacility      FacilityType = "BLOCK"
	BedFacility        FacilityType = "BED"
)
