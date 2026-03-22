package postgres

import (
	"context"
	"database/sql"
	"fmt"
	sqlErrors "samurenkoroma/services/internal/core/errors"
	"time"

	"samurenkoroma/services/internal/modules/crop/domain/croptype"
)

type cropTypeRepository struct {
	tx *sql.Tx
}

// NewCropTypeRepository создает новый репозиторий типов культур
func NewCropTypeRepository(tx *sql.Tx) croptype.Repository {
	return &cropTypeRepository{tx: tx}
}

// Save сохраняет или обновляет тип культуры
func (r *cropTypeRepository) Save(ctx context.Context, ct *croptype.CropType) error {
	query := `
        INSERT INTO crop_types (id, name, category, description, is_perennial, is_active, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name,
            category = EXCLUDED.category,
            description = EXCLUDED.description,
            is_perennial = EXCLUDED.is_perennial,
            updated_at = EXCLUDED.updated_at
    `

	_, err := r.tx.ExecContext(ctx, query,
		ct.GetID(),
		ct.GetName(),
		string(ct.GetCategory()),
		ct.GetDescription(),
		ct.IsPerennial(),
		ct.IsActive(),
		ct.GetCreatedAt(),
		ct.GetUpdatedAt(),
	)

	if err != nil {
		if sqlErrors.IsUniqueViolation(err) {
			return croptype.ErrCropTypeAlreadyExists
		}
		return fmt.Errorf("failed to save crop type: %w", err)
	}

	return nil
}

// FindByID находит тип культуры по ID
func (r *cropTypeRepository) FindByID(ctx context.Context, search croptype.CropTypeID) (*croptype.CropType, error) {
	query := `
        SELECT id, name,  category, description, is_perennial, is_active, created_at, updated_at 
        FROM crop_types
        WHERE id = $1
    `

	var (
		id          string
		name        string
		category    string
		description sql.NullString
		isPerennial bool
		isActive    bool
		createdAt   time.Time
		updatedAt   time.Time
	)

	err := r.tx.QueryRowContext(ctx, query, string(search)).Scan(
		&id, &name, &category, &description, &isPerennial, &isActive, &createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, croptype.ErrCropTypeNotFound
		}
		return nil, fmt.Errorf("failed to find crop type: %w", err)
	}

	ct := croptype.Rehydrate(
		id,
		name,
		category,
		description.String,
		isPerennial,
		isActive,
	)

	return ct, nil
}

// FindByName находит тип культуры по имени
func (r *cropTypeRepository) FindByName(ctx context.Context, search string) (*croptype.CropType, error) {
	query := `
        SELECT id, name,  category, description, is_perennial, is_active, created_at, updated_at 
        FROM crop_types
        WHERE name = $1
    `

	var (
		id          string
		name        string
		category    string
		description sql.NullString
		isPerennial bool
		isActive    bool
		createdAt   time.Time
		updatedAt   time.Time
	)

	err := r.tx.QueryRowContext(ctx, query, search).Scan(
		&id, &name, &category, &description, &isPerennial, &isActive, &createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, croptype.ErrCropTypeNotFound
		}
		return nil, fmt.Errorf("failed to find crop type: %w", err)
	}

	ct := croptype.Rehydrate(
		id,
		name,
		category,
		description.String,
		isPerennial,
		isActive,
	)

	return ct, nil
}

// FindAll возвращает все типы культур
func (r *cropTypeRepository) FindAll(ctx context.Context) ([]*croptype.CropType, error) {
	query := `
        SELECT id, name,  category, description, is_perennial, is_active, created_at, updated_at 
        FROM crop_types
        ORDER BY name
    `

	rows, err := r.tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query crop types: %w", err)
	}
	defer rows.Close()

	var cropTypes []*croptype.CropType

	for rows.Next() {
		var (
			id          string
			name        string
			category    string
			description sql.NullString
			isPerennial bool
			isActive    bool
			createdAt   time.Time
			updatedAt   time.Time
		)

		err := rows.Scan(
			&id, &name, &category, &description,
			&isPerennial, &isActive, &createdAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan crop type: %w", err)
		}

		ct := croptype.Rehydrate(
			id,
			name,
			category,
			description.String,
			isPerennial,
			isActive,
		)

		cropTypes = append(cropTypes, ct)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return cropTypes, nil
}

