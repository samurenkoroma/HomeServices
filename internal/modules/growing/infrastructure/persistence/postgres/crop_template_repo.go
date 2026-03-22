package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"samurenkoroma/services/internal/modules/growing/domain/croptemplate"
)

type cropTemplateRepository struct {
	tx *sql.Tx
}

func NewCropTemplateRepository(tx *sql.Tx) croptemplate.Repository {
	return &cropTemplateRepository{tx: tx}
}

// Save сохраняет шаблон
func (r *cropTemplateRepository) Save(ctx context.Context, t *croptemplate.CropTemplate) error {
	// Сохраняем основную информацию
	query := `
        INSERT INTO crop_templates (
            id, crop_plan_id, name, version, status, requirements, created_at, published_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name,
            version = EXCLUDED.version,
            status = EXCLUDED.status,
            requirements = EXCLUDED.requirements,
            published_at = EXCLUDED.published_at,
            updated_at = EXCLUDED.updated_at
    `

	reqJSON, err := json.Marshal(t.GetRequirements())
	if err != nil {
		return fmt.Errorf("failed to marshal requirements: %w", err)
	}

	_, err = r.tx.ExecContext(ctx, query,
		string(t.GetID()),
		t.GetCropPlanID(),
		t.GetName(),
		t.GetVersion(),
		string(t.GetStatus()),
		reqJSON,
		t.GetCreatedAt(),
		t.GetPublishedAt(),
		t.GetUpdatedAt(),
	)
	if err != nil {
		return fmt.Errorf("failed to save template: %w", err)
	}

	// Сохраняем этапы
	if err := r.saveStages(ctx, t); err != nil {
		return err
	}

	return nil
}

// saveStages сохраняет этапы шаблона
func (r *cropTemplateRepository) saveStages(ctx context.Context, t *croptemplate.CropTemplate) error {
	// Удаляем старые
	_, err := r.tx.ExecContext(ctx, `DELETE FROM template_stages WHERE template_id = $1`, string(t.GetID()))
	if err != nil {
		return fmt.Errorf("failed to delete old stages: %w", err)
	}

	// Вставляем новые
	for _, stage := range t.GetStages() {
		_, err := r.tx.ExecContext(ctx, `
            INSERT INTO template_stages (
                template_id, stage_order, name, duration, min_temp, max_temp, optimal_temp, water_per_day, description
            ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        `,
			string(t.GetID()),
			stage.Order,
			stage.Name,
			stage.Duration,
			stage.MinTemp,
			stage.MaxTemp,
			stage.OptimalTemp,
			stage.WaterPerDay,
			stage.Description,
		)
		if err != nil {
			return fmt.Errorf("failed to insert stage: %w", err)
		}
	}

	return nil
}

// FindByID находит шаблон по ID
func (r *cropTemplateRepository) FindByID(ctx context.Context, id croptemplate.TemplateID) (*croptemplate.CropTemplate, error) {
	query := `
        SELECT id, crop_plan_id, name, version, status, requirements, created_at, published_at, updated_at
        FROM crop_templates
        WHERE id = $1
    `

	var (
		tid         string
		cropPlanID  string
		name        string
		version     int
		status      string
		reqJSON     []byte
		createdAt   time.Time
		publishedAt sql.NullTime
		updatedAt   time.Time
	)

	err := r.tx.QueryRowContext(ctx, query, string(id)).Scan(
		&tid, &cropPlanID, &name, &version, &status, &reqJSON,
		&createdAt, &publishedAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, croptemplate.ErrTemplateNotFound
		}
		return nil, fmt.Errorf("failed to find template: %w", err)
	}

	// Загружаем этапы
	stages, err := r.loadStages(ctx, tid)
	if err != nil {
		return nil, err
	}

	// Парсим требования
	var requirements croptemplate.Requirements
	if len(reqJSON) > 0 {
		if err := json.Unmarshal(reqJSON, &requirements); err != nil {
			return nil, fmt.Errorf("failed to unmarshal requirements: %w", err)
		}
	}

	// Создаём шаблон
	template := croptemplate.NewCropTemplate(cropPlanID, name, version)

	var pubAt *time.Time
	if publishedAt.Valid {
		pubAt = &publishedAt.Time
	}

	template.Rehydrate(
		croptemplate.TemplateID(tid),
		cropPlanID,
		name,
		version,
		croptemplate.TemplateStatus(status),
		stages,
		requirements,
		createdAt,
		pubAt,
		updatedAt,
	)

	return template, nil
}

