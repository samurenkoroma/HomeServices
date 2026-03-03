package persistence

import (
	"context"
	"database/sql"
	"errors"
	"samurenkoroma/services/internal/crop/application"
	"samurenkoroma/services/internal/crop/domain"
)

type CropPlanRepoImp struct {
	tx *sql.Tx
}

func NewCropRepo(tx *sql.Tx) application.CropPlanRepository {
	return &CropPlanRepoImp{
		tx: tx,
	}
}

func (r *CropPlanRepoImp) ByID(
	ctx context.Context,
	id domain.CropPlanID,
) (*domain.CropPlan, error) {

	row := r.tx.QueryRowContext(ctx, `
		SELECT id, crop_type_id, name, duration, version, status
		FROM crop_plans
		WHERE id = $1
	`, id)

	var (
		planID     string
		cropTypeID string
		name       string
		duration   int
		version    int
		status     string
	)

	if err := row.Scan(
		&planID,
		&cropTypeID,
		&name,
		&duration,
		&version,
		&status,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	plan, err := domain.NewCropPlan(
		domain.CropPlanID(planID),
		domain.CropTypeID(cropTypeID),
		name,
		duration,
	)
	if err != nil {
		return nil, err
	}

	// важно: восстановить статус и версию
	plan.Rehydrate(version, domain.CropPlanStatus(status))

	// далее загрузить stages и rotation rules
	// (сортировать по stage_order)

	return plan, nil
}

func (r *CropPlanRepoImp) Save(
	ctx context.Context,
	plan *domain.CropPlan,
) error {
	// 1. upsert crop_plans
	_, err := r.tx.ExecContext(ctx, `
		INSERT INTO crop_plans (
			id, crop_type_id, variety_id,
			name, duration, version, status,
			min_temp, max_temp, min_humidity, max_humidity,
			min_ph, max_ph,
			nitrogen, phosphorus, potassium
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)
		ON CONFLICT (id) DO UPDATE SET
			name = EXCLUDED.name,
			duration = EXCLUDED.duration,
			version = EXCLUDED.version,
			status = EXCLUDED.status
	`,
		plan.ID(),
		plan.CropTypeID(),
		plan.VarietyID(),
		plan.Name(),
		plan.Duration(),
		plan.Version(),
		plan.Status(),
		plan.Env().MinTemp(),
		plan.Env().MaxTemp(),
		plan.Env().MinHumidity(),
		plan.Env().MaxHumidity(),
		plan.Env().MinPH(),
		plan.Env().MaxPH(),
		plan.Nutrients().Nitrogen(),
		plan.Nutrients().Phosphorus(),
		plan.Nutrients().Potassium(),
	)

	if err != nil {
		return err
	}

	// 2. delete old children
	if _, err = r.tx.ExecContext(ctx,
		`DELETE FROM crop_plan_stages WHERE plan_id = $1`,
		plan.ID(),
	); err != nil {
		return err
	}

	if _, err = r.tx.ExecContext(ctx,
		`DELETE FROM crop_rotation_rules WHERE plan_id = $1`,
		plan.ID(),
	); err != nil {
		return err
	}

	// 3. insert stages
	for _, s := range plan.Stages() {
		_, err = r.tx.ExecContext(ctx, `
			INSERT INTO crop_plan_stages (
				plan_id, stage_order, name, duration,
				min_temp, max_temp, water_per_day
			)
			VALUES ($1,$2,$3,$4,$5,$6,$7)
		`,
			plan.ID(),
			s.Order(),
			s.Name(),
			s.Duration(),
			s.MinTemp(),
			s.MaxTemp(),
			s.WaterPerDay(),
		)
		if err != nil {
			return err
		}
	}

	// 4. insert rotation rules
	for _, rrule := range plan.RotationRules() {
		_, err = r.tx.ExecContext(ctx, `
			INSERT INTO crop_rotation_rules (
				plan_id, predecessor_crop_type_id, min_years
			)
			VALUES ($1,$2,$3)
		`,
			plan.ID(),
			rrule.Predecessor(),
			rrule.MinYears(),
		)
		if err != nil {
			return err
		}
	}

	return r.tx.Commit()
}
