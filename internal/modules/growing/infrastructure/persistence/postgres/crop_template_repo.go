package postgres

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/modules/growing/domain/croptemplate"
	"time"
)

type SqlCropTemplateRepository struct {
	tx *sql.Tx
}

func NewSqlCropTemplateRepository(tx *sql.Tx) *SqlCropTemplateRepository {
	return &SqlCropTemplateRepository{tx: tx}
}
func (r *SqlCropTemplateRepository) Save(
	ctx context.Context,
	t *croptemplate.CropTemplate,
) error {

	_, err := r.tx.ExecContext(ctx, `
		INSERT INTO crop_templates (plan_id, version, active, updated_at)
		VALUES ($1,$2,$3,$4)
		ON CONFLICT (plan_id)
		DO UPDATE SET
		    version = EXCLUDED.version,
		    active = EXCLUDED.active,
		    updated_at = EXCLUDED.updated_at
	`,
		t.PlanID(),
		t.Version(),
		t.Active(),
		time.Now(),
	)

	return err
}

func (r *SqlCropTemplateRepository) ByPlanID(
	ctx context.Context,
	planID string,
) (*croptemplate.CropTemplate, error) {

	row := r.tx.QueryRowContext(ctx, `
		SELECT plan_id, version, active
		FROM crop_templates
		WHERE plan_id = $1
	`, planID)

	var (
		id      string
		version int
		active  bool
	)

	err := row.Scan(&id, &version, &active)
	if err != nil {
		return nil, err
	}

	template := croptemplate.New(id, version)
	if !active {
		template.Disable()
	}

	return template, nil
}
