package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"strings"
)

// CatalogRepository реализация репозитория каталога для PostgreSQL
type CatalogRepository struct {
	tx *sql.Tx
}

// NewCatalogRepository создает новый репозиторий каталога
func NewCatalogRepository(tx *sql.Tx) *CatalogRepository {
	return &CatalogRepository{tx: tx}
}

// ========== SPECIES (ВИДЫ) ==========

// GetSpecies возвращает вид по ключу
func (r *CatalogRepository) GetCrop(ctx context.Context, key string) (*catalog.Crop, error) {
	query := `
        SELECT key, name, family, category, image_url, description
        FROM public.growing_crops
        WHERE key = $1
    `

	var species catalog.Crop
	err := r.tx.QueryRowContext(ctx, query, key).Scan(
		&species.Key,
		&species.Name,
		&species.Family,
		&species.Category,
		&species.ImageUrl,
		&species.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, catalog.ErrCropNotFound
		}
		return nil, fmt.Errorf("failed to get species: %w", err)
	}

	return &species, nil
}

// ListSpecies возвращает все виды
func (r *CatalogRepository) ListCrops(ctx context.Context) ([]catalog.Crop, error) {
	query := `
        SELECT key, name, family, category, image_url, description
        FROM public.growing_crops
        ORDER BY name
    `

	rows, err := r.tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list species: %w", err)
	}
	defer rows.Close()

	var speciesList []catalog.Crop
	for rows.Next() {
		var s catalog.Crop
		err := rows.Scan(
			&s.Key,
			&s.Name,
			&s.Family,
			&s.Category,
			&s.ImageUrl,
			&s.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan species: %w", err)
		}
		speciesList = append(speciesList, s)
	}

	return speciesList, nil
}

// ListSpeciesByCategory возвращает виды по категории
func (r *CatalogRepository) ListSpeciesByCategory(ctx context.Context, category string) ([]catalog.Crop, error) {
	query := `
        SELECT key, name, family, category, image_url, description
        FROM public.growing_crops
        WHERE category = $1
        ORDER BY name
    `

	rows, err := r.tx.QueryContext(ctx, query, category)
	if err != nil {
		return nil, fmt.Errorf("failed to list species by category: %w", err)
	}
	defer rows.Close()

	var speciesList []catalog.Crop
	for rows.Next() {
		var s catalog.Crop
		err := rows.Scan(
			&s.Key,
			&s.Name,
			&s.Family,
			&s.Category,
			&s.ImageUrl,
			&s.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan species: %w", err)
		}
		speciesList = append(speciesList, s)
	}

	return speciesList, nil
}

// SaveSpecies сохраняет вид
func (r *CatalogRepository) SaveSpecies(ctx context.Context, species *catalog.Crop) error {
	query := `
        INSERT INTO public.growing_crops (key, name, family, category, image_url, description, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
        ON CONFLICT (key) DO UPDATE SET
            name = EXCLUDED.name,
            family = EXCLUDED.family,
            category = EXCLUDED.category,
            image_url = EXCLUDED.image_url,
            description = EXCLUDED.description,
            updated_at = NOW()
    `

	_, err := r.tx.ExecContext(ctx, query,
		species.Key,
		species.Name,
		species.Family,
		species.Category,
		species.ImageUrl,
		species.Description,
	)
	if err != nil {
		return fmt.Errorf("failed to save species: %w", err)
	}

	return nil
}

