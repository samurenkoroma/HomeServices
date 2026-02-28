package http

type CreateLandUnitRequest struct {
	Name     string  `json:"name"`
	ParentID string  `json:"parentId"`
	UnitType string  `json:"unitType"`
	LandType string  `json:"landType"`
	Length   float64 `json:"length"`
	Width    float64 `json:"width"`
}
