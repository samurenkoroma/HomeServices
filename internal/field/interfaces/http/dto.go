package http

type CreateLandUnitRequest struct {
	Name     string  `json:"name"`
	ParentID string  `json:"parentId"`
	Unit     string  `json:"unit"`
	Space    string  `json:"space"`
	Length   float64 `json:"length"`
	Width    float64 `json:"width"`
}
