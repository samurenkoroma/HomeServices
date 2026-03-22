package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"samurenkoroma/services/internal/modules/growing/domain/cropcycle"
)

type cropCycleRepository struct {
	tx *sql.Tx
}

func NewCropCycleRepository(tx *sql.Tx) cropcycle.Repository {
	return &cropCycleRepository{tx: tx}
}

// Save сохраняет цикл (вставка или обновление с оптимистичной блокировкой)
func (r *cropCycleRepository) Save(ctx context.Context, cycle *cropcycle.CropCycle) error {
	// Для новых циклов (version = 0) делаем INSERT
	if cycle.GetVersion() == 0 {
		return r.insert(ctx, cycle)
	}
	// Для существующих — UPDATE с проверкой версии
	return r.update(ctx, cycle)
}

// insert вставляет новый цикл
func (r *cropCycleRepository) insert(ctx context.Context, cycle *cropcycle.CropCycle) error {
	query := `
        INSERT INTO crop_cycles (
            id, template_id, area_id, season_id, crop_plan_id, crop_plan_version,
            status, started_at, finished_at,
            yield_actual, yield_estimated, yield_quality, yield_notes,
            version, created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
    `

	_, err := r.tx.ExecContext(ctx, query,
		string(cycle.GetID()),
		cycle.GetTemplateID(),
		cycle.GetAreaID(),
		cycle.GetSeasonID(),
		cycle.GetCropPlanID(),
		cycle.GetCropPlanVersion(),
		string(cycle.GetStatus()),
		cycle.GetStartedAt(),
		cycle.GetFinishedAt(),
		cycle.GetYieldActual(),
		cycle.GetYieldEstimated(),
		cycle.GetYieldQuality(),
		cycle.GetYieldNotes(),
		1, // начальная версия
		cycle.GetCreatedAt(),
		cycle.GetUpdatedAt(),
	)

	if err != nil {
		return fmt.Errorf("failed to insert crop cycle: %w", err)
	}

	// Сохраняем операции
	if err := r.saveOperations(ctx, cycle); err != nil {
		return err
	}

	return nil
}

