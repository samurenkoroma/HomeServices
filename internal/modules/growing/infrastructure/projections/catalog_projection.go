package projections

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"samurenkoroma/services/internal/infrastructure/persistence"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"time"
)

type catalogProjection struct {
	db persistence.DBTX
}

func (c catalogProjection) GetCropPlans(ctx context.Context, filter catalog.CropPlanFilter) ([]catalog.CropPlanListItemDTO, error) {
	query := `
        SELECT 
           plan.id,
       plan.status,
       plan.expected_harvest_date,
       plan.planting_date,
       crop.key,
       crop.name,
       variety.id,
       variety.name,
       variety.days_to_maturity,
       prod_unit.id,
       prod_unit.area,
       prod_unit.name,
       cultivator.id,
       cultivator.name
FROM growing_crop_plans plan
         LEFT JOIN growing_crops crop on crop.key = plan.crop_key
         LEFT JOIN growing_varieties variety on variety.id = plan.variety_id
         LEFT JOIN growing_cultivation_areas prod_unit on prod_unit.id = plan.area_id
         LEFT JOIN growing_cultivation_plans cultivator on cultivator.id = plan.cultivation_plan_id
--         WHERE pu.owner_id IN (
--             SELECT id FROM farm_physical_objects WHERE  IN (
--                 SELECT id FROM land_structure WHERE unit_type = 'land' AND root_id = $1
--                 UNION
--                 SELECT $1
--             )
--         )
        ORDER BY plan.planting_date ASC 
    `

	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query plans: %w", err)
	}
	defer rows.Close()

	var plans []catalog.CropPlanListItemDTO

	for rows.Next() {
		var p catalog.CropPlanListItemDTO
		var cropName, cropKey, cropNameFull, prodUnitId, prodUnitName, cultivatorId, cultivatorName string
		var prodUnitArea float64
		var daysToMaturity sql.NullInt16
		var varietyID, varietyName sql.NullString

		var plantingDate, expectedHarvestDate time.Time

		err := rows.Scan(
			&p.ID,
			&p.Status,
			&expectedHarvestDate,
			&plantingDate,

			&cropKey,
			&cropName,

			&varietyID,
			&varietyName,
			&daysToMaturity,
			&prodUnitId,
			&prodUnitArea,
			&prodUnitName,

			&cultivatorId,
			&cultivatorName,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan plan: %w", err)
		}

		// Заполняем DTO
		p.PlantingDate = plantingDate.Format("2006-01-02")
		p.ExpectedHarvestDate = expectedHarvestDate.Format("2006-01-02")

		// Культура
		p.Crop = catalog.CropDTO{
			Key:  cropKey,
			Name: cropNameFull,
		}
		if p.Crop.Name == "" {
			p.Crop.Name = cropName
		}

		// Сорт
		if varietyID.Valid {
			p.Variety = catalog.VarietyDTO{
				ID:             varietyID.String,
				Name:           varietyName.String,
				DaysToMaturity: daysToMaturity.Int16,
			}
		}

		if p.Variety.Name == "" {
			p.Variety.Name = cropName
		}

		// Производственная единица
		p.ProductionUnit = catalog.UnitDTO{
			ID:   prodUnitId,
			Area: prodUnitArea,
			Name: prodUnitName,
		}

		// Шаблон плана (пока заглушка)
		p.CultivationPlan = catalog.PlanRefDTO{
			ID:   cultivatorId,
			Name: cultivatorName,
		}

		// Прогресс (расчет на основе завершенных этапов)
		p.Progress = 0 // нужно будет рассчитать

		plans = append(plans, p)
	}

	return plans, nil

}

