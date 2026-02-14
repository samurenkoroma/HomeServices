package storehouse

import (
	"time"

	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Type string
	gorm.Model
	Vendors []*Vendor `gorm:"many2many:vendor_seeds;"`
}

type Vendor struct {
	Name string
	Url  string
	gorm.Model
	Seeds []*Seed `gorm:"many2many:vendor_seeds;"`
}

type VendorSeeds struct {
	Weight    float64
	Price     float64
	Link      string
	SeedID    uint `gorm:"primaryKey"`
	VendorID  uint `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