// FindByCropPlanID находит шаблон по ID плана (последнюю версию)
func (r *cropTemplateRepository) FindByCropPlanID(ctx context.Context, cropPlanID string) (*croptemplate.CropTemplate, error) {
	query := `
        SELECT id, crop_plan_id, name, version, status, requirements, created_at, published_at, updated_at
        FROM crop_templates
        WHERE crop_plan_id = $1
        ORDER BY version DESC
        LIMIT 1
    `

	var (
		tid         string
		cpid        string
		name        string
		version     int
		status      string
		reqJSON     []byte
		createdAt   time.Time
		publishedAt sql.NullTime
		updatedAt   time.Time
	)

	err := r.tx.QueryRowContext(ctx, query, cropPlanID).Scan(
		&tid, &cpid, &name, &version, &status, &reqJSON,
		&createdAt, &publishedAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find template by crop plan: %w", err)
	}

	stages, err := r.loadStages(ctx, tid)
	if err != nil {
		return nil, err
	}

	var requirements croptemplate.Requirements
	if len(reqJSON) > 0 {
		if err := json.Unmarshal(reqJSON, &requirements); err != nil {
			return nil, fmt.Errorf("failed to unmarshal requirements: %w", err)
		}
	}

	template := croptemplate.NewCropTemplate(cpid, name, version)

	var pubAt *time.Time
	if publishedAt.Valid {
		pubAt = &publishedAt.Time
	}

	template.Rehydrate(
		croptemplate.TemplateID(tid),
		cpid,
		name,
		version,
		croptemplate.TemplateStatus(status),
		stages,
		requirements,
		createdAt,
		pubAt,
		updatedAt,
	)

	return template, nil
}

// FindByCropPlanIDAndVersion находит шаблон по ID плана и версии
func (r *cropTemplateRepository) FindByCropPlanIDAndVersion(ctx context.Context, cropPlanID string, version int) (*croptemplate.CropTemplate, error) {
	query := `
        SELECT id, crop_plan_id, name, version, status, requirements, created_at, published_at, updated_at
        FROM crop_templates
        WHERE crop_plan_id = $1 AND version = $2
    `

	var (
		tid         string
		cpid        string
		name        string
		ver         int
		status      string
		reqJSON     []byte
		createdAt   time.Time
		publishedAt sql.NullTime
		updatedAt   time.Time
	)

	err := r.tx.QueryRowContext(ctx, query, cropPlanID, version).Scan(
		&tid, &cpid, &name, &ver, &status, &reqJSON,
		&createdAt, &publishedAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, croptemplate.ErrTemplateNotFound
		}
		return nil, fmt.Errorf("failed to find template: %w", err)
	}

	stages, err := r.loadStages(ctx, tid)
	if err != nil {
		return nil, err
	}

	var requirements croptemplate.Requirements
	if len(reqJSON) > 0 {
		if err := json.Unmarshal(reqJSON, &requirements); err != nil {
			return nil, fmt.Errorf("failed to unmarshal requirements: %w", err)
		}
	}

	template := croptemplate.NewCropTemplate(cpid, name, ver)

	var pubAt *time.Time
	if publishedAt.Valid {
		pubAt = &publishedAt.Time
	}

	template.Rehydrate(
		croptemplate.TemplateID(tid),
		cpid,
		name,
		ver,
		croptemplate.TemplateStatus(status),
		stages,
		requirements,
		createdAt,
		pubAt,
		updatedAt,
	)

	return template, nil
}

// FindAll возвращает все шаблоны
func (r *cropTemplateRepository) FindAll(ctx context.Context) ([]*croptemplate.CropTemplate, error) {
	query := `
        SELECT id, crop_plan_id, name, version, status, requirements, created_at, published_at, updated_at
        FROM crop_templates
        ORDER BY name, version DESC
    `

	rows, err := r.tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query templates: %w", err)
	}
	defer rows.Close()

	var templates []*croptemplate.CropTemplate

	for rows.Next() {
		var (
			tid         string
			cropPlanID  string
			name        string
			version     int
			status      string
			reqJSON     []byte
			createdAt   time.Time
			publishedAt sql.NullTime
			updatedAt   time.Time
		)

		err := rows.Scan(
			&tid, &cropPlanID, &name, &version, &status, &reqJSON,
			&createdAt, &publishedAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan template: %w", err)
		}

		stages, err := r.loadStages(ctx, tid)
		if err != nil {
			return nil, err
		}

		var requirements croptemplate.Requirements
		if len(reqJSON) > 0 {
			if err := json.Unmarshal(reqJSON, &requirements); err != nil {
				return nil, fmt.Errorf("failed to unmarshal requirements: %w", err)
			}
		}

		template := croptemplate.NewCropTemplate(cropPlanID, name, version)

		var pubAt *time.Time
		if publishedAt.Valid {
			pubAt = &publishedAt.Time
		}

		template.Rehydrate(
			croptemplate.TemplateID(tid),
			cropPlanID,
			name,
			version,
			croptemplate.TemplateStatus(status),
			stages,
			requirements,
			createdAt,
			pubAt,
			updatedAt,
		)

		templates = append(templates, template)
	}

	return templates, nil
}

