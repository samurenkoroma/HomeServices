package db_table

import (
	"database/sql"

	"gorm.io/gorm"
)

/*
id — уникальный идентификатор таксона
parent_id — ссылка на родительский уровень (иерархия tree structure)

	например: вид → род → семейство → отряд

type - тип таксономии

	plants
	animals
	tools

rank — уровень классификации:

	domain
	kingdom
	phylum
	class
	order
	family
	genus
	species

scientific_name — научное имя (например Homo sapiens)
common_name — обычное название (человек разумный)
latin_name — иногда полезно отдельно (можно объединить с scientific_name)
is_extinct — вымерший вид или нет
description — свободное описание
*/
type TaxonomyNode struct {
	ID   int64 `gorm:"primarykey" nestedset:"id"`
	Type uint
	Rank uint
	Name string
	gorm.Model

	ParentID      sql.NullInt64 `nestedset:"parent_id"`
	Rgt           int           `nestedset:"rgt"`
	Lft           int           `nestedset:"lft"`
	Depth         int           `nestedset:"depth"`
	ChildrenCount int           `nestedset:"children_count"`
}

type Seed struct {
	Link     string
	Plant    *TaxonomyNode
	PlantID  int64
	Vendor   *Vendor
	VendorID int64
	gorm.Model
	Variants []SeedVariant
}

type SeedVariant struct {
	Weight float64
	Price  float64
	Seed   Seed
	SeedId int64
}
type Vendor struct {
	Name string
	Url  string
	gorm.Model
	Seeds []*Seed
}
