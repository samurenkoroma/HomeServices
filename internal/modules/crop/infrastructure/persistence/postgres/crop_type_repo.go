package postgres

import (
	"context"
	"database/sql"
	"fmt"
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
        INSERT INTO crop_types (
            id, name, scientific_name, category, description,
            root_depth, is_perennial, vegetation_days,
            default_yield, market_price, is_active,
            created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name,
            scientific_name = EXCLUDED.scientific_name,
            category = EXCLUDED.category,
            description = EXCLUDED.description,
            root_depth = EXCLUDED.root_depth,
            is_perennial = EXCLUDED.is_perennial,
            vegetation_days = EXCLUDED.vegetation_days,
            default_yield = EXCLUDED.default_yield,
            market_price = EXCLUDED.market_price,
            is_active = EXCLUDED.is_active,
            updated_at = EXCLUDED.updated_at
    `

	_, err := r.tx.ExecContext(ctx, query,
		ct.GetID(),
		ct.GetName(),
		ct.GetScientificName(),
		string(ct.GetCategory()),
		ct.GetDescription(),
		ct.GetRootDepth(),
		ct.IsPerennial(),
		ct.GetVegetationDays(),
		ct.GetDefaultYield(),
		ct.GetMarketPrice(),
		ct.IsActive(),
		ct.GetCreatedAt(),
		ct.GetUpdatedAt(),
	)

	if err != nil {
		return fmt.Errorf("failed to save crop type: %w", err)
	}

	return nil
}

// FindByID находит тип культуры по ID
func (r *cropTypeRepository) FindByID(ctx context.Context, id croptype.CropTypeID) (*croptype.CropType, error) {
	query := `
        SELECT 
            id, name, scientific_name, category, description,
            root_depth, is_perennial, vegetation_days,
            default_yield, market_price, is_active,
            created_at, updated_at
        FROM crop_types
        WHERE id = $1
    `

	var (
		ctID           string
		name           string
		scientificName sql.NullString
		category       string
		description    sql.NullString
		rootDepth      sql.NullInt64
		isPerennial    bool
		vegetationDays int
		defaultYield   sql.NullFloat64
		marketPrice    sql.NullFloat64
		isActive       bool
		createdAt      time.Time
		updatedAt      time.Time
	)

	err := r.tx.QueryRowContext(ctx, query, string(id)).Scan(
		&ctID, &name, &scientificName, &category, &description,
		&rootDepth, &isPerennial, &vegetationDays,
		&defaultYield, &marketPrice, &isActive,
		&createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, croptype.ErrCropTypeNotFound
		}
		return nil, fmt.Errorf("failed to find crop type: %w", err)
	}

	// Создаем тип культуры через конструктор
	ct, err := croptype.NewCropType(
		name,
		scientificName.String,
		croptype.CropCategory(category),
		vegetationDays,
	)
	if err != nil {
		return nil, err
	}

	// Восстанавливаем остальные поля
	ct.Rehydrate(
		croptype.CropTypeID(ctID),
		description.String,
		int(rootDepth.Int64),
		isPerennial,
		defaultYield.Float64,
		marketPrice.Float64,
		isActive,
		createdAt,
		updatedAt,
	)

	return ct, nil
}

// FindByName находит тип культуры по имени
func (r *cropTypeRepository) FindByName(ctx context.Context, name string) (*croptype.CropType, error) {
	query := `
        SELECT 
            id, name, scientific_name, category, description,
            root_depth, is_perennial, vegetation_days,
            default_yield, market_price, is_active,
            created_at, updated_at
        FROM crop_types
        WHERE name = $1
    `

	var (
		ctID           string
		ctName         string
		scientificName sql.NullString
		category       string
		description    sql.NullString
		rootDepth      sql.NullInt64
		isPerennial    bool
		vegetationDays int
		defaultYield   sql.NullFloat64
		marketPrice    sql.NullFloat64
		isActive       bool
		createdAt      time.Time
		updatedAt      time.Time
	)

	err := r.tx.QueryRowContext(ctx, query, name).Scan(
		&ctID, &ctName, &scientificName, &category, &description,
		&rootDepth, &isPerennial, &vegetationDays,
		&defaultYield, &marketPrice, &isActive,
		&createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, croptype.ErrCropTypeNotFound
		}
		return nil, fmt.Errorf("failed to find crop type by name: %w", err)
	}

	ct, err := croptype.NewCropType(
		ctName,
		scientificName.String,
		croptype.CropCategory(category),
		vegetationDays,
	)
	if err != nil {
		return nil, err
	}

	ct.Rehydrate(
		croptype.CropTypeID(ctID),
		description.String,
		int(rootDepth.Int64),
		isPerennial,
		defaultYield.Float64,
		marketPrice.Float64,
		isActive,
		createdAt,
		updatedAt,
	)

	return ct, nil
}

// FindAll возвращает все типы культур
func (r *cropTypeRepository) FindAll(ctx context.Context) ([]*croptype.CropType, error) {
	query := `
        SELECT 
            id, name, scientific_name, category, description,
            root_depth, is_perennial, vegetation_days,
            default_yield, market_price, is_active,
            created_at, updated_at
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
			ctID           string
			name           string
			scientificName sql.NullString
			category       string
			description    sql.NullString
			rootDepth      sql.NullInt64
			isPerennial    bool
			vegetationDays int
			defaultYield   sql.NullFloat64
			marketPrice    sql.NullFloat64
			isActive       bool
			createdAt      time.Time
			updatedAt      time.Time
		)

		err := rows.Scan(
			&ctID, &name, &scientificName, &category, &description,
			&rootDepth, &isPerennial, &vegetationDays,
			&defaultYield, &marketPrice, &isActive,
			&createdAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan crop type: %w", err)
		}

		ct, err := croptype.NewCropType(
			name,
			scientificName.String,
			croptype.CropCategory(category),
			vegetationDays,
		)
		if err != nil {
			return nil, err
		}

		ct.Rehydrate(
			croptype.CropTypeID(ctID),
			description.String,
			int(rootDepth.Int64),
			isPerennial,
			defaultYield.Float64,
			marketPrice.Float64,
			isActive,
			createdAt,
			updatedAt,
		)

		cropTypes = append(cropTypes, ct)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return cropTypes, nil
}

