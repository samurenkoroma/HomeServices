package postgres

import (
	"context"
	"encoding/json"
	"samurenkoroma/services/internal/infrastructure/persistence"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cultivation"
)

type cultivationPlanRepo struct {
	db persistence.DBTX
}

func NewCultivationPlanRepository(db persistence.DBTX) cultivation.Repository {
	return &cultivationPlanRepo{db: db}
}

func (r *cultivationPlanRepo) Save(ctx context.Context, p *cultivation.CultivationPlan) error {
	stepsBytes, err := json.Marshal(p.Steps)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO growing_cultivation_plans (
		id, version, name, crop_key, steps, created_at
	)
	VALUES ($1,$2,$3,$4,$5,$6)
	ON CONFLICT (id, version) DO NOTHING
	`

	_, err = r.db.ExecContext(ctx, query,
		p.ID,
		p.Version,
		p.Name,
		p.CropKey,
		stepsBytes,
		p.CreatedAt,
	)

	return err
}

func (r *cultivationPlanRepo) Get(ctx context.Context, id string, version int) (*cultivation.CultivationPlan, error) {
	query := `
	SELECT id, version, name, crop_key, steps, created_at
	FROM growing_cultivation_plans
	WHERE id = $1 AND version = $2
	`

	row := r.db.QueryRowContext(ctx, query, id, version)

	var p cultivation.CultivationPlan
	var stepsBytes []byte

	err := row.Scan(
		&p.ID,
		&p.Version,
		&p.Name,
		&p.CropKey,
		&stepsBytes,
		&p.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(stepsBytes, &p.Steps); err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *cultivationPlanRepo) GetLatest(ctx context.Context, id string) (*cultivation.CultivationPlan, error) {
	query := `
	SELECT id, version, name, crop_key, steps, created_at
	FROM growing_cultivation_plans
	WHERE id = $1
	ORDER BY version DESC
	LIMIT 1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var p cultivation.CultivationPlan
	var stepsBytes []byte

	err := row.Scan(
		&p.ID,
		&p.Version,
		&p.Name,
		&p.CropKey,
		&stepsBytes,
		&p.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(stepsBytes, &p.Steps); err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *cultivationPlanRepo) List(ctx context.Context, cropKey string) ([]*cultivation.CultivationPlan, error) {
	query := `
	SELECT id, version, name, crop_key, steps, created_at
	FROM growing_cultivation_plans
	WHERE crop_key = $1
	ORDER BY id, version DESC
	`

	rows, err := r.db.QueryContext(ctx, query, cropKey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*cultivation.CultivationPlan

	for rows.Next() {
		var p cultivation.CultivationPlan
		var stepsBytes []byte

		err := rows.Scan(
			&p.ID,
			&p.Version,
			&p.Name,
			&p.CropKey,
			&stepsBytes,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(stepsBytes, &p.Steps); err != nil {
			return nil, err
		}

		result = append(result, &p)
	}

	return result, nil
}
