package cropplan

import (
	"context"
	"time"
)

// Repository интерфейс для работы с планами
type Repository interface {
	// Базовые операции
	Save(ctx context.Context, plan *CropPlan) error
	//FindByID(ctx context.Context, id string) (*CropPlan, error)
	//Delete(ctx context.Context, id string) error
	//
	//Запросы по полям
	//FindByArea(ctx context.Context, areaId string) ([]*CropPlan, error)
	//FindByVariety(ctx context.Context, varietyID string) ([]*CropPlan, error)
	//FindByStatus(ctx context.Context, status Status) ([]*CropPlan, error)
	//FindByAssignedTo(ctx context.Context, userID string) ([]*CropPlan, error)
	//
	//Поиск по датам
	//FindActiveByDate(ctx context.Context, date time.Time) ([]*CropPlan, error)
	//FindBySeason(ctx context.Context, bedID string, seasonStart, seasonEnd time.Time) ([]*CropPlan, error)
	//
	//Статистика
	//GetStatistics(ctx context.Context, filter StatisticsFilter) (*Statistics, error)
}

// StatisticsFilter фильтр для статистики
type StatisticsFilter struct {
	BedID     string
	VarietyID string
	DateFrom  time.Time
	DateTo    time.Time
}

// Statistics статистика по планам
type Statistics struct {
	TotalPlans     int     `json:"total_plans"`
	ActivePlans    int     `json:"active_plans"`
	CompletedPlans int     `json:"completed_plans"`
	TotalHarvestKg float64 `json:"total_harvest_kg"`
	AvgYieldPerM2  float64 `json:"avg_yield_per_m2"`
}
