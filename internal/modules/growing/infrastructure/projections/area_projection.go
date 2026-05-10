package projections

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
)

type areaProjection struct {
	db *sql.DB
}

func NewAreaProjection(db *sql.DB) cultivationarea.Projections {
	return &areaProjection{db: db}
}

// GetList возвращает список мест выращивания
func (p *areaProjection) GetList(ctx context.Context, filter cultivationarea.AreaFilter) ([]cultivationarea.AreaListItem, error) {
	var query string

	query = `
            SELECT 
                ca.id, 
                ca.farm_ref_id, 
                ca.type, 
                ca.name, 
                ca.area
            FROM public.growing_cultivation_areas ca
            WHERE ($1 = '' OR ca.farm_ref_id::text = $1)
            ORDER BY ca.name
        `

	rows, err := p.db.QueryContext(ctx, query, filter.ObjectId)
	if err != nil {
		return nil, fmt.Errorf("failed to query areas: %w", err)
	}
	defer rows.Close()

	var items []cultivationarea.AreaListItem
	for rows.Next() {
		var item cultivationarea.AreaListItem

		err := rows.Scan(
			&item.ID,
			&item.ObjectId,
			&item.Type,
			&item.Name,
			&item.Area,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan area: %w", err)
		}

		items = append(items, item)
	}

	return items, nil
}

// GetByID возвращает детальную информацию о месте выращивания
func (p *areaProjection) GetByID(ctx context.Context, id string) (*cultivationarea.AreaDetail, error) {
	if id == "" {
		return nil, cultivationarea.ErrAreaNotFound
	}

	query := `
        SELECT 
            ca.id,
            ca.farm_ref_id,
            ca.type,
            ca.name,
            ST_AsGeoJSON(ca.geometry) as geometry,
            ca.area,
            ca.created_at,
            ca.updated_at,
            -- Атрибуты (из farm_physical_objects)
            fo.attributes as farm_attributes,
            -- Текущая конфигурация
            cfg.season_id as current_season_id,
            s.name as current_season_name,
            cfg.crop_plan_id,
            cp.name as crop_plan_name,
            CASE 
                WHEN cfg.crop_plan_id IS NOT NULL THEN 'monoculture'
                ELSE 'polyculture'
            END as usage_type,
            -- Статистика
            COUNT(CASE WHEN cc.status IN ('active', 'growing') THEN 1 END) as active_cycles_count,
            COUNT(CASE WHEN cc.status = 'completed' THEN 1 END) as completed_cycles_count,
            COALESCE(SUM(cc.yield_actual), 0) as total_yield
        FROM public.growing_cultivation_areas ca
        LEFT JOIN farm_physical_objects fo ON fo.id = ca.farm_ref_id
        LEFT JOIN public.growing_area_season_configs cfg ON cfg.area_id = ca.id
        LEFT JOIN public.growing_seasons s ON s.id = cfg.season_id AND s.status = 'active'
        LEFT JOIN public.crop_crop_plans cp ON cp.id = cfg.crop_plan_id
        LEFT JOIN public.growing_crop_cycles cc ON cc.area_id = ca.id
        WHERE ca.id::text = $1
        GROUP BY 
            ca.id, ca.farm_ref_id, ca.type, ca.name, ca.geometry, ca.area,
            ca.parent_id, ca.created_at, ca.updated_at,
            fo.attributes,
            cfg.season_id, s.name, cfg.crop_plan_id, cp.name
    `

	var detail cultivationarea.AreaDetail
	var geometryJSON string
	var parentID sql.NullString
	var farmAttributesJSON []byte
	var currentSeasonID sql.NullString
	var currentSeasonName sql.NullString
	var cropPlanID sql.NullString
	var cropPlanName sql.NullString
	var usageType sql.NullString

	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&detail.ID,
		&detail.FarmRefID,
		&detail.Type,
		&detail.Name,
		&geometryJSON,
		&detail.Area,
		&parentID,
		&detail.CreatedAt,
		&detail.UpdatedAt,
		&farmAttributesJSON,
		&currentSeasonID,
		&currentSeasonName,
		&cropPlanID,
		&cropPlanName,
		&usageType,
		&detail.ActiveCyclesCount,
		&detail.CompletedCyclesCount,
		&detail.TotalYield,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, cultivationarea.ErrAreaNotFound
		}
		return nil, fmt.Errorf("failed to get area detail: %w", err)
	}

	// Парсим геометрию
	var geometry interface{}
	if err := json.Unmarshal([]byte(geometryJSON), &geometry); err != nil {
		return nil, fmt.Errorf("failed to parse geometry: %w", err)
	}
	detail.Geometry = geometry

	// Парсим атрибуты из farm_physical_objects
	if len(farmAttributesJSON) > 0 {
		var farmAttrs map[string]interface{}
		if err := json.Unmarshal(farmAttributesJSON, &farmAttrs); err != nil {
			return nil, fmt.Errorf("failed to parse farm attributes: %w", err)
		}

		detail.Attributes = &cultivationarea.AreaAttributes{}

		// Для теплицы
		if detail.Type == "greenhouse" {
			if v, ok := farmAttrs["greenhouse_type"].(string); ok {
				detail.Attributes.GreenhouseType = v
			}
			if v, ok := farmAttrs["width"].(float64); ok {
				detail.Attributes.Width = v
			}
			if v, ok := farmAttrs["length"].(float64); ok {
				detail.Attributes.Length = v
			}
			if v, ok := farmAttrs["height"].(float64); ok {
				detail.Attributes.Height = v
			}
			if v, ok := farmAttrs["has_heating"].(bool); ok {
				detail.Attributes.HasHeating = v
			}
			if v, ok := farmAttrs["has_ventilation"].(bool); ok {
				detail.Attributes.HasVentilation = v
			}
			if v, ok := farmAttrs["has_lighting"].(bool); ok {
				detail.Attributes.HasLighting = v
			}
		}

		// Для поля
		if detail.Type == "field" {
			if v, ok := farmAttrs["soil_type"].(string); ok {
				detail.Attributes.SoilType = v
			}
		}

		if v, ok := farmAttrs["description"].(string); ok {
			detail.Attributes.Description = v
		}
	}

	// Опциональные поля
	if parentID.Valid {
		detail.ParentID = &parentID.String
	}
	if currentSeasonID.Valid {
		detail.CurrentSeasonID = &currentSeasonID.String
	}
	if currentSeasonName.Valid {
		detail.CurrentSeasonName = &currentSeasonName.String
	}
	if cropPlanID.Valid {
		detail.CropPlanID = &cropPlanID.String
	}
	if cropPlanName.Valid {
		detail.CropPlanName = &cropPlanName.String
	}
	if usageType.Valid {
		detail.UsageType = &usageType.String
	}

	return &detail, nil
}