// DeleteSpecies удаляет вид
func (r *CatalogRepository) DeleteSpecies(ctx context.Context, key string) error {
	query := `DELETE FROM public.growing_crops WHERE key = $1`

	result, err := r.tx.ExecContext(ctx, query, key)
	if err != nil {
		return fmt.Errorf("failed to delete species: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return catalog.ErrCropNotFound
	}

	return nil
}

// ========== VARIETIES (СОРТА) ==========

// GetVariety возвращает сорт по ID
func (r *CatalogRepository) GetVariety(ctx context.Context, varietyID string) (*catalog.Variety, error) {
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
	args = []interface{}{varietyID}

	var variety catalog.Variety
	var recommendedSeasonsStr, growingTypesStr string // ПРОМЕЖУТОЧНЫЕ ПЕРЕМЕННЫЕ ДЛЯ МАССИВОВ
	var characteristicsJSON, waterReqJSON, lightReqJSON, phenophaseJSON, seedingRatesJSON []byte

	err := r.tx.QueryRowContext(ctx, query, args...).Scan(
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
		return nil, fmt.Errorf("failed to get variety: %w", err)
	}

	// Парсим массивы PostgreSQL в []string
	variety.RecommendedSeasons = parsePostgresArray(recommendedSeasonsStr)
	variety.GrowingTypes = parsePostgresArray(growingTypesStr)

	// Декодируем JSON поля
	if len(characteristicsJSON) > 0 {
		if err := json.Unmarshal(characteristicsJSON, &variety.Characteristics); err != nil {
			return nil, fmt.Errorf("failed to unmarshal characteristics: %w", err)
		}
	}

	if len(waterReqJSON) > 0 {
		if err := json.Unmarshal(waterReqJSON, &variety.WaterRequirement); err != nil {
			return nil, fmt.Errorf("failed to unmarshal water requirement: %w", err)
		}
	}

	if len(lightReqJSON) > 0 {
		if err := json.Unmarshal(lightReqJSON, &variety.LightRequirement); err != nil {
			return nil, fmt.Errorf("failed to unmarshal light requirement: %w", err)
		}
	}

	if len(phenophaseJSON) > 0 {
		if err := json.Unmarshal(phenophaseJSON, &variety.PhenophaseGDD); err != nil {
			return nil, fmt.Errorf("failed to unmarshal phenophase: %w", err)
		}
	}

	if len(seedingRatesJSON) > 0 {
		if err := json.Unmarshal(seedingRatesJSON, &variety.SeedingRates); err != nil {
			return nil, fmt.Errorf("failed to unmarshal seeding rates: %w", err)
		}
	}

	return &variety, nil
}

// ListVarieties возвращает все сорта вида
func (r *CatalogRepository) ListVarieties(ctx context.Context, speciesKey string) ([]catalog.Variety, error) {
	query := `
        SELECT id, name, species_key, species_name, base_temperature, max_temperature,
               days_to_maturity, yield_potential, plant_height,
               COALESCE(recommended_seasons, '{}') as recommended_seasons,
               COALESCE(growing_types, '{}') as growing_types,
               characteristics, description, water_requirement,
               light_requirement, phenophase_gdd, seeding_rates
        FROM public.growing_varieties
        WHERE species_key = $1
        ORDER BY name
    `

	rows, err := r.tx.QueryContext(ctx, query, speciesKey)
	if err != nil {
		return nil, fmt.Errorf("failed to list varieties: %w", err)
	}
	defer rows.Close()

	var varieties []catalog.Variety
	for rows.Next() {
		var v catalog.Variety
		var recommendedSeasonsStr, growingTypesStr string
		var characteristicsJSON, waterReqJSON, lightReqJSON, phenophaseJSON, seedingRatesJSON []byte

		err := rows.Scan(
			&v.ID,
			&v.Name,
			&v.SpeciesKey,
			&v.SpeciesName,
			&v.BaseTemperature,
			&v.MaxTemperature,
			&v.DaysToMaturity,
			&v.YieldPotential,
			&v.PlantHeight,
			&recommendedSeasonsStr,
			&growingTypesStr,
			&characteristicsJSON,
			&v.Description,
			&waterReqJSON,
			&lightReqJSON,
			&phenophaseJSON,
			&seedingRatesJSON,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan variety: %w", err)
		}

		v.RecommendedSeasons = parsePostgresArray(recommendedSeasonsStr)
		v.GrowingTypes = parsePostgresArray(growingTypesStr)

		json.Unmarshal(characteristicsJSON, &v.Characteristics)
		json.Unmarshal(waterReqJSON, &v.WaterRequirement)
		json.Unmarshal(lightReqJSON, &v.LightRequirement)
		json.Unmarshal(phenophaseJSON, &v.PhenophaseGDD)
		json.Unmarshal(seedingRatesJSON, &v.SeedingRates)

		varieties = append(varieties, v)
	}

	return varieties, nil
}

// SearchVarieties ищет сорта по фильтру
func (r *CatalogRepository) SearchVarieties(ctx context.Context, filter catalog.VarietyFilter) ([]catalog.Variety, error) {
	query := `
        SELECT id, name, species_key, species_name, base_temperature, max_temperature,
               days_to_maturity, yield_potential, plant_height,
               COALESCE(recommended_seasons, '{}') as recommended_seasons,
               COALESCE(growing_types, '{}') as growing_types,
               characteristics, description, water_requirement,
               light_requirement, phenophase_gdd, seeding_rates
        FROM public.growing_varieties
        WHERE 1=1
    `
	var args []interface{}
	argIndex := 1

	if filter.SpeciesKey != "" {
		query += fmt.Sprintf(" AND species_key = $%d", argIndex)
		args = append(args, filter.SpeciesKey)
		argIndex++
	}

	if filter.GrowingType != "" {
		query += fmt.Sprintf(" AND $%d = ANY(growing_types)", argIndex)
		args = append(args, filter.GrowingType)
		argIndex++
	}

	if filter.Season != "" {
		query += fmt.Sprintf(" AND $%d = ANY(recommended_seasons)", argIndex)
		args = append(args, filter.Season)
		argIndex++
	}

	if filter.MaxDaysToMaturity > 0 {
		query += fmt.Sprintf(" AND days_to_maturity <= $%d", argIndex)
		args = append(args, filter.MaxDaysToMaturity)
		argIndex++
	}

	if filter.Query != "" {
		query += fmt.Sprintf(" AND (name ILIKE $%d OR id ILIKE $%d)", argIndex, argIndex)
		args = append(args, "%"+filter.Query+"%")
		argIndex++
	}

	query += " ORDER BY name"

	rows, err := r.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to search varieties: %w", err)
	}
	defer rows.Close()

	var varieties []catalog.Variety
	for rows.Next() {
		var v catalog.Variety
		var recommendedSeasonsStr, growingTypesStr string
		var characteristicsJSON, waterReqJSON, lightReqJSON, phenophaseJSON, seedingRatesJSON []byte

		err := rows.Scan(
			&v.ID,
			&v.Name,
			&v.SpeciesKey,
			&v.SpeciesName,
			&v.BaseTemperature,
			&v.MaxTemperature,
			&v.DaysToMaturity,
			&v.YieldPotential,
			&v.PlantHeight,
			&recommendedSeasonsStr,
			&growingTypesStr,
			&characteristicsJSON,
			&v.Description,
			&waterReqJSON,
			&lightReqJSON,
			&phenophaseJSON,
			&seedingRatesJSON,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan variety: %w", err)
		}

		v.RecommendedSeasons = parsePostgresArray(recommendedSeasonsStr)
		v.GrowingTypes = parsePostgresArray(growingTypesStr)

		json.Unmarshal(characteristicsJSON, &v.Characteristics)
		json.Unmarshal(waterReqJSON, &v.WaterRequirement)
		json.Unmarshal(lightReqJSON, &v.LightRequirement)
		json.Unmarshal(phenophaseJSON, &v.PhenophaseGDD)
		json.Unmarshal(seedingRatesJSON, &v.SeedingRates)

		varieties = append(varieties, v)
	}

	return varieties, nil
}

