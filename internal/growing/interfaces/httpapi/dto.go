package httpapi

type CreateFacilityRequest struct {
	Name     string  `json:"name"`
	ParentID string  `json:"parentId"`
	Unit     string  `json:"unit"`
	Facility string  `json:"facility"`
	Length   float64 `json:"length"`
	Width    float64 `json:"width"`
}
