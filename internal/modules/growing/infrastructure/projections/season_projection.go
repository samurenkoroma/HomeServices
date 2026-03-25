package projections

import (
	"context"
	"database/sql"
	"time"
)

// SeasonDetailDTO — детальная информация о сезоне
type SeasonDetailDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Status      string    `json:"status"`
	Description string    `json:"description"`

	// Статистика
	AreasCount      int     `json:"areas_count"`
	ConfiguredAreas int     `json:"configured_areas"`
	ActiveCycles    int     `json:"active_cycles"`
	TotalArea       float64 `json:"total_area"`
}

// SeasonProjection — read-модель для сезонов
type SeasonProjection struct {
	db *sql.DB
}

func NewSeasonProjection(db *sql.DB) *SeasonProjection {
	return &SeasonProjection{db: db}
}

// GetSeasonWithStats — получить сезон со статистикой
func (p *SeasonProjection) GetSeasonWithStats(ctx context.Context, id string) (*SeasonDetailDTO, error) {
	query := `
        SELECT 
            s.id, s.name, s.start_date, s.end_date, s.status, s.description,
            COUNT(DISTINCT ca.id) as areas_count,
            COUNT(DISTINCT CASE WHEN asc.area_id IS NOT NULL THEN ca.id END) as configured_areas,
            COUNT(DISTINCT CASE WHEN cc.status IN ('active', 'growing') THEN cc.id END) as active_cycles,
            COALESCE(SUM(ca.area), 0) as total_area
        FROM seasons s
        LEFT JOIN area_season_configs asc ON asc.season_id = s.id
        LEFT JOIN cultivation_areas ca ON ca.id = asc.area_id
        LEFT JOIN crop_cycles cc ON cc.season_id = s.id
        WHERE s.id = $1
        GROUP BY s.id
    `

	var dto SeasonDetailDTO
	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&dto.ID, &dto.Name, &dto.StartDate, &dto.EndDate, &dto.Status, &dto.Description,
		&dto.AreasCount, &dto.ConfiguredAreas, &dto.ActiveCycles, &dto.TotalArea,
	)

	return &dto, err
}
