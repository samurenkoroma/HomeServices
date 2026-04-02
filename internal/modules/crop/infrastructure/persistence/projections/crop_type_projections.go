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
        SELECT ct.id, 
               ct.name, 
               ct.category,
               ct.family,
               ct.description,
               ct.is_perennial,
               ct.imageurl
        FROM crop_crop_types ct
        WHERE ct.id = $1 AND ct.is_active = true
    `

	var dto croptype.CropTypeWithVarietiesDTO

	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&dto.ID, &dto.Name, &dto.Category, &dto.Family, &dto.Description, &dto.IsPerennial, &dto.ImageUrl,
	)
	if err != nil {
		return nil, err
	}

	// Получаем сорта
	varietiesQuery := `
        SELECT id, name, 
               attributes->>'YieldPotential' as yieldEstimate,
               attributes->>'VegetationDays' as growingTime,
               'Март - Апрель' as plantingTime,
        '' as image
        FROM crop_varieties
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
		if err := rows.Scan(&v.ID, &v.Name, &v.YieldPotential, &v.VegetationDays, &v.PlantingTime, &v.Image); err != nil {
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
        FROM crop_crop_types ct
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
		if err := rows.Scan(&dto.ID, &dto.Name, &dto.Category, &dto.Category); err != nil {
			return nil, err
		}
		result = append(result, dto)
	}

	return result, nil
}

// GetList — получить список типов культур
func (p *cropTypeProjections) GetList(ctx context.Context, filter croptype.Filter) ([]croptype.CropTypeSimpleDTO, error) {
	query := `
SELECT ct.id, ct.name, ct.icon, ct.imageurl,
	   (select ru from translations where entity = 'crop_category' and latin = ct.category) as category,
       (select ru from translations where entity = 'crop_family' and latin = ct.family) as family,
--        ct.category, ct.family,
       
       (SELECT count(*) from crop_varieties cv where cv.crop_type_id = ct.id) as countVarieties,
       (SELECT min(split_part(cv.attributes->>'YieldPotential', '..', 1)::int) from crop_varieties cv where  cv.crop_type_id = ct.id) as productivityMin,
       (SELECT min(split_part(cv.attributes->>'YieldPotential', '..', 2)::int) from crop_varieties cv where  cv.crop_type_id = ct.id) as productivityMax,
       (SELECT min(split_part(cv.attributes->>'VegetationDays', '..', 1)::int) from crop_varieties cv where  cv.crop_type_id = ct.id) as growingPeriodMin,
       (SELECT min(split_part(cv.attributes->>'VegetationDays', '..', 2)::int) from crop_varieties cv where  cv.crop_type_id = ct.id) as growingPeriodMax
FROM crop_crop_types ct
        WHERE ($1 = '' OR ct.category = $1)
          AND ($2 = '' OR ct.family = $2)
          AND ($3 = false OR ct.is_active = true)
ORDER BY ct.name;

    `

	rows, err := p.db.QueryContext(ctx, query, filter.Category, filter.Family, filter.IsActive)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []croptype.CropTypeSimpleDTO
	for rows.Next() {
		var dto croptype.CropTypeSimpleDTO
		if err := rows.Scan(&dto.ID, &dto.Name, &dto.Icon, &dto.ImageUrl, &dto.Category, &dto.Family, &dto.CountVarieties,
			&dto.YieldEstimateMin, &dto.YieldEstimateMax, &dto.GrowingDaysMin, &dto.GrowingDaysMax); err != nil {
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
               (SELECT COUNT(*) FROM crop_varieties WHERE crop_type_id = ct.id AND is_active = true) as varieties_count
        FROM crop_crop_types ct
--         LEFT JOIN crop_category_translations cat ON cat.code = ct.category AND cat.language = 'ru'
        WHERE ct.id = $1
    `

	var dto croptype.CropTypeDetailDTO
	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&dto.ID, &dto.Name, &dto.Category, &dto.Description,
		&dto.IsPerennial, &dto.IsActive, &dto.CreatedAt, &dto.UpdatedAt,
		&dto.CategoryName, &dto.VarietiesCount,
	)

	if err != nil {
		return nil, err
	}

	return &dto, nil
}