func (c catalogProjection) GetCultivationPlans(ctx context.Context, filter catalog.CultivationPlansFilter) ([]catalog.CultivationPlanItem, error) {
	query := `
	SELECT id, version, name, crop_key, steps
	FROM growing_cultivation_plans
	WHERE crop_key = $1
	ORDER BY id, version DESC
	`

	rows, err := c.db.QueryContext(ctx, query, filter.CropKey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []catalog.CultivationPlanItem{}

	for rows.Next() {
		var p catalog.CultivationPlanItem
		var stepsBytes []byte

		err := rows.Scan(
			&p.ID,
			&p.Version,
			&p.Name,
			&p.CropKey,
			&stepsBytes,
		)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(stepsBytes, &p.Steps); err != nil {
			return nil, err
		}

		result = append(result, p)
	}

	return result, nil
}

func (c catalogProjection) GetSeasons(ctx context.Context, filter catalog.SeasonFilter) ([]catalog.SeasonItem, error) {
	query := `
SELECT id, name, start_date, end_date, description, status
FROM growing_seasons
WHERE  created_by = $1 ORDER BY start_date DESC`

	rows, err := c.db.QueryContext(ctx, query, filter.OwnerId)
	if err != nil {
		return nil, fmt.Errorf("failed to query seasons: %w", err)
	}
	defer rows.Close()
	var items []catalog.SeasonItem
	for rows.Next() {
		var item catalog.SeasonItem
		if err := rows.Scan(&item.Id, &item.Name, &item.StartDate, &item.EndDate, &item.Description, &item.Status); err != nil {
			return nil, fmt.Errorf("failed to scan season: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate seasons: %w", err)
	}
	return items, nil
}

func (c catalogProjection) GetCrops(ctx context.Context) ([]catalog.CropDto, error) {
	query := `
        SELECT key, name, family, category, image_url, description
        FROM public.growing_crops
        ORDER BY name
    `

	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var crops []catalog.CropDto
	for rows.Next() {
		var s catalog.CropDto
		err := rows.Scan(
			&s.Key,
			&s.Name,
			&s.Family,
			&s.Category,
			&s.ImageUrl,
			&s.Description,
		)
		if err != nil {
			return nil, err
		}
		crops = append(crops, s)
	}

	return crops, nil
}

func (c catalogProjection) GetVarieties(ctx context.Context, cropKey string) ([]catalog.VarietyItem, error) {
	query := `
        SELECT id, name,description, yield_potential, plant_height, COALESCE(growing_types, '{}') as growing_types
        FROM public.growing_varieties
        WHERE species_key = $1
        ORDER BY name
    `
	rows, err := c.db.QueryContext(ctx, query, cropKey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var varieties []catalog.VarietyItem
	for rows.Next() {
		var v catalog.VarietyItem
		var growingTypesStr string

		err := rows.Scan(
			&v.ID,
			&v.Name,
			&v.Desc,
			&v.YieldPotential,
			&v.PlantHeight,
			&growingTypesStr,
		)
		if err != nil {
			return nil, err
		}

		v.GrowingTypes = persistence.ParsePostgresArray(growingTypesStr)

		varieties = append(varieties, v)
	}

	return varieties, nil
}

func (c catalogProjection) GetVariety(ctx context.Context, id string) (*catalog.VarietyDetail, error) {
	var query string
	var args []interface{}

	query = `
            SELECT id, name, species_key, species_name, base_temperature, max_temperature,
                   days_to_maturity, yield_potential, plant_height,
                   COALESCE(recommended_seasons, '{}') as recommended_seasons,
                   COALESCE(growing_types, '{}') as growing_types,
                   characteristics, description, water_requirement,
                   light_requirement, phenophase_gdd, seeding_rates
            FROM public.growing_varieties
            WHERE id = $1
        `
	args = []interface{}{id}

	var variety catalog.VarietyDetail
	var recommendedSeasonsStr, growingTypesStr string // ПРОМЕЖУТОЧНЫЕ ПЕРЕМЕННЫЕ ДЛЯ МАССИВОВ
	var characteristicsJSON, waterReqJSON, lightReqJSON, phenophaseJSON, seedingRatesJSON []byte

	err := c.db.QueryRowContext(ctx, query, args...).Scan(
		&variety.ID,
		&variety.Name,
		&variety.SpeciesKey,
		&variety.SpeciesName,
		&variety.BaseTemperature,
		&variety.MaxTemperature,
		&variety.DaysToMaturity,
		&variety.YieldPotential,
		&variety.PlantHeight,
		&recommendedSeasonsStr,
		&growingTypesStr,
		&characteristicsJSON,
		&variety.Description,
		&waterReqJSON,
		&lightReqJSON,
		&phenophaseJSON,
		&seedingRatesJSON,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, catalog.ErrVarietyNotFound
		}
		return nil, err
	}

	// Парсим массивы PostgreSQL в []string
	variety.RecommendedSeasons = persistence.ParsePostgresArray(recommendedSeasonsStr)
	variety.GrowingTypes = persistence.ParsePostgresArray(growingTypesStr)

	// Декодируем JSON поля
	if len(characteristicsJSON) > 0 {
		if err := json.Unmarshal(characteristicsJSON, &variety.Characteristics); err != nil {
			return nil, err
		}
	}

	if len(waterReqJSON) > 0 {
		if err := json.Unmarshal(waterReqJSON, &variety.WaterRequirement); err != nil {
			return nil, err
		}
	}

	if len(lightReqJSON) > 0 {
		if err := json.Unmarshal(lightReqJSON, &variety.LightRequirement); err != nil {
			return nil, err
		}
	}

	if len(phenophaseJSON) > 0 {
		if err := json.Unmarshal(phenophaseJSON, &variety.PhenophaseGDD); err != nil {
			return nil, err
		}
	}

	if len(seedingRatesJSON) > 0 {
		if err := json.Unmarshal(seedingRatesJSON, &variety.SeedingRates); err != nil {
			return nil, err
		}
	}

	return &variety, nil
}

func NewCatalogProjection(db persistence.DBTX) catalog.CatalogProjections {
	return &catalogProjection{db: db}
}
