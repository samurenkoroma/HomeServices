package projections

import (
	"context"
	"database/sql"
	"fmt"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
)

type CultivationAreaProjection struct {
	db *sql.DB
}

func (p *CultivationAreaProjection) GetList(ctx context.Context, filter cultivationarea.Filter) ([]*cultivationarea.ListItem, error) {
	query := `
SELECT 
    id, name, farm_ref_id, type, ST_AsGeoJSON(geometry),area,parent_id,created_at 
FROM public.growing_cultivation_areas 
LIMIT $1 OFFSET $2`

	rows, err := p.db.QueryContext(ctx, query, filter.Limit, filter.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query growing_cultivation_areas: %w", err)
	}
	defer rows.Close()
	var items []*cultivationarea.ListItem
	for rows.Next() {
		var item cultivationarea.ListItem
		if err := rows.Scan(&item.Id, &item.Name, &item.FarmRefId, &item.Type, &item.Geometry, &item.Area, &item.ParentId, &item.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan growing_cultivation_areas: %w", err)
		}
		items = append(items, &item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate seasons: %w", err)
	}
	return items, nil
}

//func (p *CultivationAreaProjection) GetByID(ctx context.Context, id string) (*season.Detail, error) {
//	query := `
//        SELECT
//            s.id, s.name, s.start_date, s.end_date, s.status, s.description,
//            COUNT(DISTINCT ca.id) as areas_count,
//            COUNT(DISTINCT CASE WHEN cnf.area_id IS NOT NULL THEN ca.id END) as configured_areas,
//            COUNT(DISTINCT CASE WHEN cc.status IN ('active', 'growing') THEN cc.id END) as active_cycles,
//            COALESCE(SUM(ca.area), 0) as total_area
//        FROM public.growing_seasons s
//        LEFT JOIN public.growing_area_season_configs cnf ON cnf.season_id = s.id
//        LEFT JOIN growing_cultivation_areas ca ON ca.id = cnf.area_id
//        LEFT JOIN growing_crop_cycles cc ON cc.season_id = s.id
//        WHERE s.id = $1
//        GROUP BY s.id
//    `
//
//	var dto season.Detail
//	err := p.db.QueryRowContext(ctx, query, id).Scan(
//		&dto.ID, &dto.Name, &dto.StartDate, &dto.EndDate, &dto.Status, &dto.Description,
//		&dto.AreasCount, &dto.ConfiguredAreas, &dto.ActiveCycles, &dto.TotalArea,
//	)
//
//	return &dto, err
//}

func NewCultivationAreaProjections(db *sql.DB) cultivationarea.Projections {
	return &CultivationAreaProjection{db: db}
}
