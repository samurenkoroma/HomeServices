package db_table

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model

	Product   Product
	ProductId uint
	Quantity  float64
	Price     float64
	Total     float64 `gorm:"->;numeric:GENERATED ALWAYS AS (quantity * price);default:(0.0);"`

	OrderID uint
	Order   Order
}

type OrderType uint8

const (
	Purchase OrderType = iota
	Sale
)

type Order struct {
	gorm.Model
	Type       OrderType `gorm:"default:0"`
	Items      []OrderItem
	File       string
	SupplierId uint
	Supplier   Supplier
	Status     string
}
