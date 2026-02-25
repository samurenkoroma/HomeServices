package db_table

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Name string
	Site string
}
