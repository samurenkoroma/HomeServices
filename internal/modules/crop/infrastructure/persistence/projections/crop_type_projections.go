package projections

import (
	"context"
	"samurenkoroma/services/internal/modules/crop/domain/croptype"
)

import (
	"database/sql"
)

type cropTypeProjections struct {
	db *sql.DB
}

func NewCropTypeProjections(db *sql.DB) croptype.Projections {
	return &cropTypeProjections{db: db}
}

// GetCropTypeWithVarieties — получить тип культуры со всеми сортами
func (p *cropTypeProjections) GetCropTypeWithVarieties(ctx context.Context, id string) (*croptype.CropTypeWithVarietiesDTO, error) {
	// Основная информация о типе культуры
	query := `
        SELECT ct.id, ct.name, ct.category, ct.is_perennial, ct.created_at,
               ct.category as category_name,
--                COALESCE(cat.name, ct.category) as category_name,
               (SELECT COUNT(*) FROM varieties WHERE crop_type_id = ct.id AND is_active = true) as varieties_count
        FROM crop_types ct
--         LEFT JOIN crop_category_translations cat ON cat.code = ct.category AND cat.language = 'ru'
        WHERE ct.id = $1 AND ct.is_active = true
    `

	var dto croptype.CropTypeWithVarietiesDTO

	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&dto.ID, &dto.Name, &dto.Category, &dto.IsPerennial, &dto.CreatedAt,
		&dto.CategoryName, &dto.VarietiesCount,
	)
	if err != nil {
		return nil, err
	}

	// Получаем сорта
	varietiesQuery := `
        SELECT id, name, 
               45 /*vegetation_days_min || '–' || vegetation_days_max*/ as vegetation_days,
               is_active
        FROM varieties
        WHERE crop_type_id = $1
        ORDER BY name
    `

	rows, err := p.db.QueryContext(ctx, varietiesQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var varieties []croptype.VarietySimpleDTO
	for rows.Next() {
		var v croptype.VarietySimpleDTO
		if err := rows.Scan(&v.ID, &v.Name, &v.VegetationDays, &v.IsActive); err != nil {
			return nil, err
		}
		varieties = append(varieties, v)
	}

	dto.Varieties = varieties
	return &dto, nil
}

// GetAllCropTypesSimple — упрощённый список для селектов
func (p *cropTypeProjections) GetAllCropTypesSimple(ctx context.Context) ([]croptype.CropTypeSimpleDTO, error) {
	query := `
        SELECT ct.id, ct.name, ct.category, ct.category as category_name
--                COALESCE(cat.name, ct.category) as category_name
        FROM crop_types ct
--         LEFT JOIN crop_category_translations cat ON cat.code = ct.category AND cat.language = 'ru'
        WHERE ct.is_active = true
        ORDER BY ct.name
    `

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []croptype.CropTypeSimpleDTO
	for rows.Next() {
		var dto croptype.CropTypeSimpleDTO
		if err := rows.Scan(&dto.ID, &dto.Name, &dto.Category, &dto.CategoryName); err != nil {
			return nil, err
		}
		result = append(result, dto)
	}

	return result, nil
}

// GetList — получить список типов культур
func (p *cropTypeProjections) GetList(ctx context.Context, filter croptype.Filter) ([]croptype.CropTypeSimpleDTO, error) {
	query := `
        SELECT ct.id, ct.name, ct.category, ct.is_perennial,ct.category as category_name
        FROM crop_types ct
        WHERE ($1 = '' OR ct.category = $1)
          AND ($2 = false OR ct.is_active = true)
        ORDER BY ct.name
    `

	rows, err := p.db.QueryContext(ctx, query, filter.Category, filter.IsActive)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []croptype.CropTypeSimpleDTO
	for rows.Next() {
		var dto croptype.CropTypeSimpleDTO
		if err := rows.Scan(&dto.ID, &dto.Name, &dto.Category, &dto.IsPerennial, &dto.CategoryName); err != nil {
			return nil, err
		}
		result = append(result, dto)
	}

	return result, nil
}

// GetByID — получить тип культуры по ID
func (p *cropTypeProjections) GetByID(ctx context.Context, id string) (*croptype.CropTypeDetailDTO, error) {
	query := `
        SELECT ct.id, ct.name, ct.category, ct.description, 
               ct.is_perennial, ct.is_active, ct.created_at, ct.updated_at,
               ct.category as category_name,
--                COALESCE(cat.name, ct.category) as category_name,
               (SELECT COUNT(*) FROM varieties WHERE crop_type_id = ct.id AND is_active = true) as varieties_count
        FROM crop_types ct
--         LEFT JOIN crop_category_translations cat ON cat.code = ct.category AND cat.language = 'ru'
        WHERE ct.id = $1
    `

	var dto croptype.CropTypeDetailDTO
	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&dto.ID, &dto.Name, &dto.ScientificName, &dto.Category, &dto.Description,
		&dto.IsPerennial, &dto.IsActive, &dto.CreatedAt, &dto.UpdatedAt,
		&dto.CategoryName, &dto.VarietiesCount,
	)

	if err != nil {
		return nil, err
	}

	return &dto, nil
}
