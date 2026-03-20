package valueobject

// NutrientRequirements — требования к питательным веществам
type NutrientRequirements struct {
	Nitrogen   float64 `json:"nitrogen"`   // кг/га
	Phosphorus float64 `json:"phosphorus"` // кг/га
	Potassium  float64 `json:"potassium"`  // кг/га
	Calcium    float64 `json:"calcium"`
	Magnesium  float64 `json:"magnesium"`
	Sulfur     float64 `json:"sulfur"`
}
