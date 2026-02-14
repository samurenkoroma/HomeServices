package db_table

import "gorm.io/gorm"

type InvoiceItem struct {
	gorm.Model
	Name      string
	Quantity  int
	Price     Price `gorm:"embedded;embeddedPrefix:price_"`
	InvoiceID uint
}

type Price struct {
	Value float64
	Vat   float64
	Total float64
}

type Invoice struct {
	gorm.Model
	Items      []InvoiceItem
	File       string
	SupplierId uint
	Supplier   Supplier
}

type Supplier struct {
	gorm.Model
	Name   string
	Site   string
	Rating float64
}
