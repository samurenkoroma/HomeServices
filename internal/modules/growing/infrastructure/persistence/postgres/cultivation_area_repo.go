package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"samurenkoroma/services/internal/core/domain/types"
	"time"

	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
)

type cultivationAreaRepository struct {
	tx *sql.Tx
}

func NewCultivationAreaRepository(tx *sql.Tx) cultivationarea.Repository {
	return &cultivationAreaRepository{tx: tx}
}

// Save сохраняет место выращивания
// Save сохраняет место выращивания
func (r *cultivationAreaRepository) Save(ctx context.Context, area cultivationarea.CultivationArea) error {
	query := `
        INSERT INTO public.growing_cultivation_areas (
            id, farm_ref_id, type, name, geometry, area, attributes, created_at, updated_at
        ) VALUES ($1, $2, $3, $4, ST_SetSRID(ST_GeomFromGeoJSON($5), 4326), $6, $7, $8, $9)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name,
            geometry = EXCLUDED.geometry,
            area = EXCLUDED.area,
            attributes = EXCLUDED.attributes,
            updated_at = EXCLUDED.updated_at
    `

	geomData, err := json.Marshal(area.GetGeometry())
	if err != nil {
		return fmt.Errorf("failed to marshal geometry: %w", err)
	}

	var attributesJSON []byte

	switch a := area.(type) {
	case *cultivationarea.Bed:
		// Сохраняем атрибуты грядки
		attrs := a.GetAttributes()
		attributesJSON, err = json.Marshal(attrs)
		if err != nil {
			return fmt.Errorf("failed to marshal bed attributes: %w", err)
		}
	case *cultivationarea.Block:
	}

	_, err = r.tx.ExecContext(ctx, query,
		area.GetID(),
		area.GetFarmRefID(),
		string(area.GetType()),
		area.GetName(),
		string(geomData),
		area.GetArea(),
		attributesJSON,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("failed to save cultivation area: %w", err)
	}

	return nil
}

// FindByID находит место по ID
func (r *cultivationAreaRepository) FindByID(ctx context.Context, id string) (cultivationarea.CultivationArea, error) {
	query := `
        SELECT id, farm_ref_id, type, name, ST_AsGeoJSON(geometry), area, parent_id, created_at, updated_at
        FROM public.growing_cultivation_areas
        WHERE id = $1
    `

	var (
		areaID    string
		farmRefID string
		areaType  string
		name      string
		geomJSON  string
		areaValue float64
		parentID  sql.NullString
		createdAt time.Time
		updatedAt time.Time
	)

	err := r.tx.QueryRowContext(ctx, query, id).Scan(
		&areaID, &farmRefID, &areaType, &name, &geomJSON, &areaValue, &parentID,
		&createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, cultivationarea.ErrAreaNotFound
		}
		return nil, fmt.Errorf("failed to find cultivation area: %w", err)
	}

	var geom spatial.GeoJSON
	if err := json.Unmarshal([]byte(geomJSON), &geom); err != nil {
		return nil, fmt.Errorf("failed to unmarshal geometry: %w", err)
	}
	seasons, err := r.GetSeasonConfigs(ctx, areaID)
	if err != nil {
		return nil, err
	}
	return r.hydrateArea(areaID, farmRefID, cultivationarea.AreaType(areaType), name, geom, areaValue, parentID, createdAt, updatedAt, seasons)
}

// FindByFarmRefID находит место по ссылке на farm модуль
func (r *cultivationAreaRepository) FindByFarmRefID(ctx context.Context, farmRefID string) (cultivationarea.CultivationArea, error) {
	query := `
        SELECT id, farm_ref_id, type, name, ST_AsGeoJSON(geometry), area, parent_id, created_at, updated_at
        FROM public.growing_cultivation_areas
        WHERE farm_ref_id = $1
        LIMIT 1
    `

	var (
		areaID    string
		refID     string
		areaType  string
		name      string
		geomJSON  string
		areaValue float64
		parentID  sql.NullString
		createdAt time.Time
		updatedAt time.Time
	)

	err := r.tx.QueryRowContext(ctx, query, farmRefID).Scan(
		&areaID, &refID, &areaType, &name, &geomJSON, &areaValue, &parentID,
		&createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find cultivation area by farm_ref_id: %w", err)
	}

	var geom spatial.GeoJSON
	if err := json.Unmarshal([]byte(geomJSON), &geom); err != nil {
		return nil, fmt.Errorf("failed to unmarshal geometry: %w", err)
	}
	seasons, err := r.GetSeasonConfigs(ctx, areaID)
	if err != nil {
		return nil, err
	}
	return r.hydrateArea(areaID, refID, cultivationarea.AreaType(areaType), name, geom, areaValue, parentID, createdAt, updatedAt, seasons)
}

