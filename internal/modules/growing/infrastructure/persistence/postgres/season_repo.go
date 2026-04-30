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

// FindAll возвращает все сезоны
func (r *seasonRepo) FindAll(ctx context.Context, filter season.Filter) ([]*season.Season, error) {
	query := `
        SELECT id, name, start_date, end_date, description, status, created_by, created_at, updated_at
        FROM growing_seasons s
        WHERE ($1 = '' OR s.created_by::text = $1)
        ORDER BY start_date DESC
    `

	rows, err := r.tx.QueryContext(ctx, query, filter.OwnerId)
	if err != nil {
		return nil, fmt.Errorf("failed to query seasons: %w", err)
	}
	defer rows.Close()

	var seasons []*season.Season

	for rows.Next() {
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

		err := rows.Scan(
			&sid, &name, &startDate, &endDate, &description, &status,
			&createdBy, &createdAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan season: %w", err)
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

		seasons = append(seasons, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return seasons, nil
}

// Delete удаляет сезон (мягкое удаление — архивирование)
func (r *seasonRepo) Delete(ctx context.Context, id season.SeasonID) error {
	// Проверяем, есть ли активные циклы в этом сезоне
	var cycleCount int
	err := r.tx.QueryRowContext(ctx, `
        SELECT COUNT(*) FROM crop_cycles WHERE season_id = $1 AND status != 'completed'
    `, string(id)).Scan(&cycleCount)
	if err != nil {
		return fmt.Errorf("failed to check cycles: %w", err)
	}

	if cycleCount > 0 {
		return fmt.Errorf("cannot delete season with active cycles")
	}

	// Мягкое удаление — архивируем
	_, err = r.tx.ExecContext(ctx, `
        UPDATE growing_seasons SET status = 'archived', updated_at = $1 WHERE id = $2
    `, time.Now(), string(id))

	if err != nil {
		return fmt.Errorf("failed to archive season: %w", err)
	}

	return nil
}

// FindByName находит сезон по имени
func (r *seasonRepo) FindByName(ctx context.Context, name string) (*season.Season, error) {
	query := `
        SELECT id, name, start_date, end_date, description, status, created_by, created_at, updated_at
        FROM growing_seasons
        WHERE name = $1
    `

	var (
		sid         string
		sname       string
		startDate   time.Time
		endDate     time.Time
		description sql.NullString
		status      string
		createdBy   string
		createdAt   time.Time
		updatedAt   time.Time
	)

	err := r.tx.QueryRowContext(ctx, query, name).Scan(
		&sid, &sname, &startDate, &endDate, &description, &status,
		&createdBy, &createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, season.ErrSeasonNotFound
		}
		return nil, fmt.Errorf("failed to find season by name: %w", err)
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

// FindByStatus возвращает сезоны по статусу
func (r *seasonRepo) FindByStatus(ctx context.Context, status season.SeasonStatus) ([]*season.Season, error) {
	query := `
        SELECT id, name, start_date, end_date, description, status, created_by, created_at, updated_at
        FROM growing_seasons
        WHERE status = $1
        ORDER BY start_date DESC
    `

	rows, err := r.tx.QueryContext(ctx, query, string(status))
	if err != nil {
		return nil, fmt.Errorf("failed to query seasons by status: %w", err)
	}
	defer rows.Close()

	var seasons []*season.Season

	for rows.Next() {
		var (
			sid         string
			name        string
			startDate   time.Time
			endDate     time.Time
			description sql.NullString
			st          string
			createdBy   string
			createdAt   time.Time
			updatedAt   time.Time
		)

		err := rows.Scan(
			&sid, &name, &startDate, &endDate, &description, &st,
			&createdBy, &createdAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan season: %w", err)
		}

		s := season.Rehydrate(
			season.SeasonID(sid),
			status,
			createdBy,
			name,
			description.String,
			startDate,
			endDate,
			createdAt,
			updatedAt,
		)

		seasons = append(seasons, s)
	}

	return seasons, nil
}

// FindActive возвращает текущий активный сезон
func (r *seasonRepo) FindActive(ctx context.Context) (*season.Season, error) {
	now := time.Now()

	query := `
        SELECT id, name, start_date, end_date, description, status, created_by, created_at, updated_at
        FROM growing_seasons
        WHERE status = 'active' AND start_date <= $1 AND end_date >= $1
        LIMIT 1
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

	err := r.tx.QueryRowContext(ctx, query, now).Scan(
		&sid, &name, &startDate, &endDate, &description, &status,
		&createdBy, &createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find active season: %w", err)
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

// FindOverlapping возвращает сезоны, пересекающиеся с указанным периодом
func (r *seasonRepo) FindOverlapping(ctx context.Context, start, end time.Time) ([]*season.Season, error) {
	query := `
        SELECT id, name, start_date, end_date, description, status, created_by, created_at, updated_at
        FROM growing_seasons
        WHERE (start_date, end_date) OVERLAPS ($1, $2)
        ORDER BY start_date
    `

	rows, err := r.tx.QueryContext(ctx, query, start, end)
	if err != nil {
		return nil, fmt.Errorf("failed to query overlapping seasons: %w", err)
	}
	defer rows.Close()

	var seasons []*season.Season

	for rows.Next() {
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

		err := rows.Scan(
			&sid, &name, &startDate, &endDate, &description, &status,
			&createdBy, &createdAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan season: %w", err)
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

		seasons = append(seasons, s)
	}

	return seasons, nil
}

// ExistsByName проверяет существование сезона с указанным именем
func (r *seasonRepo) ExistsByName(ctx context.Context, name string) (bool, error) {
	var exists bool
	err := r.tx.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM seasons WHERE name = $1)`, name).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check existence: %w", err)
	}
	return exists, nil
}

// ExistsOverlapping проверяет наличие пересекающихся сезонов
func (r *seasonRepo) ExistsOverlapping(ctx context.Context, start, end time.Time) (bool, error) {
	var exists bool
	err := r.tx.QueryRowContext(ctx, `
        SELECT EXISTS(
            SELECT 1 FROM growing_seasons 
            WHERE (start_date, end_date) OVERLAPS ($1, $2)
        )
    `, start, end).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check overlapping: %w", err)
	}
	return exists, nil
}
