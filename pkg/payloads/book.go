package payloads

type BookRequest struct {
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}
type ResourceResponse struct {
	Type uint   `json:"type"`
	Link string `json:"link"`
}

type AuthorResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type BookResponse struct {
	Id        uint               `json:"id"`
	Title     string             `json:"title"`
	Authors   []AuthorResponse   `json:"authors"`
	Resources []ResourceResponse `json:"resources"`
}

type BookListResponse struct {
	Data  []BookResponse `json:"data"`
	Count int            `json:"count"`
}
