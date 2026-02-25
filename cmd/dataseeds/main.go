package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"samurenkoroma/services/configs"
	"samurenkoroma/services/internal/infrastructure/db_table"
	"samurenkoroma/services/internal/usecase/taxa/create"
	"samurenkoroma/services/pkg/db"

	"gorm.io/gorm"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)

	file, err := os.ReadFile("./data/db_seeds.json")
	if err != nil {
		return
	}
	var data Data
	if err := json.Unmarshal(file, &data); err != nil {
		return
	}
	tables := []interface{}{
		db_table.ProductVariant{},
		db_table.Product{},
		db_table.TaxonomyNode{},
		db_table.Supplier{},
	}

	newSeedManager(database.DB).cleanTables(tables).execute(data)
}

type SeedManager struct {
	database    *gorm.DB
	nodeCreator *create.UC
	ctx         context.Context
}

func newSeedManager(database *gorm.DB) *SeedManager {
	return &SeedManager{
		database:    database,
		nodeCreator: create.NewUC(database),
		ctx:         context.Background(),
	}
}

func (m *SeedManager) execute(data Data) {
	m.addNode(data.Nodes, nil)
	m.addSuppliers(data.Suppliers)
}

func (m *SeedManager) cleanTables(tables []interface{}) *SeedManager {
	for _, t := range tables {
		m.database.Unscoped().Where("id > 0").Delete(&t)
	}
	return m
}

func (m *SeedManager) addSuppliers(suppliers []Supplier) {
	for _, s := range suppliers {
		gorm.G[db_table.Supplier](m.database).Create(m.ctx, &db_table.Supplier{
			Name: s.Name,
			Site: s.Site,
		})
	}
}

func (m *SeedManager) addNode(data []Node, parent *db_table.TaxonomyNode) {
	for _, seed := range data {
		payload := &create.Payload{
			Name:   seed.Name,
			Rank:   seed.Rank,
			Type:   seed.Type,
			Parent: parent,
		}
		node, err := m.nodeCreator.Payload(payload).Execute(m.ctx)
		if err != nil {
			fmt.Println(err)
		}
		if seed.Items != nil {
			m.addNode(seed.Items, node)
		}
	}
}
