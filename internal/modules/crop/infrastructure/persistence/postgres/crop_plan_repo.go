package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"samurenkoroma/services/internal/modules/crop/domain/cropplan"
	"time"
)

type cropPlanRepository struct {
	tx *sql.Tx
}

func NewCropPlanRepository(tx *sql.Tx) cropplan.Repository {
	return &cropPlanRepository{tx: tx}
}

func (r *cropPlanRepository) Save(ctx context.Context, plan *cropplan.CropPlan) error {
	// Основная информация
	query := `
        INSERT INTO crop_crop_plans (
            id, crop_type_id, variety_id, name, description, duration,
            version, status, environment, nutrients, created_by,
            created_at, updated_at, published_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name,
            description = EXCLUDED.description,
            duration = EXCLUDED.duration,
            version = EXCLUDED.version,
            status = EXCLUDED.status,
            environment = EXCLUDED.environment,
            nutrients = EXCLUDED.nutrients,
            updated_at = EXCLUDED.updated_at,
            published_at = EXCLUDED.published_at
    `

	envJSON, _ := json.Marshal(plan.GetEnvironment())
	nutJSON, _ := json.Marshal(plan.GetNutrients())

	_, err := r.tx.ExecContext(ctx, query,
		string(plan.GetID()),
		plan.GetCropTypeID(),
		plan.GetVarietyID(),
		plan.GetName(),
		plan.Description,
		plan.GetDuration(),
		plan.GetVersion(),
		string(plan.GetStatus()),
		envJSON,
		nutJSON,
		plan.CreatedBy,
		plan.CreatedAt,
		plan.UpdatedAt,
		plan.PublishedAt,
	)
	if err != nil {
		return err
	}

	// Сохраняем этапы
	if err := r.saveStages(ctx, plan); err != nil {
		return err
	}

	// Сохраняем правила севооборота
	if err := r.saveRotationRules(ctx, plan); err != nil {
		return err
	}

	return nil
}

func (r *cropPlanRepository) saveStages(ctx context.Context, plan *cropplan.CropPlan) error {
	// Удаляем старые
	_, err := r.tx.ExecContext(ctx, `DELETE FROM crop_crop_plan_stages WHERE plan_id = $1`, string(plan.GetID()))
	if err != nil {
		return err
	}

	// Вставляем новые
	for _, stage := range plan.GetStages() {
		recommendData, err := stage.Recommendations.Marshal()
		if err != nil {
			return err
		}
		_, err = r.tx.ExecContext(ctx, `
            INSERT INTO crop_crop_plan_stages (
                plan_id, stage_order, name, duration, recommendations
            ) VALUES ($1, $2, $3, $4, $5)
        `,
			string(plan.GetID()),
			stage.Order,
			stage.Name,
			stage.Duration,
			recommendData,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *cropPlanRepository) saveRotationRules(ctx context.Context, plan *cropplan.CropPlan) error {
	// Удаляем старые
	_, err := r.tx.ExecContext(ctx, `DELETE FROM crop_crop_rotation_rules WHERE plan_id = $1`, string(plan.GetID()))
	if err != nil {
		return err
	}

	// Вставляем новые
	for _, rule := range plan.GetRotationRules() {
		_, err := r.tx.ExecContext(ctx, `
            INSERT INTO crop_crop_rotation_rules (
                plan_id, predecessor_crop_type_id, min_years, recommended, notes
            ) VALUES ($1, $2, $3, $4, $5)
        `,
			string(plan.GetID()),
			rule.PredecessorCropTypeID,
			rule.MinYears,
			rule.Recommended,
			rule.Notes,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *cropPlanRepository) GetByID(ctx context.Context, id cropplan.PlanID) (*cropplan.CropPlan, error) {
	query := `
        SELECT 
            id, crop_type_id, variety_id, name, description, duration,
            version, status, environment, nutrients, created_by,
            created_at, updated_at, published_at
        FROM crop_crop_plans
        WHERE id = $1
    `

	var (
		planID      string
		cropTypeID  string
		varietyID   *string
		name        string
		description string
		duration    int
		version     int
		status      string
		envJSON     []byte
		nutJSON     []byte
		createdBy   string
		createdAt   time.Time
		updatedAt   time.Time
		publishedAt *time.Time
	)

	err := r.tx.QueryRowContext(ctx, query, string(id)).Scan(
		&planID, &cropTypeID, &varietyID, &name, &description, &duration,
		&version, &status, &envJSON, &nutJSON, &createdBy,
		&createdAt, &updatedAt, &publishedAt,
	)
	if err != nil {
		return nil, err
	}

	// Восстанавливаем план
	plan, err := cropplan.NewCropPlan(cropTypeID, *varietyID, name, duration, createdBy)
	if err != nil {
		return nil, err
	}

	plan.ID = cropplan.PlanID(planID)
	plan.Description = description
	plan.VarietyID = varietyID
	plan.Rehydrate(version, cropplan.PlanStatus(status))
	plan.CreatedAt = createdAt
	plan.UpdatedAt = updatedAt
	plan.PublishedAt = publishedAt

	// Загружаем этапы
	stages, err := r.loadStages(ctx, string(plan.ID))
	if err != nil {
		return nil, err
	}
	plan.Stages = stages

	// Загружаем правила севооборота
	rules, err := r.loadRotationRules(ctx, string(plan.ID))
	if err != nil {
		return nil, err
	}
	plan.RotationRules = rules

	return plan, nil
}

func (r *cropPlanRepository) loadStages(ctx context.Context, planID string) ([]cropplan.GrowthStage, error) {
	rows, err := r.tx.QueryContext(ctx, `
        SELECT stage_order, name, duration, recommendations
        FROM crop_crop_plan_stages
        WHERE plan_id = $1
        ORDER BY stage_order
    `, planID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stages []cropplan.GrowthStage
	for rows.Next() {
		var stage cropplan.GrowthStage
		var attrJSON []byte
		if err := rows.Scan(&stage.Order, &stage.Name, &stage.Duration, &attrJSON); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(attrJSON, &stage.Recommendations); err != nil {
			return nil, err
		}
		stages = append(stages, stage)
	}

	return stages, nil
}

func (r *cropPlanRepository) loadRotationRules(ctx context.Context, planID string) ([]cropplan.RotationRule, error) {
	rows, err := r.tx.QueryContext(ctx, `
        SELECT predecessor_crop_type_id, min_years, recommended, notes
        FROM crop_crop_rotation_rules
        WHERE plan_id = $1
    `, planID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []cropplan.RotationRule
	for rows.Next() {
		var rule cropplan.RotationRule
		err := rows.Scan(
			&rule.PredecessorCropTypeID,
			&rule.MinYears,
			&rule.Recommended,
			&rule.Notes,
		)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}

	return rules, nil
}
