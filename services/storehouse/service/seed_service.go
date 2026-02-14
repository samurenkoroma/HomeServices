package service

import "samurenkoroma/services/services/storehouse"

type SeedService struct {
	Repo *storehouse.SeedsRepository
}

func NewSeedService(repo *storehouse.SeedsRepository) *SeedService {
	return &SeedService{
		Repo: repo,
	}
}

func (s *SeedService) Add(dto *storehouse.CreateSeedRequest) (storehouse.CreateSeedResponse, error) {
	seed := &storehouse.Seed{
		Name: dto.Name,
		Type: dto.Type,
	}
	err := s.Repo.Crud.Save(seed)
	if err != nil {
		return storehouse.CreateSeedResponse{}, err
	}
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
