package db_table

import (
	"gorm.io/gorm"
)

type Product struct {
	Link     string
	Name     string
	Category string

	gorm.Model
	Variants []ProductVariant

	Supplier   Supplier
	SupplierID int64
}

type ProductVariant struct {
	Weight    float64
	Price     float64
	Product   Product
	ProductId int64
}
