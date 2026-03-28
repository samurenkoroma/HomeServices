package projections

import (
	"context"
	"samurenkoroma/services/internal/modules/crop/domain/valueobject"
	"samurenkoroma/services/internal/modules/crop/domain/variety"
)

import (
	"database/sql"
)

type varietyProjections struct {
	db *sql.DB
}

func NewVarietyProjections(db *sql.DB) variety.Projections {
	return &varietyProjections{db: db}
}

func (p *varietyProjections) GetVariety(ctx context.Context, s string) (any, error) {
	query := `
SELECT  v.id, v.name, attributes, crop_type_id,  ct.name as crop_type_name, v.is_active
FROM varieties v
LEFT OUTER JOIN public.crop_types ct on v.crop_type_id = ct.id
WHERE  v.id = $1 
	`

	row := p.db.QueryRowContext(ctx, query, s)

	var dto variety.VarietyDTO
	var attrJSON []byte
	if err := row.Scan(&dto.ID, &dto.Name, &attrJSON, &dto.CropTypeID, &dto.CropTypeName, &dto.IsActive); err != nil {
		return nil, err
	}

	var attrs variety.Attributes
	attrs.Unmarshal(attrJSON)

	dto.VegetationDays = valueobject.MinMax{
		Min: attrs.VD().Min,
		Max: attrs.VD().Max,
	}
	//dto.VegetationDaysMin = attrs.VegetDays().Min
	dto.YieldPotential = valueobject.MinMax{
		Min: attrs.YP().Min,
		Max: attrs.YP().Max,
	}

	return dto, nil
}

// GetVarieties — получить сорта
func (p *varietyProjections) GetVarieties(ctx context.Context, filter variety.Filter) ([]*variety.VarietyDTO, error) {
	// Основная информация о типе культуры
	query := `
SELECT  v.id, v.name, attributes, crop_type_id,  ct.name as crop_type_name, v.is_active
FROM varieties v
LEFT OUTER JOIN public.crop_types ct on v.crop_type_id = ct.id
WHERE  v.is_active = true ORDER BY v.name
	`

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var varieties []*variety.VarietyDTO

	for rows.Next() {
		var dto variety.VarietyDTO
		var attrJSON []byte
		if err := rows.Scan(&dto.ID, &dto.Name, &attrJSON, &dto.CropTypeID, &dto.CropTypeName, &dto.IsActive); err != nil {
			return nil, err
		}

		var attrs variety.Attributes
		attrs.Unmarshal(attrJSON)

		dto.VegetationDays = valueobject.MinMax{
			Min: attrs.VD().Min,
			Max: attrs.VD().Max,
		}
		//dto.VegetationDaysMin = attrs.VegetDays().Min
		dto.YieldPotential = valueobject.MinMax{
			Min: attrs.YP().Min,
			Max: attrs.YP().Max,
		}

		varieties = append(varieties, &dto)
	}

	return varieties, nil
}
