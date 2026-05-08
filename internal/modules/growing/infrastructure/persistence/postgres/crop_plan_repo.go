package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"samurenkoroma/services/internal/infrastructure/persistence"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"strings"
)

type cropPlanRepo struct {
	db persistence.DBTX
}

func NewCropPlanRepository(db persistence.DBTX) cropplan.Repository {
	return &cropPlanRepo{db: db}
}

func (r *cropPlanRepo) Save(ctx context.Context, p *cropplan.Plan) error {
	snapshotBytes, err := json.Marshal(p.Snapshot)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO growing_crop_plans (
		id, crop_key, variety_id,
		season_id, area_id, owner_id,
		start_date, status,
		cultivation_plan_id, cultivation_plan_version,
		cultivation_plan_snapshot,
		created_at, updated_at
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
	ON CONFLICT (id) DO UPDATE SET
		crop_key = EXCLUDED.crop_key,
		variety_id = EXCLUDED.variety_id,
		season_id = EXCLUDED.season_id,
		area_id = EXCLUDED.area_id,
		start_date = EXCLUDED.start_date,
		status = EXCLUDED.status,
		cultivation_plan_id = EXCLUDED.cultivation_plan_id,
		cultivation_plan_version = EXCLUDED.cultivation_plan_version,
		cultivation_plan_snapshot = EXCLUDED.cultivation_plan_snapshot,
		updated_at = EXCLUDED.updated_at
	`

	_, err = r.db.ExecContext(ctx, query,
		p.ID,
		p.CropKey,
		p.VarietyID,
		p.SeasonID,
		p.AreaID,
		p.Organization,
		p.StartDate,
		p.Status,
		p.CultivationPlanID,
		p.CultivationPlanVersion,
		snapshotBytes,
		p.CreatedAt,
		p.UpdatedAt,
	)

	return err
}

func (r *cropPlanRepo) GetByID(ctx context.Context, id string) (*cropplan.Plan, error) {
	query := `
	SELECT id, crop_key, variety_id,
	       season_id, area_id, owner_id,
	       start_date, status,
	       cultivation_plan_id, cultivation_plan_version,
	       cultivation_plan_snapshot,
	       created_at, updated_at
	FROM growing_crop_plans
	WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var p cropplan.Plan
	var snapshotBytes []byte

	err := row.Scan(
		&p.ID,
		&p.CropKey,
		&p.VarietyID,
		&p.SeasonID,
		&p.AreaID,
		&p.Organization,
		&p.StartDate,
		&p.Status,
		&p.CultivationPlanID,
		&p.CultivationPlanVersion,
		&snapshotBytes,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(snapshotBytes, &p.Snapshot); err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *cropPlanRepo) All(ctx context.Context, f cropplan.Filter) ([]*cropplan.Plan, error) {
	base := `
	SELECT id, crop_key, variety_id,
	       season_id, area_id, owner_id,
	       start_date, status,
	       cultivation_plan_id, cultivation_plan_version,
	       cultivation_plan_snapshot,
	       created_at, updated_at
	FROM growing_crop_plans
	`

	var conditions []string
	var args []interface{}
	argIdx := 1

	// 🔥 multi-tenant защита
	if f.OwnerID == nil {
		return nil, fmt.Errorf("owner_id is required")
	}

	conditions = append(conditions, fmt.Sprintf("owner_id = $%d", argIdx))
	args = append(args, *f.OwnerID)
	argIdx++

	if f.SeasonID != nil {
		conditions = append(conditions, fmt.Sprintf("season_id = $%d", argIdx))
		args = append(args, *f.SeasonID)
		argIdx++
	}

	if f.AreaID != nil {
		conditions = append(conditions, fmt.Sprintf("area_id = $%d", argIdx))
		args = append(args, *f.AreaID)
		argIdx++
	}

	if f.Status != nil {
		conditions = append(conditions, fmt.Sprintf("status = $%d", argIdx))
		args = append(args, *f.Status)
		argIdx++
	}

	query := base + " WHERE " + strings.Join(conditions, " AND ")
	query += " ORDER BY created_at DESC"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*cropplan.Plan{}

	for rows.Next() {
		var p cropplan.Plan
		var snapshotBytes []byte

		err := rows.Scan(
			&p.ID,
			&p.CropKey,
			&p.VarietyID,
			&p.SeasonID,
			&p.AreaID,
			&p.Organization,
			&p.StartDate,
			&p.Status,
			&p.CultivationPlanID,
			&p.CultivationPlanVersion,
			&snapshotBytes,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(snapshotBytes, &p.Snapshot); err != nil {
			return nil, err
		}

		result = append(result, &p)
	}

	return result, nil
}
