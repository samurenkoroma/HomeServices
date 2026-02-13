package payload

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateSupplierRequest struct {
	Name string `json:"name"`
}
type CreateSupplierResponse struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
}

type SupplierGorm struct {
	ID     uuid.UUID `gorm:"type:uuid"`
	Name   string
	Rating float64
	gorm.Model
}

func (SupplierGorm) TableName() string {
	return "suppliers"
}
