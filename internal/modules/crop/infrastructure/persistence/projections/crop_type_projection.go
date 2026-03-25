package projections

import (
	"context"
)

// GetCropTypeWithVarieties — получить тип культуры со всеми сортами
func (p *CropProjection) GetCropTypeWithVarieties(ctx context.Context, id string) (*CropTypeWithVarietiesDTO, error) {
	// Основная информация о типе культуры
	query := `
        SELECT ct.id, ct.name, ct.category, ct.is_perennial, ct.created_at,
               COALESCE(cat.name, ct.category) as category_name,
               (SELECT COUNT(*) FROM varieties WHERE crop_type_id = ct.id AND is_active = true) as varieties_count
        FROM crop_types ct
        LEFT JOIN crop_category_translations cat ON cat.code = ct.category AND cat.language = 'ru'
        WHERE ct.id = $1 AND ct.is_active = true
    `

	var dto CropTypeWithVarietiesDTO
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
               vegetation_days_min || '–' || vegetation_days_max as vegetation_days,
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

	var varieties []VarietySimpleDTO
	for rows.Next() {
		var v VarietySimpleDTO
		if err := rows.Scan(&v.ID, &v.Name, &v.VegetationDays, &v.IsActive); err != nil {
			return nil, err
		}
		varieties = append(varieties, v)
	}

	dto.Varieties = varieties
	return &dto, nil
}

// GetAllCropTypesSimple — упрощённый список для селектов
func (p *CropProjection) GetAllCropTypesSimple(ctx context.Context) ([]CropTypeSimpleDTO, error) {
	query := `
        SELECT ct.id, ct.name, ct.category,
               COALESCE(cat.name, ct.category) as category_name
        FROM crop_types ct
        LEFT JOIN crop_category_translations cat ON cat.code = ct.category AND cat.language = 'ru'
        WHERE ct.is_active = true
        ORDER BY ct.name
    `

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []CropTypeSimpleDTO
	for rows.Next() {
		var dto CropTypeSimpleDTO
		if err := rows.Scan(&dto.ID, &dto.Name, &dto.Category, &dto.CategoryName); err != nil {
			return nil, err
		}
		result = append(result, dto)
	}

	return result, nil
}

// GetList — получить список типов культур
func (p *CropProjection) GetList(ctx context.Context, category string, activeOnly bool) ([]CropTypeSimpleDTO, error) {
	query := `
        SELECT ct.id, ct.name, ct.category, ct.is_perennial,
               COALESCE(cat.name, ct.category) as category_name
        FROM crop_types ct
        LEFT JOIN crop_category_translations cat ON cat.code = ct.category AND cat.language = 'ru'
        WHERE ($1 = '' OR ct.category = $1)
          AND ($2 = false OR ct.is_active = true)
        ORDER BY ct.name
    `

	rows, err := p.db.QueryContext(ctx, query, category, activeOnly)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []CropTypeSimpleDTO
	for rows.Next() {
		var dto CropTypeSimpleDTO
		if err := rows.Scan(&dto.ID, &dto.Name, &dto.Category, &dto.IsPerennial, &dto.CategoryName); err != nil {
			return nil, err
		}
		result = append(result, dto)
	}

	return result, nil
}

// GetByID — получить тип культуры по ID
func (p *CropProjection) GetByID(ctx context.Context, id string) (*CropTypeDetailDTO, error) {
	query := `
        SELECT ct.id, ct.name, ct.scientific_name, ct.category, ct.description, 
               ct.is_perennial, ct.is_active, ct.created_at, ct.updated_at,
               COALESCE(cat.name, ct.category) as category_name,
               (SELECT COUNT(*) FROM varieties WHERE crop_type_id = ct.id AND is_active = true) as varieties_count
        FROM crop_types ct
        LEFT JOIN crop_category_translations cat ON cat.code = ct.category AND cat.language = 'ru'
        WHERE ct.id = $1
    `

	var dto CropTypeDetailDTO
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