// update обновляет цикл с оптимистичной блокировкой
func (r *cropCycleRepository) update(ctx context.Context, cycle *cropcycle.CropCycle) error {
	query := `
        UPDATE crop_cycles SET
            template_id = $2,
            area_id = $3,
            season_id = $4,
            crop_plan_id = $5,
            crop_plan_version = $6,
            status = $7,
            started_at = $8,
            finished_at = $9,
            yield_actual = $10,
            yield_estimated = $11,
            yield_quality = $12,
            yield_notes = $13,
            version = version + 1,
            updated_at = $14
        WHERE id = $1 AND version = $15
    `

	result, err := r.tx.ExecContext(ctx, query,
		string(cycle.GetID()),
		cycle.GetTemplateID(),
		cycle.GetAreaID(),
		cycle.GetSeasonID(),
		cycle.GetCropPlanID(),
		cycle.GetCropPlanVersion(),
		string(cycle.GetStatus()),
		cycle.GetStartedAt(),
		cycle.GetFinishedAt(),
		cycle.GetYieldActual(),
		cycle.GetYieldEstimated(),
		cycle.GetYieldQuality(),
		cycle.GetYieldNotes(),
		cycle.GetUpdatedAt(),
		cycle.GetVersion(),
	)

	if err != nil {
		return fmt.Errorf("failed to update crop cycle: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		return cropcycle.ErrConcurrentModification
	}

	// Сохраняем операции
	if err := r.saveOperations(ctx, cycle); err != nil {
		return err
	}

	return nil
}

// saveOperations сохраняет операции цикла
func (r *cropCycleRepository) saveOperations(ctx context.Context, cycle *cropcycle.CropCycle) error {
	// Удаляем старые операции
	_, err := r.tx.ExecContext(ctx, `DELETE FROM cycle_operations WHERE cycle_id = $1`, string(cycle.GetID()))
	if err != nil {
		return fmt.Errorf("failed to delete old operations: %w", err)
	}

	// Вставляем новые
	for _, op := range cycle.GetOperations() {
		_, err := r.tx.ExecContext(ctx, `
            INSERT INTO cycle_operations (
                id, cycle_id, type, description, amount, unit, performed_by, notes, created_at
            ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        `,
			op.ID,
			string(cycle.GetID()),
			string(op.Type),
			op.Description,
			op.Amount,
			op.Unit,
			op.PerformedBy,
			op.Notes,
			op.CreatedAt,
		)
		if err != nil {
			return fmt.Errorf("failed to insert operation: %w", err)
		}
	}

	return nil
}

// FindByID находит цикл по ID
func (r *cropCycleRepository) FindByID(ctx context.Context, id cropcycle.CycleID) (*cropcycle.CropCycle, error) {
	query := `
        SELECT 
            id, template_id, area_id, season_id, crop_plan_id, crop_plan_version,
            status, started_at, finished_at,
            yield_actual, yield_estimated, yield_quality, yield_notes,
            version, created_at, updated_at
        FROM crop_cycles
        WHERE id = $1
    `

	var (
		cycleID         string
		templateID      string
		areaID          string
		seasonID        string
		cropPlanID      string
		cropPlanVersion int
		status          string
		startedAt       sql.NullTime
		finishedAt      sql.NullTime
		yieldActual     sql.NullFloat64
		yieldEstimated  sql.NullFloat64
		yieldQuality    sql.NullString
		yieldNotes      sql.NullString
		version         int
		createdAt       time.Time
		updatedAt       time.Time
	)

	err := r.tx.QueryRowContext(ctx, query, string(id)).Scan(
		&cycleID, &templateID, &areaID, &seasonID, &cropPlanID, &cropPlanVersion,
		&status, &startedAt, &finishedAt,
		&yieldActual, &yieldEstimated, &yieldQuality, &yieldNotes,
		&version, &createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find crop cycle: %w", err)
	}

	// Загружаем операции
	operations, err := r.loadOperations(ctx, cycleID)
	if err != nil {
		return nil, err
	}

	// Создаём цикл
	cycle := cropcycle.NewCropCycle(
		templateID,
		areaID,
		seasonID,
		cropPlanID,
		cropPlanVersion,
	)

	var startedAtPtr, finishedAtPtr *time.Time
	if startedAt.Valid {
		startedAtPtr = &startedAt.Time
	}
	if finishedAt.Valid {
		finishedAtPtr = &finishedAt.Time
	}

	var yieldPtr *cropcycle.Yield
	if yieldActual.Valid || yieldEstimated.Valid {
		yieldPtr = &cropcycle.Yield{
			ActualWeight:    yieldActual.Float64,
			EstimatedWeight: yieldEstimated.Float64,
			Quality:         yieldQuality.String,
			Notes:           yieldNotes.String,
		}
	}

	cycle.Rehydrate(
		cropcycle.CycleID(cycleID),
		templateID,
		areaID,
		seasonID,
		cropPlanID,
		cropPlanVersion,
		cropcycle.Status(status),
		startedAtPtr,
		finishedAtPtr,
		operations,
		yieldPtr,
		createdAt,
		updatedAt,
		version,
	)

	return cycle, nil
}

// FindByAreaID находит все циклы по месту выращивания
func (r *cropCycleRepository) FindByAreaID(ctx context.Context, areaID string) ([]*cropcycle.CropCycle, error) {
	query := `
        SELECT 
            id, template_id, area_id, season_id, crop_plan_id, crop_plan_version,
            status, started_at, finished_at,
            yield_actual, yield_estimated, yield_quality, yield_notes,
            version, created_at, updated_at
        FROM crop_cycles
        WHERE area_id = $1
        ORDER BY created_at DESC
    `

	rows, err := r.tx.QueryContext(ctx, query, areaID)
	if err != nil {
		return nil, fmt.Errorf("failed to query cycles by area: %w", err)
	}
	defer rows.Close()

	return r.scanCycles(ctx, rows)
}

// FindBySeasonID находит все циклы по сезону
func (r *cropCycleRepository) FindBySeasonID(ctx context.Context, seasonID string) ([]*cropcycle.CropCycle, error) {
	query := `
        SELECT 
            id, template_id, area_id, season_id, crop_plan_id, crop_plan_version,
            status, started_at, finished_at,
            yield_actual, yield_estimated, yield_quality, yield_notes,
            version, created_at, updated_at
        FROM crop_cycles
        WHERE season_id = $1
        ORDER BY created_at DESC
    `

	rows, err := r.tx.QueryContext(ctx, query, seasonID)
	if err != nil {
		return nil, fmt.Errorf("failed to query cycles by season: %w", err)
	}
	defer rows.Close()

	return r.scanCycles(ctx, rows)
}

// FindByStatus находит все циклы по статусу
func (r *cropCycleRepository) FindByStatus(ctx context.Context, status cropcycle.Status) ([]*cropcycle.CropCycle, error) {
	query := `
        SELECT 
            id, template_id, area_id, season_id, crop_plan_id, crop_plan_version,
            status, started_at, finished_at,
            yield_actual, yield_estimated, yield_quality, yield_notes,
            version, created_at, updated_at
        FROM crop_cycles
        WHERE status = $1
        ORDER BY created_at DESC
    `

	rows, err := r.tx.QueryContext(ctx, query, string(status))
	if err != nil {
		return nil, fmt.Errorf("failed to query cycles by status: %w", err)
	}
	defer rows.Close()

	return r.scanCycles(ctx, rows)
}

// FindAll возвращает все циклы
func (r *cropCycleRepository) FindAll(ctx context.Context) ([]*cropcycle.CropCycle, error) {
	query := `
        SELECT 
            id, template_id, area_id, season_id, crop_plan_id, crop_plan_version,
            status, started_at, finished_at,
            yield_actual, yield_estimated, yield_quality, yield_notes,
            version, created_at, updated_at
        FROM crop_cycles
        ORDER BY created_at DESC
    `

	rows, err := r.tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query all cycles: %w", err)
	}
	defer rows.Close()

	return r.scanCycles(ctx, rows)
}

