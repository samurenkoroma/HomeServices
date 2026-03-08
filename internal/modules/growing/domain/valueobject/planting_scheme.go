package valueobject

type PlantingScheme struct {
	RowSpacing   float64
	PlantSpacing float64
}

func (s PlantingScheme) CalculatePlants(d Dimension) int {
	rows := int(d.Width() / s.RowSpacing)
	plantsPerRow := int(d.Length() / s.PlantSpacing)
	return rows * plantsPerRow
}
