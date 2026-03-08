package domain

type CropTypeID string

type CropType struct {
	id             CropTypeID
	name           string
	scientificName string
	category       string
	rootDepth      int
	isPerennial    bool
}

func NewCropType(
	id CropTypeID,
	name string,
	scientificName string,
	category string,
	rootDepth int,
	isPerennial bool,
) *CropType {
	return &CropType{
		id:             id,
		name:           name,
		scientificName: scientificName,
		category:       category,
		rootDepth:      rootDepth,
		isPerennial:    isPerennial,
	}
}

func (c *CropType) ID() CropTypeID { return c.id }
func (c *CropType) Name() string   { return c.name }