// FindByType возвращает все места указанного типа
func (r *cultivationAreaRepository) FindByType(ctx context.Context, areaType cultivationarea.AreaType) ([]cultivationarea.CultivationArea, error) {
	query := `
        SELECT id, farm_ref_id, type, name, ST_AsGeoJSON(geometry), area, parent_id, created_at, updated_at
        FROM public.growing_cultivation_areas
        WHERE type = $1
        ORDER BY name
    `

	rows, err := r.tx.QueryContext(ctx, query, string(areaType))
	if err != nil {
		return nil, fmt.Errorf("failed to query cultivation areas by type: %w", err)
	}
	defer rows.Close()

	var areas []cultivationarea.CultivationArea

	for rows.Next() {
		var (
			areaID    string
			farmRefID string
			at        string
			name      string
			geomJSON  string
			areaValue float64
			parentID  sql.NullString
			createdAt time.Time
			updatedAt time.Time
		)

		err := rows.Scan(
			&areaID, &farmRefID, &at, &name, &geomJSON, &areaValue, &parentID,
			&createdAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cultivation area: %w", err)
		}

		var geom spatial.GeoJSON
		if err := json.Unmarshal([]byte(geomJSON), &geom); err != nil {
			return nil, fmt.Errorf("failed to unmarshal geometry: %w", err)
		}
		seasons, err := r.GetSeasonConfigs(ctx, areaID)
		if err != nil {
			return nil, err
		}
		cultArea, err := r.hydrateArea(areaID, farmRefID, cultivationarea.AreaType(at), name, geom, areaValue, parentID, createdAt, updatedAt, seasons)
		if err != nil {
			return nil, err
		}

		areas = append(areas, cultArea)
	}

	return areas, nil
}

// FindByParentID возвращает все места с указанным родителем
func (r *cultivationAreaRepository) FindByParentID(ctx context.Context, parentID string) ([]cultivationarea.CultivationArea, error) {
	query := `
        SELECT id, farm_ref_id, type, name, ST_AsGeoJSON(geometry), area, parent_id, created_at, updated_at
        FROM public.growing_cultivation_areas
        WHERE parent_id = $1
        ORDER BY name
    `

	rows, err := r.tx.QueryContext(ctx, query, parentID)
	if err != nil {
		return nil, fmt.Errorf("failed to query cultivation areas by parent: %w", err)
	}
	defer rows.Close()

	var areas []cultivationarea.CultivationArea

	for rows.Next() {
		var (
			areaID    string
			farmRefID string
			at        string
			name      string
			geomJSON  string
			areaValue float64
			pid       sql.NullString
			createdAt time.Time
			updatedAt time.Time
		)

		err := rows.Scan(
			&areaID, &farmRefID, &at, &name, &geomJSON, &areaValue, &pid,
			&createdAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cultivation area: %w", err)
		}

		var geom spatial.GeoJSON
		if err := json.Unmarshal([]byte(geomJSON), &geom); err != nil {
			return nil, fmt.Errorf("failed to unmarshal geometry: %w", err)
		}
		seasons, err := r.GetSeasonConfigs(ctx, areaID)
		if err != nil {
			return nil, err
		}
		cultArea, err := r.hydrateArea(areaID, farmRefID, cultivationarea.AreaType(at), name, geom, areaValue, pid, createdAt, updatedAt, seasons)
		if err != nil {
			return nil, err
		}

		areas = append(areas, cultArea)
	}

	return areas, nil
}

