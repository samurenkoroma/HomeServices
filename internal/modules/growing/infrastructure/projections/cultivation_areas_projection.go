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

func (p *CultivationAreaProjection) GetByID(ctx context.Context, id string) (*cultivationarea.Detail, error) {
	query := `
SELECT 
    id, name, farm_ref_id, type, ST_AsGeoJSON(geometry),area,parent_id,created_at 
FROM public.growing_cultivation_areas 
WHERE id=$1`

	rows := p.db.QueryRowContext(ctx, query, id)
	var item cultivationarea.Detail
	if err := rows.Scan(&item.Id, &item.Name, &item.FarmRefId, &item.Type, &item.Geometry, &item.Area, &item.ParentId, &item.CreatedAt); err != nil {
		return nil, fmt.Errorf("failed to scan growing_cultivation_areas: %w", err)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate seasons: %w", err)
	}
	return &item, nil
}

func NewCultivationAreaProjections(db *sql.DB) cultivationarea.Projections {
	return &CultivationAreaProjection{db: db}
}
