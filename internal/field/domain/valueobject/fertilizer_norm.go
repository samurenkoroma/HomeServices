package valueobject

type FertilizerNorm struct {
	KgPerHectare float64
}

func (n FertilizerNorm) ForArea(areaM2 float64) float64 {
	hectare := 10000.0
	return (areaM2 / hectare) * n.KgPerHectare
}
