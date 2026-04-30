package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"time"
)

// CropPlanRepository реализация репозитория для PostgreSQL
type cropPlanRepository struct {
	tx *sql.Tx
}

// NewCropPlanRepository создает новый репозиторий в транзакции
func NewCropPlanRepository(tx *sql.Tx) cropplan.Repository {
	return &cropPlanRepository{tx: tx}
}

// cropPlanRow структура для сканирования из БД
type cropPlanRow struct {
	ID           string
	Name         string
	AreaId       string
	VarietyID    string
	Season       string
	AssignedTo   string
	Status       string
	Stages       []byte // JSONB
	Metadata     []byte // JSONB
	PlantingDate time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	StartedAt    *time.Time
	CompletedAt  *time.Time
}

// ========== БАЗОВЫЕ ОПЕРАЦИИ ==========

// Save сохраняет новый план
func (r *cropPlanRepository) Save(ctx context.Context, plan *cropplan.CropPlan) error {
	// Сериализуем этапы в JSONB
	stagesJSON, err := json.Marshal(plan.Stages())
	if err != nil {
		return fmt.Errorf("failed to marshal stages: %w", err)
	}
	metadataJSON, err := json.Marshal(plan.Metadata())
	if err != nil {
		return fmt.Errorf("failed to marshal stages: %w", err)
	}

	query := `
        INSERT INTO growing_crop_plans (
            id, name, 
			area_id, variety_id, season_id,  assigned_to, 
			status, metadata, stages, 
			planting_date, created_at, updated_at, started_at, completed_at
        ) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13,$14)
    `

	_, err = r.tx.ExecContext(ctx, query,
		plan.ID(), plan.Name(),
		plan.Area().GetId(),
		plan.Variety().GetId(), plan.Season().GetId(), plan.AssignedTo(),
		string(plan.Status()), metadataJSON, stagesJSON,
		plan.PlantingDate(), plan.CreatedAt(), plan.UpdatedAt(), plan.StartedAt(), plan.CompletedAt(),
	)

	if err != nil {
		return fmt.Errorf("failed to save crop plan: %w", err)
	}

	return nil
}

// FindByID находит план по ID
func (r *cropPlanRepository) FindByID(ctx context.Context, id string) (*cropplan.CropPlan, error) {
	//_ := `
	//    SELECT
	//        id, name,
	//        area_id, variety_id, season_id,  assigned_to,
	//        status, metadata, stages,
	//        planting_date, created_at, updated_at, started_at, completed_at
	//    FROM growing_crop_plans
	//    WHERE id = $1
	//`

	//var row cropPlanRow
	//err := r.tx.QueryRowContext(ctx, query, id).Scan(
	//	&row.ID, &row.Name,
	//	&row.AreaId, &row.VarietyID,
	//	&row.CropName,
	//	&row.Status,
	//	&row.Stages,
	//	&row.SeasonStart,
	//	&row.SeasonEnd,
	//	&row.PlantingDate,
	//	&row.SeedsPlanted,
	//	&row.ExpectedYield,
	//	&row.HarvestKg,
	//	&row.CreatedAt,
	//	&row.UpdatedAt,
	//	&row.StartedAt,
	//	&row.CompletedAt,
	//	&row.Latitude,
	//	&row.Longitude,
	//	&row.AssignedTo,
	//	&row.AssignedName,
	//)
	//
	//if err != nil {
	//	if errors.Is(err, sql.ErrNoRows) {
	//		return nil, cropplan.ErrPlanNotFound
	//	}
	//	return nil, fmt.Errorf("failed to find crop plan: %w", err)
	//}

	//return r.mapRowToPlan(row)

	return nil, nil
}

// // Delete удаляет план
//
//	func (r *CropPlanRepository) Delete(ctx context.Context, id string) error {
//		query := `DELETE FROM crop_plans WHERE id = $1`
//
//		result, err := r.tx.ExecContext(ctx, query, id)
//		if err != nil {
//			return fmt.Errorf("failed to delete crop plan: %w", err)
//		}
//
//		rows, err := result.RowsAffected()
//		if err != nil {
//			return err
//		}
//		if rows == 0 {
//			return cropplan.ErrPlanNotFound
//		}
//
//		return nil
//	}
//
// // ========== ЗАПРОСЫ ПО ПОЛЯМ ==========
//
// // FindByArea находит все планы на грядке (area = bed)
func (r *cropPlanRepository) FindByObject(ctx context.Context, areaId string) ([]*cropplan.CropPlan, error) {
	//query := `
	//    SELECT
	//        id, bed_id, name, variety_id, crop_name, status, stages,
	//        season_start, season_end, planting_date, seeds_planted,
	//        expected_yield, harvest_kg, created_at, updated_at,
	//        started_at, completed_at, latitude, longitude,
	//        assigned_to, assigned_name
	//    FROM crop_plans
	//    WHERE bed_id = $1
	//    ORDER BY created_at DESC
	//`
	//
	//rows, err := r.tx.QueryContext(ctx, query, areaId)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to find plans by area: %w", err)
	//}
	//defer rows.Close()
	//
	//return r.scanRows(rows)
	return nil, nil
}

