package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinkUpdateRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}

type LinkListResponse struct {
	Data  []Link `json:"data"`
	Count int64  `json:"count"`
}
