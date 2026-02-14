package accountant

type CreateSupplierRequest struct {
	Name string `json:"name"`
	Site string `json:"site"`
}
type CreateSupplierResponse struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
}