// Delete удаляет цикл
func (r *cropCycleRepository) Delete(ctx context.Context, id cropcycle.CycleID) error {
	// Операции удалятся каскадно
	_, err := r.tx.ExecContext(ctx, `DELETE FROM crop_cycles WHERE id = $1`, string(id))
	if err != nil {
		return fmt.Errorf("failed to delete crop cycle: %w", err)
	}
	return nil
}

// FindActiveByArea находит активный цикл на месте
func (r *cropCycleRepository) FindActiveByArea(ctx context.Context, areaID string) (*cropcycle.CropCycle, error) {
	query := `
        SELECT 
            id, template_id, area_id, season_id, crop_plan_id, crop_plan_version,
            status, started_at, finished_at,
            yield_actual, yield_estimated, yield_quality, yield_notes,
            version, created_at, updated_at
        FROM crop_cycles
        WHERE area_id = $1 AND status IN ('active', 'growing')
        LIMIT 1
    `

	var (
		cycleID         string
		templateID      string
		areaID2         string
		seasonID        string
		cropPlanID      string
		cropPlanVersion int
		status          string
		startedAt       sql.NullTime
		finishedAt      sql.NullTime
		yieldActual     sql.NullFloat64
		yieldEstimated  sql.NullFloat64
		yieldQuality    sql.NullString
		yieldNotes      sql.NullString
		version         int
		createdAt       time.Time
		updatedAt       time.Time
	)

	err := r.tx.QueryRowContext(ctx, query, areaID).Scan(
		&cycleID, &templateID, &areaID2, &seasonID, &cropPlanID, &cropPlanVersion,
		&status, &startedAt, &finishedAt,
		&yieldActual, &yieldEstimated, &yieldQuality, &yieldNotes,
		&version, &createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find active cycle by area: %w", err)
	}

	operations, err := r.loadOperations(ctx, cycleID)
	if err != nil {
		return nil, err
	}

	cycle := cropcycle.NewCropCycle(
		templateID,
		areaID2,
		seasonID,
		cropPlanID,
		cropPlanVersion,
	)

	var startedAtPtr, finishedAtPtr *time.Time
	if startedAt.Valid {
		startedAtPtr = &startedAt.Time
	}
	if finishedAt.Valid {
		finishedAtPtr = &finishedAt.Time
	}

	var yieldPtr *cropcycle.Yield
	if yieldActual.Valid || yieldEstimated.Valid {
		yieldPtr = &cropcycle.Yield{
			ActualWeight:    yieldActual.Float64,
			EstimatedWeight: yieldEstimated.Float64,
			Quality:         yieldQuality.String,
			Notes:           yieldNotes.String,
		}
	}

	cycle.Rehydrate(
		cropcycle.CycleID(cycleID),
		templateID,
		areaID2,
		seasonID,
		cropPlanID,
		cropPlanVersion,
		cropcycle.Status(status),
		startedAtPtr,
		finishedAtPtr,
		operations,
		yieldPtr,
		createdAt,
		updatedAt,
		version,
	)

	return cycle, nil
}

