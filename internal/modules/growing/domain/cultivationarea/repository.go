package cultivationarea

import (
	"context"
)

type Repository interface {
	// CRUD
	Save(ctx context.Context, area CultivationArea) error
	FindByID(ctx context.Context, id string) (CultivationArea, error)
	FindByFarmRefID(ctx context.Context, farmRefID string) (CultivationArea, error)
	FindByType(ctx context.Context, areaType AreaType) ([]CultivationArea, error)
	FindAll(ctx context.Context) ([]CultivationArea, error)

	// Конфигурации
	SaveSeasonConfig(ctx context.Context, area CultivationArea, seasonID string) error
	GetSeasonConfig(ctx context.Context, areaID, seasonID string) (*SeasonConfig, error)

	// Поиск по родителю
	FindByParentID(ctx context.Context, parentID string) ([]CultivationArea, error)

	// Проверки
	Exists(ctx context.Context, id string) (bool, error)
	ExistsByFarmRefID(ctx context.Context, farmRefID string) (bool, error)
}
