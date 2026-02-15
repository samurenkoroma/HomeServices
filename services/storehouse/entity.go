package storehouse

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Seed struct {
	ID   int `gorm:"primarykey" nestedset:"id"`
	Name string
	Type string
	gorm.Model
	Vendors       []*Vendor     `gorm:"many2many:vendor_seeds;"`
	ParentID      sql.NullInt64 `nestedset:"parent_id"`
	Rgt           int           `nestedset:"rgt"`
	Lft           int           `nestedset:"lft"`
	Depth         int           `nestedset:"depth"`
	ChildrenCount int           `nestedset:"children_count"`
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
