package service

import (
	"database/sql"
	"fmt"
	"samurenkoroma/services/pkg/db"
	"samurenkoroma/services/services/storehouse"

	nestedset "github.com/longbridgeapp/nested-set"
)

type SeedService struct {
	Repo     *storehouse.SeedsRepository
	database *db.Db
}

func NewSeedService(repo *storehouse.SeedsRepository, db *db.Db) *SeedService {
	return &SeedService{
		Repo:     repo,
		database: db,
	}
}

func (s *SeedService) Add(dto *storehouse.CreateSeedRequest) (storehouse.CreateSeedResponse, error) {

	parent, err := s.Repo.Crud.Get(fmt.Sprintf("%d", dto.Parent))
	if err != nil && dto.Parent != 0 {
		return storehouse.CreateSeedResponse{}, err
	}

	seed := storehouse.Seed{
		Name: dto.Name,
		Type: dto.Type,
	}
	if parent != nil {
		seed.ParentID = sql.NullInt64{Valid: true, Int64: int64(parent.ID)}
	}

	if err := nestedset.Create(s.database.DB, &seed, parent); err != nil {
		return storehouse.CreateSeedResponse{}, err
	}

	//err := s.Repo.Crud.Save(seed)
	//if err != nil {
	//	return storehouse.CreateSeedResponse{}, err
	//}
	return storehouse.CreateSeedResponse{
		ID:   seed.ID,
		Name: seed.Name,
		Type: seed.Type,
	}, nil
}
func (s *SeedService) List() ([]storehouse.CreateSeedResponse, error) {
	seeds, err := s.Repo.Crud.List("")
	if err != nil {
		return nil, err
	}

	var data []storehouse.CreateSeedResponse
	for _, seed := range seeds {
		data = append(data, storehouse.CreateSeedResponse{
			ID:   seed.ID,
			Name: seed.Name,
			Type: seed.Type,
		})
	}
	return data, nil
}