// FindAll возвращает все места выращивания
func (r *cultivationAreaRepository) FindAll(ctx context.Context) ([]cultivationarea.CultivationArea, error) {
	query := `
        SELECT id, farm_ref_id, type, name, ST_AsGeoJSON(geometry), area, parent_id, created_at, updated_at
        FROM growing_cultivation_areas
        ORDER BY type, name
    `

	rows, err := r.tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query cultivation areas: %w", err)
	}
	defer rows.Close()

	var areas []cultivationarea.CultivationArea

	for rows.Next() {
		var (
			areaID    string
			farmRefID string
			at        string
			name      string
			geomJSON  string
			areaValue float64
			parentID  sql.NullString
			createdAt time.Time
			updatedAt time.Time
		)

		err := rows.Scan(
			&areaID, &farmRefID, &at, &name, &geomJSON, &areaValue, &parentID,
			&createdAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cultivation area: %w", err)
		}

		var geom spatial.GeoJSON
		if err := json.Unmarshal([]byte(geomJSON), &geom); err != nil {
			return nil, fmt.Errorf("failed to unmarshal geometry: %w", err)
		}
		seasons, err := r.GetSeasonConfigs(ctx, areaID)
		if err != nil {
			return nil, err
		}
		cultArea, err := r.hydrateArea(areaID, farmRefID, cultivationarea.AreaType(at), name, geom, areaValue, parentID, createdAt, updatedAt, seasons)
		if err != nil {
			return nil, err
		}

		areas = append(areas, cultArea)
	}

	return areas, nil
}