//
//// FindByVariety находит все планы по сорту
//func (r *CropPlanRepository) FindByVariety(ctx context.Context, varietyID string) ([]*cropplan.CropPlan, error) {
//	query := `
//        SELECT
//            id, bed_id, name, variety_id, crop_name, status, stages,
//            season_start, season_end, planting_date, seeds_planted,
//            expected_yield, harvest_kg, created_at, updated_at,
//            started_at, completed_at, latitude, longitude,
//            assigned_to, assigned_name
//        FROM crop_plans
//        WHERE variety_id = $1
//        ORDER BY created_at DESC
//    `
//
//	rows, err := r.tx.QueryContext(ctx, query, varietyID)
//	if err != nil {
//		return nil, fmt.Errorf("failed to find plans by variety: %w", err)
//	}
//	defer rows.Close()
//
//	return r.scanRows(rows)
//}
//
//// FindByStatus находит все планы по статусу
//func (r *CropPlanRepository) FindByStatus(ctx context.Context, status cropplan.Status) ([]*cropplan.CropPlan, error) {
//	query := `
//        SELECT
//            id, bed_id, name, variety_id, crop_name, status, stages,
//            season_start, season_end, planting_date, seeds_planted,
//            expected_yield, harvest_kg, created_at, updated_at,
//            started_at, completed_at, latitude, longitude,
//            assigned_to, assigned_name
//        FROM crop_plans
//        WHERE status = $1
//        ORDER BY created_at DESC
//    `
//
//	rows, err := r.tx.QueryContext(ctx, query, string(status))
//	if err != nil {
//		return nil, fmt.Errorf("failed to find plans by status: %w", err)
//	}
//	defer rows.Close()
//
//	return r.scanRows(rows)
//}
//
//// FindByAssignedTo находит планы, назначенные агроному
//func (r *CropPlanRepository) FindByAssignedTo(ctx context.Context, userID string) ([]*cropplan.CropPlan, error) {
//	query := `
//        SELECT
//            id, bed_id, name, variety_id, crop_name, status, stages,
//            season_start, season_end, planting_date, seeds_planted,
//            expected_yield, harvest_kg, created_at, updated_at,
//            started_at, completed_at, latitude, longitude,
//            assigned_to, assigned_name
//        FROM crop_plans
//        WHERE assigned_to = $1
//        ORDER BY created_at DESC
//    `
//
//	rows, err := r.tx.QueryContext(ctx, query, userID)
//	if err != nil {
//		return nil, fmt.Errorf("failed to find plans by assigned user: %w", err)
//	}
//	defer rows.Close()
//
//	return r.scanRows(rows)
//}
//
//// ========== ПОИСК ПО ДАТАМ ==========
//
//// FindActiveByDate находит активные планы на указанную дату
//func (r *CropPlanRepository) FindActiveByDate(ctx context.Context, date time.Time) ([]*cropplan.CropPlan, error) {
//	query := `
//        SELECT
//            id, bed_id, name, variety_id, crop_name, status, stages,
//            season_start, season_end, planting_date, seeds_planted,
//            expected_yield, harvest_kg, created_at, updated_at,
//            started_at, completed_at, latitude, longitude,
//            assigned_to, assigned_name
//        FROM crop_plans
//        WHERE status = 'active'
//            AND season_start <= $1
//            AND season_end >= $1
//        ORDER BY created_at DESC
//    `
//
//	rows, err := r.tx.QueryContext(ctx, query, date)
//	if err != nil {
//		return nil, fmt.Errorf("failed to find active plans by date: %w", err)
//	}
//	defer rows.Close()
//
//	return r.scanRows(rows)
//}
//
//// FindBySeason находит планы на грядке за сезон
//func (r *CropPlanRepository) FindBySeason(ctx context.Context, bedID string, seasonStart, seasonEnd time.Time) ([]*cropplan.CropPlan, error) {
//	query := `
//        SELECT
//            id, bed_id, name, variety_id, crop_name, status, stages,
//            season_start, season_end, planting_date, seeds_planted,
//            expected_yield, harvest_kg, created_at, updated_at,
//            started_at, completed_at, latitude, longitude,
//            assigned_to, assigned_name
//        FROM crop_plans
//        WHERE bed_id = $1
//            AND season_start <= $3
//            AND season_end >= $2
//        ORDER BY created_at DESC
//    `
//
//	rows, err := r.tx.QueryContext(ctx, query, bedID, seasonStart, seasonEnd)
//	if err != nil {
//		return nil, fmt.Errorf("failed to find plans by season: %w", err)
//	}
//	defer rows.Close()
//
//	return r.scanRows(rows)
//}
//
//// ========== СТАТИСТИКА ==========
//
//// GetStatistics возвращает статистику по планам
//func (r *CropPlanRepository) GetStatistics(ctx context.Context, filter cropplan.StatisticsFilter) (*cropplan.Statistics, error) {
//	query := `
//        SELECT
//            COUNT(*) as total_plans,
//            COUNT(CASE WHEN status = 'active' THEN 1 END) as active_plans,
//            COUNT(CASE WHEN status = 'completed' THEN 1 END) as completed_plans,
//            COALESCE(SUM(harvest_kg), 0) as total_harvest_kg,
//            COALESCE(AVG(CASE WHEN status = 'completed' THEN harvest_kg / NULLIF(expected_yield, 0) * 100 END), 0) as avg_yield_percent
//        FROM crop_plans
//        WHERE 1=1
//    `
//
//	var args []interface{}
//	argIndex := 1
//
//	if filter.BedID != "" {
//		query += fmt.Sprintf(" AND bed_id = $%d", argIndex)
//		args = append(args, filter.BedID)
//		argIndex++
//	}
//
//	if filter.VarietyID != "" {
//		query += fmt.Sprintf(" AND variety_id = $%d", argIndex)
//		args = append(args, filter.VarietyID)
//		argIndex++
//	}
//
//	if !filter.DateFrom.IsZero() {
//		query += fmt.Sprintf(" AND created_at >= $%d", argIndex)
//		args = append(args, filter.DateFrom)
//		argIndex++
//	}
//
//	if !filter.DateTo.IsZero() {
//		query += fmt.Sprintf(" AND created_at <= $%d", argIndex)
//		args = append(args, filter.DateTo)
//		argIndex++
//	}
//
//	var stats cropplan.Statistics
//	var avgYieldPercent float64
//
//	err := r.tx.QueryRowContext(ctx, query, args...).Scan(
//		&stats.TotalPlans,
//		&stats.ActivePlans,
//		&stats.CompletedPlans,
//		&stats.TotalHarvestKg,
//		&avgYieldPercent,
//	)
//	if err != nil {
//		return nil, fmt.Errorf("failed to get statistics: %w", err)
//	}
//
//	stats.AvgYieldPerM2 = avgYieldPercent / 100
//	return &stats, nil
//}
//
//// ========== ВСПОМОГАТЕЛЬНЫЕ МЕТОДЫ ==========
//
//// scanRows сканирует строки в список планов
//func (r *CropPlanRepository) scanRows(rows *sql.Rows) ([]*cropplan.CropPlan, error) {
//	var plans []*cropplan.CropPlan
//
//	for rows.Next() {
//		var row cropPlanRow
//		err := rows.Scan(
//			&row.ID,
//			&row.BedID,
//			&row.Name,
//			&row.VarietyID,
//			&row.CropName,
//			&row.Status,
//			&row.Stages,
//			&row.SeasonStart,
//			&row.SeasonEnd,
//			&row.PlantingDate,
//			&row.SeedsPlanted,
//			&row.ExpectedYield,
//			&row.HarvestKg,
//			&row.CreatedAt,
//			&row.UpdatedAt,
//			&row.StartedAt,
//			&row.CompletedAt,
//			&row.Latitude,
//			&row.Longitude,
//			&row.AssignedTo,
//			&row.AssignedName,
//		)
//		if err != nil {
//			return nil, fmt.Errorf("failed to scan row: %w", err)
//		}
//
//		plan, err := r.mapRowToPlan(row)
//		if err != nil {
//			return nil, err
//		}
//		plans = append(plans, plan)
//	}
//
//	if err := rows.Err(); err != nil {
//		return nil, fmt.Errorf("rows iteration error: %w", err)
//	}
//
//	return plans, nil
//}
//
//// mapRowToPlan преобразует строку БД в доменный объект
//func (r *CropPlanRepository) mapRowToPlan(row cropPlanRow) (*cropplan.CropPlan, error) {
//	// Десериализуем этапы из JSONB
//	var stages []cropplan.Stage
//	if len(row.Stages) > 0 {
//		if err := json.Unmarshal(row.Stages, &stages); err != nil {
//			return nil, fmt.Errorf("failed to unmarshal stages: %w", err)
//		}
//	}
//
//	// Создаем план через конструктор
//	plan, err := cropplan.NewCropPlan(
//		row.ID,
//		row.BedID,
//		row.Name,
//		row.VarietyID,
//		row.VarietyID, // varietyName нужно будет получить из каталога
//		row.CropName,
//		row.SeasonStart,
//		row.SeasonEnd,
//		row.PlantingDate,
//		row.Latitude,
//		row.Longitude,
//		row.AssignedTo,
//		row.AssignedName,
//	)
//	if err != nil {
//		return nil, fmt.Errorf("failed to create plan from row: %w", err)
//	}
//
//	// Устанавливаем остальные поля (потребуется добавить setters в агрегат)
//	plan.SetSeedsPlanted(row.SeedsPlanted)
//	plan.SetExpectedYield(row.ExpectedYield)
//	plan.SetHarvestKg(row.HarvestKg)
//	plan.SetCreatedAt(row.CreatedAt)
//	plan.SetUpdatedAt(row.UpdatedAt)
//	plan.SetStartedAt(row.StartedAt)
//	plan.SetCompletedAt(row.CompletedAt)
//	plan.SetStatus(cropplan.Status(row.Status))
//
//	// Восстанавливаем этапы (потребуется метод RestoreStages)
//	for _, stage := range stages {
//		plan.AddStage(stage)
//	}
//
//	return plan, nil
//}
