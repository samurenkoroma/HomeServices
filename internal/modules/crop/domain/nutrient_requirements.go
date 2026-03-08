package domain

type NutrientRequirements struct {
	nitrogen   float64
	phosphorus float64
	potassium  float64
}

func (n NutrientRequirements) Nitrogen() float64 {
	return n.nitrogen
}

func (n NutrientRequirements) Phosphorus() float64 {
	return n.phosphorus
}

func (n NutrientRequirements) Potassium() float64 {
	return n.potassium
}
