package storehouse

type CreatePlantRequest struct {
	Name   string `json:"name" validate:"required"`
	Rank   uint   `json:"rank" validate:"required"`
	Parent uint   `json:"parent"`
}

type CreatePlantResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CreateSeedRequest struct {
	Name     string `json:"name" validate:"required"`
	Plant    uint   `json:"plant" validate:"required"`
	Vendor   uint   `json:"vendor" validate:"required"`
	Link     string `json:"link"`
	Variants []struct {
		Weight float64 `json:"weight"`
		Price  float64 `json:"price"`
	} `json:"variants"`
}
type CreateSeedResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Plant    string `json:"plant"`
	Vendor   string `json:"vendor"`
	Link     string `json:"link"`
	Variants []struct {
		Weight float64 `json:"weight"`
		Price  float64 `json:"price"`
	} `json:"variants"`
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
