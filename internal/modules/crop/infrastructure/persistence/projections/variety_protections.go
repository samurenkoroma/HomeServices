package projections

import (
	"context"
	"samurenkoroma/services/internal/modules/crop/domain/variety"
)

// GetCropTypeWithVarieties — получить тип культуры со всеми сортами
func (p *CropProjection) GetVarieties(ctx context.Context) ([]*VarietyDTO, error) {
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

	var varieties []*VarietyDTO

	for rows.Next() {
		var dto VarietyDTO
		var attrJSON []byte
		if err := rows.Scan(&dto.ID, &dto.Name, &attrJSON, &dto.CropTypeID, &dto.CropTypeName, &dto.IsActive); err != nil {
			return nil, err
		}

		var attrs variety.Attributes
		attrs.Unmarshal(attrJSON)

		dto.VegetationDays = MinMax{
			Min: attrs.VD().Min,
			Max: attrs.VD().Max,
		}
		//dto.VegetationDaysMin = attrs.VegetDays().Min
		dto.YieldPotential = MinMax{
			Min: attrs.YP().Min,
			Max: attrs.YP().Max,
		}

		varieties = append(varieties, &dto)
	}

	return varieties, nil
}
