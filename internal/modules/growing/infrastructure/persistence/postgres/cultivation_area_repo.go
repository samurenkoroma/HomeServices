package postgres

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
)

type cultivationAreaRepository struct {
	tx *sql.Tx
}

func NewCultivationAreaRepository(tx *sql.Tx) cultivationarea.Repository {
	return &cultivationAreaRepository{tx: tx}
}

func (c cultivationAreaRepository) Save(ctx context.Context, area cultivationarea.CultivationArea) error {
	//TODO implement me
	panic("implement me")
}

func (c cultivationAreaRepository) FindByID(ctx context.Context, id string) (cultivationarea.CultivationArea, error) {
	//TODO implement me
	panic("implement me")
}

func (c cultivationAreaRepository) FindByFarmRefID(ctx context.Context, farmRefID string) (cultivationarea.CultivationArea, error) {
	//TODO implement me
	panic("implement me")
}

func (c cultivationAreaRepository) FindByType(ctx context.Context, areaType cultivationarea.AreaType) ([]cultivationarea.CultivationArea, error) {
	//TODO implement me
	panic("implement me")
}

func (c cultivationAreaRepository) FindAll(ctx context.Context) ([]cultivationarea.CultivationArea, error) {
	//TODO implement me
	panic("implement me")
}

func (c cultivationAreaRepository) SaveSeasonConfig(ctx context.Context, area cultivationarea.CultivationArea, seasonID string) error {
	//TODO implement me
	panic("implement me")
}

func (c cultivationAreaRepository) GetSeasonConfig(ctx context.Context, areaID, seasonID string) (*cultivationarea.SeasonConfig, error) {
	//TODO implement me
	panic("implement me")
}

func (c cultivationAreaRepository) FindByParentID(ctx context.Context, parentID string) ([]cultivationarea.CultivationArea, error) {
	//TODO implement me
	panic("implement me")
}

func (c cultivationAreaRepository) Exists(ctx context.Context, id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c cultivationAreaRepository) ExistsByFarmRefID(ctx context.Context, farmRefID string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