// FindByCategory возвращает типы культур по категории
func (r *cropTypeRepository) FindByCategory(ctx context.Context, search croptype.CropCategory) ([]*croptype.CropType, error) {
	query := `
	SELECT  id, name, category,  description, is_perennial,is_active, created_at, updated_at 
        FROM crop_types
        WHERE category = $1
        ORDER BY name
    `

	rows, err := r.tx.QueryContext(ctx, query, string(search))
	if err != nil {
		return nil, fmt.Errorf("failed to query crop types by category: %w", err)
	}
	defer rows.Close()

	var cropTypes []*croptype.CropType

	for rows.Next() {
		var (
			ctID        string
			name        string
			category    string
			description sql.NullString
			isPerennial bool
			isActive    bool
			createdAt   time.Time
			updatedAt   time.Time
		)

		err := rows.Scan(
			&ctID, &name, &category, &description,
			&isPerennial, &isActive, &createdAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan crop type: %w", err)
		}

		ct := croptype.Rehydrate(
			ctID,
			name,
			category,
			description.String,
			isPerennial, isActive,
		)
		cropTypes = append(cropTypes, ct)
	}

	return cropTypes, nil
}

// FindActive возвращает все активные типы культур
func (r *cropTypeRepository) FindActive(ctx context.Context) ([]*croptype.CropType, error) {
	query := `
        SELECT id, name, category,  description, is_perennial,is_active, created_at, updated_at 
        FROM crop_types
        WHERE is_active = true
        ORDER BY name
    `

	rows, err := r.tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query active crop types: %w", err)
	}
	defer rows.Close()

	var cropTypes []*croptype.CropType

	for rows.Next() {
		var (
			ctID        string
			name        string
			category    string
			description sql.NullString
			isPerennial bool
			isActive    bool
			createdAt   time.Time
			updatedAt   time.Time
		)

		err := rows.Scan(
			&ctID, &name, &category, &description,
			&isPerennial, &isActive, &createdAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan crop type: %w", err)
		}

		ct := croptype.Rehydrate(
			ctID,
			name,
			category,
			description.String,
			isPerennial, isActive,
		)

		cropTypes = append(cropTypes, ct)
	}

	return cropTypes, nil
}

// Exists проверяет существование типа культуры по имени
func (r *cropTypeRepository) Exists(ctx context.Context, name string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM crop_types WHERE name = $1)`

	var exists bool
	err := r.tx.QueryRowContext(ctx, query, name).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check existence: %w", err)
	}

	return exists, nil
}

// Delete мягко удаляет тип культуры (архивирует)
func (r *cropTypeRepository) Delete(ctx context.Context, id croptype.CropTypeID) error {
	// Проверяем, не используется ли тип в планах или сортах
	var plansCount, varietiesCount int

	err := r.tx.QueryRowContext(ctx, `
        SELECT COUNT(*) FROM crop_plans WHERE crop_type_id = $1
    `, string(id)).Scan(&plansCount)
	if err != nil {
		return fmt.Errorf("failed to check plans: %w", err)
	}

	err = r.tx.QueryRowContext(ctx, `
        SELECT COUNT(*) FROM varieties WHERE crop_type_id = $1
    `, string(id)).Scan(&varietiesCount)
	if err != nil {
		return fmt.Errorf("failed to check varieties: %w", err)
	}

	if plansCount > 0 || varietiesCount > 0 {
		return croptype.ErrCropTypeInUse
	}

	// Мягкое удаление (архивация)
	query := `UPDATE crop_types SET is_active = false, updated_at = $1 WHERE id = $2`

	_, err = r.tx.ExecContext(ctx, query, time.Now(), string(id))
	if err != nil {
		return fmt.Errorf("failed to delete crop type: %w", err)
	}

	return nil
}