// Delete удаляет место выращивания
func (r *cultivationAreaRepository) Delete(ctx context.Context, id string) error {
	var childCount int
	err := r.tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM growing_cultivation_areas WHERE parent_id = $1`, id).Scan(&childCount)
	if err != nil {
		return fmt.Errorf("failed to check children: %w", err)
	}

	if childCount > 0 {
		return fmt.Errorf("cannot delete area with children")
	}

	_, err = r.tx.ExecContext(ctx, `DELETE FROM public.growing_area_season_configs WHERE area_id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete season configs: %w", err)
	}

	_, err = r.tx.ExecContext(ctx, `DELETE FROM growing_cultivation_areas WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete cultivation area: %w", err)
	}

	return nil
}

// SaveSeasonConfig сохраняет конфигурацию места на сезон
func (r *cultivationAreaRepository) SaveSeasonConfig(ctx context.Context, areaID string, config cultivationarea.SeasonConfig) error {
	log.Printf("SaveSeasonConfig: areaID=%s, seasonID=%s, cropPlanID=%v", areaID, config.SeasonID, config.CropPlanID)

	query := `
        INSERT INTO public.growing_area_season_configs (
            area_id, season_id, name, geometry, area, crop_plan_id, block_ids, metadata, valid_from, valid_until
        ) VALUES ($1, $2, $3, ST_SetSRID(ST_GeomFromGeoJSON($4), 4326), $5, $6, $7, $8, $9, $10)
        ON CONFLICT (area_id, season_id) DO UPDATE SET
            name = EXCLUDED.name,
            geometry = EXCLUDED.geometry,
            area = EXCLUDED.area,
            crop_plan_id = EXCLUDED.crop_plan_id,
            block_ids = EXCLUDED.block_ids,
            metadata = EXCLUDED.metadata,
            valid_until = EXCLUDED.valid_until
    `

	geomData, err := json.Marshal(config.Geometry)
	if err != nil {
		return err
	}

	metadata, err := json.Marshal(config.Metadata)
	if err != nil {
		return err
	}

	blockIDsJSON, err := json.Marshal(config.BlockIDs)
	if err != nil {
		return err
	}

	_, err = r.tx.ExecContext(ctx, query,
		areaID,
		config.SeasonID,
		config.Name,
		string(geomData),
		config.Area,
		config.CropPlanID,
		blockIDsJSON,
		metadata,
		config.ValidFrom,
		config.ValidUntil,
	)

	if err != nil {
		log.Printf("SQL error: %v", err)
		return fmt.Errorf("failed to save season config: %w", err)
	}

	log.Printf("Season config saved successfully")
	return nil
}

// GetSeasonConfig получает конфигурацию места на сезон
func (r *cultivationAreaRepository) GetSeasonConfig(ctx context.Context, areaID, seasonID string) (*cultivationarea.SeasonConfig, error) {
	query := `
        SELECT name, ST_AsGeoJSON(geometry), area, crop_plan_id, block_ids, metadata, valid_from, valid_until
        FROM public.growing_area_season_configs
        WHERE area_id = $1 AND season_id = $2
    `

	var (
		name         string
		geomJSON     string
		areaValue    float64
		cropPlanID   sql.NullString
		blockIDsJSON []byte
		metadataJSON []byte
		validFrom    time.Time
		validUntil   sql.NullTime
	)

	err := r.tx.QueryRowContext(ctx, query, areaID, seasonID).Scan(
		&name, &geomJSON, &areaValue, &cropPlanID, &blockIDsJSON, &metadataJSON,
		&validFrom, &validUntil,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, cultivationarea.ErrSeasonConfigNotFound
		}
		return nil, fmt.Errorf("failed to get season config: %w", err)
	}

	var geom spatial.GeoJSON
	if err := json.Unmarshal([]byte(geomJSON), &geom); err != nil {
		return nil, fmt.Errorf("failed to unmarshal geometry: %w", err)
	}

	var blockIDs []string
	if len(blockIDsJSON) > 0 {
		if err := json.Unmarshal(blockIDsJSON, &blockIDs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal block_ids: %w", err)
		}
	}

	var metadata map[string]interface{}
	if len(metadataJSON) > 0 {
		if err := json.Unmarshal(metadataJSON, &metadata); err != nil {
			return nil, fmt.Errorf("failed to unmarshal metadata: %w", err)
		}
	}

	var cropPlanIDPtr *string
	if cropPlanID.Valid {
		cropPlanIDPtr = &cropPlanID.String
	}

	var validUntilPtr *time.Time
	if validUntil.Valid {
		validUntilPtr = &validUntil.Time
	}

	config := &cultivationarea.SeasonConfig{
		SeasonID:   seasonID,
		Name:       name,
		Geometry:   geom,
		Area:       areaValue,
		CropPlanID: cropPlanIDPtr,
		BlockIDs:   blockIDs,
		Metadata:   metadata,
		ValidFrom:  validFrom,
		ValidUntil: validUntilPtr,
	}

	return config, nil
}

// GetSeasonConfigs получает все конфигурации места
func (r *cultivationAreaRepository) GetSeasonConfigs(ctx context.Context, areaID string) ([]cultivationarea.SeasonConfig, error) {
	query := `
        SELECT season_id, name, ST_AsGeoJSON(geometry), area, crop_plan_id, block_ids, metadata, valid_from, valid_until
        FROM public.growing_area_season_configs
        WHERE area_id = $1
        ORDER BY season_id DESC
    `

	rows, err := r.tx.QueryContext(ctx, query, areaID)
	if err != nil {
		return nil, fmt.Errorf("failed to query season configs: %w", err)
	}
	defer rows.Close()

	var configs []cultivationarea.SeasonConfig

	for rows.Next() {
		var (
			seasonID     string
			name         string
			geomJSON     string
			areaValue    float64
			cropPlanID   sql.NullString
			blockIDsJSON []byte
			metadataJSON []byte
			validFrom    time.Time
			validUntil   sql.NullTime
		)

		err := rows.Scan(
			&seasonID, &name, &geomJSON, &areaValue, &cropPlanID, &blockIDsJSON, &metadataJSON,
			&validFrom, &validUntil,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan season config: %w", err)
		}

		var geom spatial.GeoJSON
		if err := json.Unmarshal([]byte(geomJSON), &geom); err != nil {
			return nil, fmt.Errorf("failed to unmarshal geometry: %w", err)
		}

		var blockIDs []string
		if len(blockIDsJSON) > 0 {
			if err := json.Unmarshal(blockIDsJSON, &blockIDs); err != nil {
				return nil, fmt.Errorf("failed to unmarshal block_ids: %w", err)
			}
		}

		var metadata map[string]interface{}
		if len(metadataJSON) > 0 {
			if err := json.Unmarshal(metadataJSON, &metadata); err != nil {
				return nil, fmt.Errorf("failed to unmarshal metadata: %w", err)
			}
		}

		var cropPlanIDPtr *string
		if cropPlanID.Valid {
			cropPlanIDPtr = &cropPlanID.String
		}

		var validUntilPtr *time.Time
		if validUntil.Valid {
			validUntilPtr = &validUntil.Time
		}

		configs = append(configs, cultivationarea.SeasonConfig{
			SeasonID:   seasonID,
			Name:       name,
			Geometry:   geom,
			Area:       areaValue,
			CropPlanID: cropPlanIDPtr,
			BlockIDs:   blockIDs,
			Metadata:   metadata,
			ValidFrom:  validFrom,
			ValidUntil: validUntilPtr,
		})
	}

	return configs, nil
}

// DeleteSeasonConfig удаляет конфигурацию места на сезон
func (r *cultivationAreaRepository) DeleteSeasonConfig(ctx context.Context, areaID, seasonID string) error {
	_, err := r.tx.ExecContext(ctx, `DELETE FROM public.growing_area_season_configs WHERE area_id = $1 AND season_id = $2`, areaID, seasonID)
	if err != nil {
		return fmt.Errorf("failed to delete season config: %w", err)
	}
	return nil
}

// Exists проверяет существование места
func (r *cultivationAreaRepository) Exists(ctx context.Context, id string) (bool, error) {
	var exists bool
	err := r.tx.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM growing_cultivation_areas WHERE id = $1)`, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check existence: %w", err)
	}
	return exists, nil
}

