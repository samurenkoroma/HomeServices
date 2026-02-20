package service

import (
	"context"
	"database/sql"
	"log"
	"samurenkoroma/services/internal/domain/taxonomy"
	"samurenkoroma/services/internal/domain/taxonomy/use_case"
	"samurenkoroma/services/internal/infrastructure/db_table"
	"samurenkoroma/services/internal/infrastructure/repo"
	"samurenkoroma/services/pkg/db"
	"samurenkoroma/services/pkg/di"
	"samurenkoroma/services/services/storehouse"

	nestedset "github.com/longbridgeapp/nested-set"
	"gorm.io/gorm"
)

type SeedService struct {
	database    *db.Db
	plantsRepo  di.CRUDRepository[db_table.TaxonomyNode]
	vendorsRepo di.CRUDRepository[db_table.Vendor]
	seedsRepo   di.CRUDRepository[db_table.Seed]
	uc          use_case.UC
}

func NewSeedService(db *db.Db) *SeedService {
	return &SeedService{
		database:    db,
		plantsRepo:  repo.NewCrudRepo[db_table.TaxonomyNode](db),
		vendorsRepo: repo.NewCrudRepo[db_table.Vendor](db),
		seedsRepo:   repo.NewCrudRepo[db_table.Seed](db),
		uc:          use_case.UC{Conn: db.DB},
	}
}

func (s *SeedService) AddSeed(dto *storehouse.CreateSeedRequest) error {
	plant, err := s.plantsRepo.Get(dto.Plant)
	if err != nil {
		return err
	}

	vendor, err := s.vendorsRepo.Get(dto.Vendor)
	if err != nil {
		return err
	}

	var seed = db_table.Seed{
		Link:   dto.Link,
		Plant:  plant,
		Vendor: vendor,
	}
	s.database.Create(&seed)
	var variants []db_table.SeedVariant
	for _, v := range dto.Variants {
		temp := db_table.SeedVariant{
			Weight: v.Weight,
			Price:  v.Price,
			Seed:   seed,
		}
		variants = append(variants, temp)
		s.database.DB.Create(&temp)
	}

	return nil
}

func (s *SeedService) AddPlant(dto *storehouse.CreatePlantRequest) error {
	parent, err := s.plantsRepo.Get(dto.Parent)
	plant := db_table.TaxonomyNode{
		Name: dto.Name,
		Type: uint(taxonomy.Plants),
		Rank: dto.Rank,
	}

	if err != nil {
		parent = nil
	} else {
		plant.ParentID = sql.NullInt64{Valid: true, Int64: parent.ID}
	}

	if err := nestedset.Create(s.database.DB, &plant, parent); err != nil {
		log.Print(err)
		return err
	}
	return nil
}
func (s *SeedService) List() ([]storehouse.CreateSeedResponse, error) {
	//result := repo.database.DB.Preload("Resources").Preload("Authors").First(&book, "id = ?", id)
	//var seeds []*db_table.Seed
	seeds, err := gorm.G[db_table.Seed](s.database.DB).Preload("Plant", nil).Preload("Vendor", nil).Find(context.Background())
	if err != nil {

		return nil, err
	}
	var data []storehouse.CreateSeedResponse
	for _, seed := range seeds {
		data = append(data, storehouse.CreateSeedResponse{
			ID:   seed.ID,
			Name: seed.Plant.Name,
		})
	}
	return data, nil
}
