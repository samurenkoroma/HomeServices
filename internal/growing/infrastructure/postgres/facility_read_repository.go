package postgres

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/growing/application"
)

type facilityReadRepository struct {
	db *sql.DB
}

func NewFacilityReadRepository(
	db *sql.DB,
) application.FacilityReadRepository {
	return &facilityReadRepository{
		db: db,
	}
}

func (r *facilityReadRepository) GetOverview(
	ctx context.Context,
	id string,
) (*application.FacilityOverviewDTO, error) {

	const query = `
	SELECT
		id,
		name,
		unit_type,
		length,
    	width
	FROM land_structure
	WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	dto := &application.FacilityOverviewDTO{}

	if err := row.Scan(
		&dto.ID,
		&dto.Name,
		&dto.Type,
		&dto.Length,
		&dto.Width,
	); err != nil {
		return nil, err
	}

	return dto, nil
}