// FindByCategory возвращает типы культур по категории
func (r *cropTypeRepository) FindByCategory(ctx context.Context, category croptype.CropCategory) ([]*croptype.CropType, error) {
	query := `
        SELECT 
            id, name, scientific_name, category, description,
            root_depth, is_perennial, vegetation_days,
            default_yield, market_price, is_active,
            created_at, updated_at
        FROM crop_types
        WHERE category = $1
        ORDER BY name
    `

	rows, err := r.tx.QueryContext(ctx, query, string(category))
	if err != nil {
		return nil, fmt.Errorf("failed to query crop types by category: %w", err)
	}
	defer rows.Close()

	var cropTypes []*croptype.CropType

	for rows.Next() {
		var (
			ctID           string
			name           string
			scientificName sql.NullString
			cat            string
			description    sql.NullString
			rootDepth      sql.NullInt64
			isPerennial    bool
			vegetationDays int
			defaultYield   sql.NullFloat64
			marketPrice    sql.NullFloat64
			isActive       bool
			createdAt      time.Time
			updatedAt      time.Time
		)

		err := rows.Scan(
			&ctID, &name, &scientificName, &cat, &description,
			&rootDepth, &isPerennial, &vegetationDays,
			&defaultYield, &marketPrice, &isActive,
			&createdAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan crop type: %w", err)
		}

		ct, err := croptype.NewCropType(
			name,
			scientificName.String,
			croptype.CropCategory(cat),
			vegetationDays,
		)
		if err != nil {
			return nil, err
		}

		ct.Rehydrate(
			croptype.CropTypeID(ctID),
			description.String,
			int(rootDepth.Int64),
			isPerennial,
			defaultYield.Float64,
			marketPrice.Float64,
			isActive,
			createdAt,
			updatedAt,
		)

		cropTypes = append(cropTypes, ct)
	}

	return cropTypes, nil
}

// FindActive возвращает все активные типы культур
func (r *cropTypeRepository) FindActive(ctx context.Context) ([]*croptype.CropType, error) {
	query := `
        SELECT 
            id, name, scientific_name, category, description,
            root_depth, is_perennial, vegetation_days,
            default_yield, market_price, is_active,
            created_at, updated_at
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
			ctID           string
			name           string
			scientificName sql.NullString
			category       string
			description    sql.NullString
			rootDepth      sql.NullInt64
			isPerennial    bool
			vegetationDays int
			defaultYield   sql.NullFloat64
			marketPrice    sql.NullFloat64
			isActive       bool
			createdAt      time.Time
			updatedAt      time.Time
		)

		err := rows.Scan(
			&ctID, &name, &scientificName, &category, &description,
			&rootDepth, &isPerennial, &vegetationDays,
			&defaultYield, &marketPrice, &isActive,
			&createdAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan crop type: %w", err)
		}

		ct, err := croptype.NewCropType(
			name,
			scientificName.String,
			croptype.CropCategory(category),
			vegetationDays,
		)
		if err != nil {
			return nil, err
		}

		ct.Rehydrate(
			croptype.CropTypeID(ctID),
			description.String,
			int(rootDepth.Int64),
			isPerennial,
			defaultYield.Float64,
			marketPrice.Float64,
			isActive,
			createdAt,
			updatedAt,
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
