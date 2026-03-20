package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
	"time"
)

type cultivationAreaRepository struct {
	tx *sql.Tx
}

func NewCultivationAreaRepository(tx *sql.Tx) cultivationarea.Repository {
	return &cultivationAreaRepository{tx: tx}
}

func (r *cultivationAreaRepository) Save(ctx context.Context, area cultivationarea.CultivationArea) error {
	// Сохраняем в таблицу cultivation_areas
	query := `
        INSERT INTO cultivation_areas (
            id, farm_ref_id, type, name, geometry, area, created_at, updated_at
        ) VALUES ($1, $2, $3, $4, ST_SetSRID(ST_GeomFromGeoJSON($5),4326), $6, $7, $8)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name,
            geometry = EXCLUDED.geometry,
            area = EXCLUDED.area,
            updated_at = EXCLUDED.updated_at
    `

	geomData, _ := json.Marshal(area.GetGeometry())

	_, err := r.tx.ExecContext(ctx, query,
		area.GetID(),
		area.GetFarmRefID(),
		string(area.GetType()),
		area.GetName(),
		geomData,
		area.GetArea(),
		time.Now(),
		time.Now(),
	)

	return err
}

func (r *cultivationAreaRepository) SaveSeasonConfig(ctx context.Context, area cultivationarea.CultivationArea, seasonID string) error {
	config, err := area.GetSeasonConfig(seasonID)
	if err != nil {
		return err
	}

	query := `
        INSERT INTO area_season_configs (
            area_id, season_id, name, geometry, area, crop_plan_id, block_ids, metadata, valid_from
        ) VALUES ($1, $2, $3, ST_SetSRID(ST_GeomFromGeoJSON($4),4326), $5, $6, $7, $8, $9)
        ON CONFLICT (area_id, season_id) DO UPDATE SET
            name = EXCLUDED.name,
            geometry = EXCLUDED.geometry,
            area = EXCLUDED.area,
            crop_plan_id = EXCLUDED.crop_plan_id,
            block_ids = EXCLUDED.block_ids,
            metadata = EXCLUDED.metadata
    `

	geomData, _ := json.Marshal(config.Geometry)
	metadata, _ := json.Marshal(config.Metadata)

	_, err = r.tx.ExecContext(ctx, query,
		area.GetID(),
		seasonID,
		config.Name,
		geomData,
		config.Area,
		config.CropPlanID,
		config.BlockIDs,
		metadata,
		config.ValidFrom,
	)

	return err
}

func (r *cultivationAreaRepository) GetByFarmRefID(ctx context.Context, farmRefID string) (cultivationarea.CultivationArea, error) {
	query := `
        SELECT id, type, name, ST_AsGeoJSON(geometry), area
        FROM cultivation_areas
        WHERE farm_ref_id = $1
    `

	var (
		id       string
		objType  string
		name     string
		geomJSON string
		area     float64
	)

	err := r.tx.QueryRowContext(ctx, query, farmRefID).Scan(&id, &objType, &name, &geomJSON, &area)
	if err != nil {
		return nil, err
	}

	var geom spatial.GeoJSON
	json.Unmarshal([]byte(geomJSON), &geom)

	// Создаем объект в зависимости от типа
	switch cultivationarea.AreaType(objType) {
	case cultivationarea.AreaTypeField:
		return r.hydrateFieldArea(id, farmRefID, name, geom, area)
	case cultivationarea.AreaTypeGreenhouse:
		return r.hydrateGreenhouseArea(id, farmRefID, name, geom, area)
	default:
		return nil, cultivationarea.ErrUnknownAreaType
	}
}