// SaveVariety сохраняет сорт
func (r *CatalogRepository) SaveVariety(ctx context.Context, speciesKey string, variety *catalog.Variety) error {
	// Сериализуем JSON поля
	characteristicsJSON, err := json.Marshal(variety.Characteristics)
	if err != nil {
		return fmt.Errorf("failed to marshal characteristics: %w", err)
	}

	waterReqJSON, err := json.Marshal(variety.WaterRequirement)
	if err != nil {
		return fmt.Errorf("failed to marshal water requirement: %w", err)
	}

	lightReqJSON, err := json.Marshal(variety.LightRequirement)
	if err != nil {
		return fmt.Errorf("failed to marshal light requirement: %w", err)
	}

	phenophaseJSON, err := json.Marshal(variety.PhenophaseGDD)
	if err != nil {
		return fmt.Errorf("failed to marshal phenophase: %w", err)
	}

	seedingRatesJSON, err := json.Marshal(variety.SeedingRates)
	if err != nil {
		return fmt.Errorf("failed to marshal seeding rates: %w", err)
	}

	query := `
        INSERT INTO public.growing_varieties (
            id, name, species_key, species_name, base_temperature, max_temperature,
            days_to_maturity, yield_potential, plant_height, recommended_seasons,
            growing_types, characteristics, description, water_requirement,
            light_requirement, phenophase_gdd, seeding_rates, created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, NOW(), NOW())
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name,
            species_key = EXCLUDED.species_key,
            species_name = EXCLUDED.species_name,
            base_temperature = EXCLUDED.base_temperature,
            max_temperature = EXCLUDED.max_temperature,
            days_to_maturity = EXCLUDED.days_to_maturity,
            yield_potential = EXCLUDED.yield_potential,
            plant_height = EXCLUDED.plant_height,
            recommended_seasons = EXCLUDED.recommended_seasons,
            growing_types = EXCLUDED.growing_types,
            characteristics = EXCLUDED.characteristics,
            description = EXCLUDED.description,
            water_requirement = EXCLUDED.water_requirement,
            light_requirement = EXCLUDED.light_requirement,
            phenophase_gdd = EXCLUDED.phenophase_gdd,
            seeding_rates = EXCLUDED.seeding_rates,
            updated_at = NOW()
    `

	_, err = r.tx.ExecContext(ctx, query,
		variety.ID,
		variety.Name,
		speciesKey,
		variety.SpeciesName,
		variety.BaseTemperature,
		variety.MaxTemperature,
		variety.DaysToMaturity,
		variety.YieldPotential,
		variety.PlantHeight,
		variety.RecommendedSeasons,
		variety.GrowingTypes,
		characteristicsJSON,
		variety.Description,
		waterReqJSON,
		lightReqJSON,
		phenophaseJSON,
		seedingRatesJSON,
	)
	if err != nil {
		return fmt.Errorf("failed to save variety: %w", err)
	}

	return nil
}