// ExistsByFarmRefID проверяет существование места по ссылке на farm
func (r *cultivationAreaRepository) ExistsByFarmRefID(ctx context.Context, farmRefID string) (bool, error) {
	var exists bool
	err := r.tx.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM growing_cultivation_areas WHERE farm_ref_id = $1)`, farmRefID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check existence: %w", err)
	}
	return exists, nil
}

// ExistsByName проверяет существование места по имени
func (r *cultivationAreaRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var exists bool
	err := r.tx.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM growing_cultivation_areas WHERE name = $1)`, name).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check existence: %w", err)
	}
	return exists, nil
}

// hydrateArea восстанавливает объект CultivationArea из данных
func (r *cultivationAreaRepository) hydrateArea(
	id, farmRefID string,
	areaType cultivationarea.AreaType,
	name string,
	geom spatial.GeoJSON,
	areaValue float64,
	parentID sql.NullString,
	createdAt, updatedAt time.Time,
	seasons []cultivationarea.SeasonConfig,
) (cultivationarea.CultivationArea, error) {
	switch areaType {
	case cultivationarea.AreaTypeField:
		field := cultivationarea.NewFieldArea(farmRefID, name, geom, areaValue)
		field.Rehydrate(id, createdAt, updatedAt, seasons)
		return field, nil

	case cultivationarea.AreaTypeBlock:
		if !parentID.Valid {
			return nil, fmt.Errorf("block requires parent_id")
		}
		block := cultivationarea.NewBlock(parentID.String, name, geom)
		block.Rehydrate(id, farmRefID, createdAt, updatedAt, seasons)
		return block, nil

	case cultivationarea.AreaTypeBed:
		if !parentID.Valid {
			return nil, fmt.Errorf("bed requires parent_id")
		}
		bed := cultivationarea.NewBed(id, parentID.String, name, geom, areaValue)
		//bed.Rehydrate(id, farmRefID, createdAt, updatedAt, seasons)
		return bed, nil

	case cultivationarea.AreaTypeGreenhouse:
		//TODO сделать атрибуты для хранения размеров
		greenhouse := cultivationarea.NewGreenhouseArea(farmRefID, name, *types.NewDimension(12, 4), geom)
		greenhouse.Rehydrate(id, createdAt, updatedAt, seasons)
		return greenhouse, nil

	default:
		return nil, fmt.Errorf("%w: %s", cultivationarea.ErrUnknownAreaType, areaType)
	}
}
