package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"samurenkoroma/services/internal/modules/growing/domain/season"
	"time"
)

type seasonRepo struct {
	tx *sql.Tx
}

func NewSeasonRepository(tx *sql.Tx) season.Repository {
	return &seasonRepo{
		tx: tx,
	}
}

// Save сохраняет или обновляет сезон
func (r *seasonRepo) Save(ctx context.Context, s *season.Season) error {
	query := `
        INSERT INTO growing_seasons (
            id, name, start_date, end_date, description, status, created_by, created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name,
            start_date = EXCLUDED.start_date,
            end_date = EXCLUDED.end_date,
            description = EXCLUDED.description,
            status = EXCLUDED.status,
            updated_at = EXCLUDED.updated_at
    `

	_, err := r.tx.ExecContext(ctx, query,
		string(s.GetId()),
		s.GetName(),
		s.GetStartDate(),
		s.GetEndDate(),
		s.GetDescription(),
		string(s.GetStatus()),
		s.GetCreatedBy(),
		s.GetCreatedAt(),
		s.GetUpdatedAt(),
	)

	if err != nil {
		return fmt.Errorf("failed to save season: %w", err)
	}

	return nil
}

func (r *seasonRepo) FindByID(ctx context.Context, id season.SeasonID) (*season.Season, error) {
	query := `
        SELECT id, name, start_date, end_date, description, status, created_by, created_at, updated_at
        FROM growing_seasons
        WHERE id = $1
    `

	var (
		sid         string
		name        string
		startDate   time.Time
		endDate     time.Time
		description sql.NullString
		status      string
		createdBy   string
		createdAt   time.Time
		updatedAt   time.Time
	)

	err := r.tx.QueryRowContext(ctx, query, string(id)).Scan(
		&sid, &name, &startDate, &endDate, &description, &status,
		&createdBy, &createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, season.ErrSeasonNotFound
		}
		return nil, fmt.Errorf("failed to find season: %w", err)
	}

	s := season.Rehydrate(
		season.SeasonID(sid),
		season.SeasonStatus(status),
		createdBy,
		name,
		description.String,
		startDate,
		endDate,
		createdAt,
		updatedAt,
	)

	return s, nil
}
