package cultivationarea

import (
	"context"
)

// Repository — репозиторий для мест выращивания
type Repository interface {
	// CRUD
	Save(ctx context.Context, area CultivationArea) error
	FindByID(ctx context.Context, id string) (CultivationArea, error)
	FindByFarmRefID(ctx context.Context, farmRefID string) (CultivationArea, error)
	FindByType(ctx context.Context, areaType AreaType) ([]CultivationArea, error)
	FindByParentID(ctx context.Context, parentID string) ([]CultivationArea, error)
	FindAll(ctx context.Context) ([]CultivationArea, error)
	Delete(ctx context.Context, id string) error

	// Конфигурации по сезонам
	SaveSeasonConfig(ctx context.Context, areaID string, config SeasonConfig) error
	GetSeasonConfig(ctx context.Context, areaID, seasonID string) (*SeasonConfig, error)
	GetSeasonConfigs(ctx context.Context, areaID string) ([]SeasonConfig, error)
	DeleteSeasonConfig(ctx context.Context, areaID, seasonID string) error

	// Проверки
	Exists(ctx context.Context, id string) (bool, error)
	ExistsByFarmRefID(ctx context.Context, farmRefID string) (bool, error)
	ExistsByName(ctx context.Context, name string) (bool, error)
}
