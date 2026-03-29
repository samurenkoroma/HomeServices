package projections

import (
	"context"
	"database/sql"
	"fmt"
	"samurenkoroma/services/internal/modules/growing/domain/season"
)

// SeasonProjection — read-модель для сезонов
type SeasonProjection struct {
	db *sql.DB
}

func (p *SeasonProjection) GetList(ctx context.Context, filter season.Filter) ([]*season.ListItem, error) {
	query := `
SELECT 
    id, name, start_date, end_date, description,status,created_by,created_at 
FROM growing_seasons 
LIMIT $1 OFFSET $2`

	rows, err := p.db.QueryContext(ctx, query, filter.Limit, filter.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query seasons: %w", err)
	}
	defer rows.Close()
	var items []*season.ListItem
	for rows.Next() {
		var item season.ListItem
		if err := rows.Scan(&item.Id, &item.Name, &item.StartDate, &item.EndDate, &item.Description, &item.Status, &item.CreatedBy, &item.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan season: %w", err)
		}
		items = append(items, &item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate seasons: %w", err)
	}
	return items, nil
}

func (p *SeasonProjection) GetByID(ctx context.Context, id string) (*season.Detail, error) {
	query := `
        SELECT 
            s.id, s.name, s.start_date, s.end_date, s.status, s.description,
            COUNT(DISTINCT ca.id) as areas_count,
            COUNT(DISTINCT CASE WHEN cnf.area_id IS NOT NULL THEN ca.id END) as configured_areas,
            COUNT(DISTINCT CASE WHEN cc.status IN ('active', 'growing') THEN cc.id END) as active_cycles,
            COALESCE(SUM(ca.area), 0) as total_area
        FROM public.growing_seasons s
        LEFT JOIN public.growing_area_season_configs cnf ON cnf.season_id = s.id
        LEFT JOIN growing_cultivation_areas ca ON ca.id = cnf.area_id
        LEFT JOIN growing_crop_cycles cc ON cc.season_id = s.id
        WHERE s.id = $1
        GROUP BY s.id
    `

	var dto season.Detail
	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&dto.ID, &dto.Name, &dto.StartDate, &dto.EndDate, &dto.Status, &dto.Description,
		&dto.AreasCount, &dto.ConfiguredAreas, &dto.ActiveCycles, &dto.TotalArea,
	)

	return &dto, err
}

func NewSeasonProjection(db *sql.DB) season.Projections {
	return &SeasonProjection{db: db}
}