// FindPublished возвращает все опубликованные шаблоны
func (r *cropTemplateRepository) FindPublished(ctx context.Context) ([]*croptemplate.CropTemplate, error) {
	query := `
        SELECT id, crop_plan_id, name, version, status, requirements, created_at, published_at, updated_at
        FROM crop_templates
        WHERE status = 'published'
        ORDER BY name, version DESC
    `

	rows, err := r.tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query published templates: %w", err)
	}
	defer rows.Close()

	var templates []*croptemplate.CropTemplate

	for rows.Next() {
		var (
			tid         string
			cropPlanID  string
			name        string
			version     int
			status      string
			reqJSON     []byte
			createdAt   time.Time
			publishedAt sql.NullTime
			updatedAt   time.Time
		)

		err := rows.Scan(
			&tid, &cropPlanID, &name, &version, &status, &reqJSON,
			&createdAt, &publishedAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan template: %w", err)
		}

		stages, err := r.loadStages(ctx, tid)
		if err != nil {
			return nil, err
		}

		var requirements croptemplate.Requirements
		if len(reqJSON) > 0 {
			if err := json.Unmarshal(reqJSON, &requirements); err != nil {
				return nil, fmt.Errorf("failed to unmarshal requirements: %w", err)
			}
		}

		template := croptemplate.NewCropTemplate(cropPlanID, name, version)

		var pubAt *time.Time
		if publishedAt.Valid {
			pubAt = &publishedAt.Time
		}

		template.Rehydrate(
			croptemplate.TemplateID(tid),
			cropPlanID,
			name,
			version,
			croptemplate.TemplateStatus(status),
			stages,
			requirements,
			createdAt,
			pubAt,
			updatedAt,
		)

		templates = append(templates, template)
	}

	return templates, nil
}

// Delete удаляет шаблон
func (r *cropTemplateRepository) Delete(ctx context.Context, id croptemplate.TemplateID) error {
	// Проверяем, есть ли активные циклы, использующие этот шаблон
	var cycleCount int
	err := r.tx.QueryRowContext(ctx, `
        SELECT COUNT(*) FROM crop_cycles WHERE template_id = $1 AND status NOT IN ('completed', 'cancelled')
    `, string(id)).Scan(&cycleCount)
	if err != nil {
		return fmt.Errorf("failed to check cycles: %w", err)
	}

	if cycleCount > 0 {
		return fmt.Errorf("cannot delete template with active cycles")
	}

	// Этажи удалятся каскадно
	_, err = r.tx.ExecContext(ctx, `DELETE FROM crop_templates WHERE id = $1`, string(id))
	if err != nil {
		return fmt.Errorf("failed to delete template: %w", err)
	}

	return nil
}

// Exists проверяет существование шаблона
func (r *cropTemplateRepository) Exists(ctx context.Context, cropPlanID string, version int) (bool, error) {
	var exists bool
	err := r.tx.QueryRowContext(ctx, `
        SELECT EXISTS(SELECT 1 FROM crop_templates WHERE crop_plan_id = $1 AND version = $2)
    `, cropPlanID, version).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check existence: %w", err)
	}
	return exists, nil
}

// GetLatestVersion возвращает последнюю версию шаблона для плана
func (r *cropTemplateRepository) GetLatestVersion(ctx context.Context, cropPlanID string) (int, error) {
	var version int
	err := r.tx.QueryRowContext(ctx, `
        SELECT COALESCE(MAX(version), 0) FROM crop_templates WHERE crop_plan_id = $1
    `, cropPlanID).Scan(&version)
	if err != nil {
		return 0, fmt.Errorf("failed to get latest version: %w", err)
	}
	return version, nil
}

// loadStages загружает этапы шаблона
func (r *cropTemplateRepository) loadStages(ctx context.Context, templateID string) ([]croptemplate.GrowthStage, error) {
	rows, err := r.tx.QueryContext(ctx, `
        SELECT stage_order, name, duration, min_temp, max_temp, optimal_temp, water_per_day, description
        FROM template_stages
        WHERE template_id = $1
        ORDER BY stage_order
    `, templateID)
	if err != nil {
		return nil, fmt.Errorf("failed to query stages: %w", err)
	}
	defer rows.Close()

	var stages []croptemplate.GrowthStage
	for rows.Next() {
		var stage croptemplate.GrowthStage
		err := rows.Scan(
			&stage.Order,
			&stage.Name,
			&stage.Duration,
			&stage.MinTemp,
			&stage.MaxTemp,
			&stage.OptimalTemp,
			&stage.WaterPerDay,
			&stage.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan stage: %w", err)
		}
		stages = append(stages, stage)
	}

	return stages, nil
}