// DeleteVariety удаляет сорт
func (r *CatalogRepository) DeleteVariety(ctx context.Context, speciesKey, varietyID string) error {
	query := `DELETE FROM public.growing_varieties WHERE id = $1 AND species_key = $2`

	result, err := r.tx.ExecContext(ctx, query, varietyID, speciesKey)
	if err != nil {
		return fmt.Errorf("failed to delete variety: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return catalog.ErrVarietyNotFound
	}

	return nil
}

// ========== STAGE TEMPLATES (ШАБЛОНЫ ЭТАПОВ) ==========

// GetStageTemplates возвращает шаблоны этапов для вида
func (r *CatalogRepository) GetStageTemplates(ctx context.Context, speciesKey string) ([]catalog.StageTemplate, error) {
	query := `
        SELECT type, name, bbch_start, bbch_end, description, priority, is_required, display_order
        FROM public.growing_stage_templates
        WHERE species_key = $1
        ORDER BY display_order ASC, bbch_start ASC
    `

	rows, err := r.tx.QueryContext(ctx, query, speciesKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get stage templates: %w", err)
	}
	defer rows.Close()

	var templates []catalog.StageTemplate
	for rows.Next() {
		var t catalog.StageTemplate
		var priorityStr string

		err := rows.Scan(
			&t.Type,
			&t.Name,
			&t.BBCHStart,
			&t.BBCHEnd,
			&t.Description,
			&priorityStr,
			&t.IsRequired,
			&t.DisplayOrder,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan stage template: %w", err)
		}

		t.Priority = priorityStr
		templates = append(templates, t)
	}

	return templates, nil
}

// GetStageTemplatesByBBCH возвращает шаблоны этапов для вида в определенном BBCH диапазоне
func (r *CatalogRepository) GetStageTemplatesByBBCH(ctx context.Context, speciesKey string, bbchCode int) ([]catalog.StageTemplate, error) {
	query := `
        SELECT type, name, bbch_start, bbch_end, description, priority, is_required, display_order
        FROM public.growing_stage_templates
        WHERE species_key = $1 AND bbch_start <= $2 AND bbch_end >= $2
        ORDER BY display_order ASC
    `

	rows, err := r.tx.QueryContext(ctx, query, speciesKey, bbchCode)
	if err != nil {
		return nil, fmt.Errorf("failed to get stage templates by BBCH: %w", err)
	}
	defer rows.Close()

	var templates []catalog.StageTemplate
	for rows.Next() {
		var t catalog.StageTemplate
		var priorityStr string

		err := rows.Scan(
			&t.Type,
			&t.Name,
			&t.BBCHStart,
			&t.BBCHEnd,
			&t.Description,
			&priorityStr,
			&t.IsRequired,
			&t.DisplayOrder,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan stage template: %w", err)
		}

		t.Priority = priorityStr
		templates = append(templates, t)
	}

	return templates, nil
}

// SaveStageTemplate сохраняет шаблон этапа
func (r *CatalogRepository) SaveStageTemplate(ctx context.Context, speciesKey string, template *catalog.StageTemplate) error {
	// Проверяем существование вида
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM public.growing_crops WHERE key = $1)`
	err := r.tx.QueryRowContext(ctx, checkQuery, speciesKey).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check species existence: %w", err)
	}
	if !exists {
		return catalog.ErrCropNotFound
	}

	query := `
        INSERT INTO public.growing_stage_templates (
            species_key, type, name, bbch_start, bbch_end, 
            description, priority, is_required, display_order,
            created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())
        ON CONFLICT (species_key, type, display_order) DO UPDATE SET
            name = EXCLUDED.name,
            bbch_start = EXCLUDED.bbch_start,
            bbch_end = EXCLUDED.bbch_end,
            description = EXCLUDED.description,
            priority = EXCLUDED.priority,
            is_required = EXCLUDED.is_required,
            updated_at = NOW()
    `

	_, err = r.tx.ExecContext(ctx, query,
		speciesKey,
		template.Type,
		template.Name,
		template.BBCHStart,
		template.BBCHEnd,
		template.Description,
		template.Priority,
		template.IsRequired,
		template.DisplayOrder,
	)
	if err != nil {
		return fmt.Errorf("failed to save stage template: %w", err)
	}

	return nil
}

// SaveStageTemplatesBatch массовое сохранение шаблонов этапов
func (r *CatalogRepository) SaveStageTemplatesBatch(ctx context.Context, speciesKey string, templates []catalog.StageTemplate) error {
	// Начинаем транзакцию

	// Удаляем старые шаблоны для этого вида
	deleteQuery := `DELETE FROM public.growing_stage_templates WHERE species_key = $1`
	if _, err := r.tx.ExecContext(ctx, deleteQuery, speciesKey); err != nil {
		return fmt.Errorf("failed to delete old templates: %w", err)
	}

	// Вставляем новые
	insertQuery := `
        INSERT INTO public.growing_stage_templates (
            species_key, type, name, bbch_start, bbch_end,
            description, priority, is_required, display_order,
            created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())
    `

	for _, template := range templates {
		_, err := r.tx.ExecContext(ctx, insertQuery,
			speciesKey,
			template.Type,
			template.Name,
			template.BBCHStart,
			template.BBCHEnd,
			template.Description,
			template.Priority,
			template.IsRequired,
			template.DisplayOrder,
		)
		if err != nil {
			return fmt.Errorf("failed to insert template %s: %w", template.Name, err)
		}
	}

	return nil
}

// DeleteStageTemplate удаляет шаблон этапа
func (r *CatalogRepository) DeleteStageTemplate(ctx context.Context, speciesKey string, templateType string, displayOrder int) error {
	query := `
        DELETE FROM public.growing_stage_templates 
        WHERE species_key = $1 AND type = $2 AND display_order = $3
    `

	result, err := r.tx.ExecContext(ctx, query, speciesKey, templateType, displayOrder)
	if err != nil {
		return fmt.Errorf("failed to delete stage template: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("stage template not found for species %s, type %s, order %d",
			speciesKey, templateType, displayOrder)
	}

	return nil
}

// DeleteAllStageTemplates удаляет все шаблоны этапов для вида
func (r *CatalogRepository) DeleteAllStageTemplates(ctx context.Context, speciesKey string) error {
	query := `DELETE FROM public.growing_stage_templates WHERE species_key = $1`

	_, err := r.tx.ExecContext(ctx, query, speciesKey)
	if err != nil {
		return fmt.Errorf("failed to delete all stage templates: %w", err)
	}

	return nil
}

// ========== ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ ==========

// parsePostgresArray преобразует строку массива PostgreSQL в []string
// Формат PostgreSQL массива: {value1,value2,value3} или {"value with space","value2"}
func parsePostgresArray(arrStr string) []string {
	if arrStr == "" || arrStr == "{}" {
		return []string{}
	}

	// Убираем фигурные скобки
	content := arrStr[1 : len(arrStr)-1]
	if content == "" {
		return []string{}
	}

	// Разделяем по запятой
	parts := strings.Split(content, ",")
	result := make([]string, 0, len(parts))

	for _, p := range parts {
		// Убираем кавычки, если есть
		p = strings.TrimSpace(p)
		p = strings.Trim(p, "\"")
		result = append(result, p)
	}

	return result
}
