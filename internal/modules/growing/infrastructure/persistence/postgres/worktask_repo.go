package postgres

import (
	"context"
	"samurenkoroma/services/internal/infrastructure/persistence"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/worktask"
)

type worktaskRepo struct {
	db persistence.DBTX
}

func NewWorktaskRepository(db persistence.DBTX) worktask.Repository {
	return &worktaskRepo{db: db}
}

func (r *worktaskRepo) SaveMany(tasks []worktask.Task) error {
	ctx := context.Background()

	query := `
		INSERT INTO tasks (id, crop_plan_id, type, scheduled_date, status)
		VALUES ($1, $2, $3, $4, $5)
	`

	for _, t := range tasks {
		_, err := r.db.ExecContext(ctx, query,
			t.ID,
			t.CropPlanID,
			t.Type,
			t.ScheduledDate,
			t.Status,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
