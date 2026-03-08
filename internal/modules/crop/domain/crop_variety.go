package domain

type CropVarietyID string

type CropVariety struct {
	id             CropVarietyID
	cropTypeID     CropTypeID
	name           string
	vegetationDays int
	potentialYield float64
}

func NewCropVariety(
	id CropVarietyID,
	cropTypeID CropTypeID,
	name string,
	vegetationDays int,
	potentialYield float64,
) *CropVariety {
	return &CropVariety{
		id:             id,
		cropTypeID:     cropTypeID,
		name:           name,
		vegetationDays: vegetationDays,
		potentialYield: potentialYield,
	}
}
