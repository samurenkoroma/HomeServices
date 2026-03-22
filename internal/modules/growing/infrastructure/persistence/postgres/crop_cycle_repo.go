package postgres

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/modules/growing/domain/cropcycle"
	"time"
)

type SqlCropCycleRepository struct {
	tx *sql.Tx
}

func NewSqlCropCycleRepository(tx *sql.Tx) *SqlCropCycleRepository {
	return &SqlCropCycleRepository{tx: tx}
}
func (r *SqlCropCycleRepository) Save(
	ctx context.Context,
	c *cropcycle.CropCycle,
) error {

	if c.Version() == 0 {
		// INSERT
		_, err := r.tx.ExecContext(ctx, `
			INSERT INTO crop_cycles (
				id, plan_id, plan_version,
				facility_id, bed_id,
				status, started_at, finished_at,
				version
			)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		`,
			c.ID(),
			c.PlanID(),
			c.PlanVersion(),
			c.FacilityID(),
			c.BedID(),
			c.Status(),
			c.StartedAt(),
			c.FinishedAt(),
			1,
		)
		return err
	}

	// UPDATE with optimistic locking
	res, err := r.tx.ExecContext(ctx, `
		UPDATE crop_cycles
		SET status = $1,
		    started_at = $2,
		    finished_at = $3,
		    version = version + 1
		WHERE id = $4
		  AND version = $5
	`,
		c.Status(),
		c.StartedAt(),
		c.FinishedAt(),
		c.ID(),
		c.Version(),
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrConcurrentModification
	}

	return nil
}

func (r *SqlCropCycleRepository) ByID(ctx context.Context, id string) (*cropcycle2.CropCycle, error) {

	row := r.tx.QueryRowContext(ctx, `
		SELECT id, plan_id, plan_version,
		       facility_id, bed_id,
		       status, started_at, finished_at,
		       version
		FROM crop_cycles
		WHERE id = $1
	`, id)

	var (
		cid         string
		planID      string
		facilityID  string
		bedID       string
		status      string
		planVersion int
		startedAt   time.Time
		finishedAt  time.Time
		version     int
	)

	err := row.Scan(
		&cid,
		&planID,
		&planVersion,
		&facilityID,
		&bedID,
		&status,
		&startedAt,
		&finishedAt,
		&version,
	)
	if err != nil {
		return nil, err
	}
	c := cropcycle2.Rehydrate(
		cid,
		planID,
		facilityID,
		bedID,
		cropcycle2.Status(status),
		planVersion,
	)

	return c, nil
}
