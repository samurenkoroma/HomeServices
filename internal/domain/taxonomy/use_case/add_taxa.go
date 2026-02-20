package use_case

import (
	"database/sql"
	"samurenkoroma/services/internal/domain/taxonomy"
	"samurenkoroma/services/internal/infrastructure/db_table"

	nestedset "github.com/longbridgeapp/nested-set"
	"gorm.io/gorm"
)

type UC struct {
	Conn *gorm.DB
}

func (uc *UC) AddPlant(name string, parent *gorm.Model) error {
	return uc.addTaxonomy(name, parent, taxonomy.Plants)
}

func (uc *UC) AddAnimal(name string, parent *gorm.Model) error {
	return uc.addTaxonomy(name, parent, taxonomy.Animals)
}
func (uc *UC) AddTools(name string, parent *gorm.Model) error {
	return uc.addTaxonomy(name, parent, taxonomy.Tools)
}

func (uc *UC) addTaxonomy(name string, parent *gorm.Model, taxa taxonomy.TypeTaxonomy) error {
	plant := db_table.TaxonomyNode{
		Name: name,
		Type: uint(taxa),
	}
	if parent.ID != 0 {
		plant.ParentID = sql.NullInt64{Valid: true, Int64: int64(parent.ID)}
	}
	if err := nestedset.Create(uc.Conn, &plant, parent); err != nil {
		return err
	}

	return nil
}
