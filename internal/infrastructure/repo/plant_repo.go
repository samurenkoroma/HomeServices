package repo

import (
	"context"
	"samurenkoroma/services/internal/domain"
	"samurenkoroma/services/internal/domain/taxonomy"
	"samurenkoroma/services/internal/infrastructure/db_table"

	"gorm.io/gorm"
)

type PlantRepo struct {
	db *gorm.DB
}

func NewPlantRepo(db *gorm.DB) *PlantRepo {
	return &PlantRepo{
		db: db,
	}
}

//func (repo *PlantRepo) GetById(id uint) (plant domain.Plant, err error) {
//	nodes, err := gorm.G[db_table.TaxonomyNode](repo.db).
//		Where("rank = ? AND type = ?", "familia", taxonomy.Plants).
//		First(context.Background())
//	if err != nil {
//		return nil, err
//	}
//
//	for _, node := range nodes {
//		result = append(result, domain.Family{
//			ID:   uint(node.ID),
//			Name: node.Name,
//		})
//	}
//
//	return result, nil
//}

func (repo *PlantRepo) GetFamily(id uint) (domain.Family, error) {
	result, err := gorm.G[db_table.TaxonomyNode](repo.db).
		Where("rank = ? AND type = ? AND ", "familia", taxonomy.Plants).
		First(context.Background())
	if err != nil {
		return domain.Family{}, err
	}

	return domain.Family{
		ID:   result.ID,
		Name: result.Name,
	}, nil
}
func (repo *PlantRepo) GetFamilies() (result []domain.Family, err error) {
	nodes, err := gorm.G[db_table.TaxonomyNode](repo.db).
		Where("rank = ? AND type = ?", "familia", taxonomy.Plants).
		Find(context.Background())
	if err != nil {
		return nil, err
	}

	for _, node := range nodes {
		result = append(result, domain.Family{
			ID:   node.ID,
			Name: node.Name,
		})
	}

	return result, nil
}

func (repo *PlantRepo) GetSpeciesByFamilia(familiaId uint) (result []domain.Species, err error) {
	nodes, err := gorm.G[db_table.TaxonomyNode](repo.db).
		Where("rank = ? AND type = ? AND parent_id = ?", "species", taxonomy.Plants, familiaId).
		Find(context.Background())
	if err != nil {
		return nil, err
	}
	f, _ := repo.GetFamily(familiaId)
	for _, node := range nodes {
		result = append(result, domain.Species{
			ID:     node.ID,
			Name:   node.Name,
			Family: f,
		})
	}

	return result, nil
}
