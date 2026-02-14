package storehouse

type CreateSeedRequest struct {
	Name string `json:"name" validate:"required"`
	Type string `json:"type" validate:"required"`
}
type CreateSeedResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type CreateVendorRequest struct {
	Name string `json:"name" validate:"required"`
	Url  string `json:"url" validate:"required"`
}
type CreateVendorResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}
