package agronomy

// Crop represents general crop information
type Crop struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	ScientificName string `json:"scientific_name"`
	BBCHStages  []string `json:"bbch_stages"`
}

// Variety is a specific variety of a crop
type Variety struct {
	ID       string `json:"id"`
	CropID   string `json:"crop_id"`
	Name     string `json:"name"`
	Maturity int    `json:"maturity_days"`
}