// FindActiveBySeason находит все активные циклы в сезоне
func (r *cropCycleRepository) FindActiveBySeason(ctx context.Context, seasonID string) ([]*cropcycle.CropCycle, error) {
	query := `
        SELECT 
            id, template_id, area_id, season_id, crop_plan_id, crop_plan_version,
            status, started_at, finished_at,
            yield_actual, yield_estimated, yield_quality, yield_notes,
            version, created_at, updated_at
        FROM crop_cycles
        WHERE season_id = $1 AND status IN ('active', 'growing')
        ORDER BY created_at DESC
    `

	rows, err := r.tx.QueryContext(ctx, query, seasonID)
	if err != nil {
		return nil, fmt.Errorf("failed to query active cycles by season: %w", err)
	}
	defer rows.Close()

	return r.scanCycles(ctx, rows)
}

// Exists проверяет существование цикла
func (r *cropCycleRepository) Exists(ctx context.Context, id cropcycle.CycleID) (bool, error) {
	var exists bool
	err := r.tx.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM crop_cycles WHERE id = $1)`, string(id)).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check existence: %w", err)
	}
	return exists, nil
}

// ExistsActiveForArea проверяет наличие активного цикла на месте
func (r *cropCycleRepository) ExistsActiveForArea(ctx context.Context, areaID string) (bool, error) {
	var exists bool
	err := r.tx.QueryRowContext(ctx, `
        SELECT EXISTS(
            SELECT 1 FROM crop_cycles 
            WHERE area_id = $1 AND status IN ('active', 'growing')
        )
    `, areaID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check active cycle: %w", err)
	}
	return exists, nil
}

// loadOperations загружает операции цикла
func (r *cropCycleRepository) loadOperations(ctx context.Context, cycleID string) ([]cropcycle.Operation, error) {
	rows, err := r.tx.QueryContext(ctx, `
        SELECT id, type, description, amount, unit, performed_by, notes, created_at
        FROM cycle_operations
        WHERE cycle_id = $1
        ORDER BY created_at
    `, cycleID)
	if err != nil {
		return nil, fmt.Errorf("failed to query operations: %w", err)
	}
	defer rows.Close()

	var operations []cropcycle.Operation
	for rows.Next() {
		var op cropcycle.Operation
		err := rows.Scan(
			&op.ID,
			&op.Type,
			&op.Description,
			&op.Amount,
			&op.Unit,
			&op.PerformedBy,
			&op.Notes,
			&op.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan operation: %w", err)
		}
		operations = append(operations, op)
	}

	return operations, nil
}

// scanCycles сканирует строки и возвращает список циклов
func (r *cropCycleRepository) scanCycles(ctx context.Context, rows *sql.Rows) ([]*cropcycle.CropCycle, error) {
	var cycles []*cropcycle.CropCycle

	for rows.Next() {
		var (
			cycleID         string
			templateID      string
			areaID          string
			seasonID        string
			cropPlanID      string
			cropPlanVersion int
			status          string
			startedAt       sql.NullTime
			finishedAt      sql.NullTime
			yieldActual     sql.NullFloat64
			yieldEstimated  sql.NullFloat64
			yieldQuality    sql.NullString
			yieldNotes      sql.NullString
			version         int
			createdAt       time.Time
			updatedAt       time.Time
		)

		err := rows.Scan(
			&cycleID, &templateID, &areaID, &seasonID, &cropPlanID, &cropPlanVersion,
			&status, &startedAt, &finishedAt,
			&yieldActual, &yieldEstimated, &yieldQuality, &yieldNotes,
			&version, &createdAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cycle: %w", err)
		}

		operations, err := r.loadOperations(ctx, cycleID)
		if err != nil {
			return nil, err
		}

		cycle := cropcycle.NewCropCycle(
			templateID,
			areaID,
			seasonID,
			cropPlanID,
			cropPlanVersion,
		)

		var startedAtPtr, finishedAtPtr *time.Time
		if startedAt.Valid {
			startedAtPtr = &startedAt.Time
		}
		if finishedAt.Valid {
			finishedAtPtr = &finishedAt.Time
		}

		var yieldPtr *cropcycle.Yield
		if yieldActual.Valid || yieldEstimated.Valid {
			yieldPtr = &cropcycle.Yield{
				ActualWeight:    yieldActual.Float64,
				EstimatedWeight: yieldEstimated.Float64,
				Quality:         yieldQuality.String,
				Notes:           yieldNotes.String,
			}
		}

		cycle.Rehydrate(
			cropcycle.CycleID(cycleID),
			templateID,
			areaID,
			seasonID,
			cropPlanID,
			cropPlanVersion,
			cropcycle.Status(status),
			startedAtPtr,
			finishedAtPtr,
			operations,
			yieldPtr,
			createdAt,
			updatedAt,
			version,
		)

		cycles = append(cycles, cycle)
	}

	return cycles, nil
}
