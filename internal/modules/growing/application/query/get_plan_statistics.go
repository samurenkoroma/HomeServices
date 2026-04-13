package query

import (
	"context"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/task"
	"time"
)

type getPlanStatisticsHandler struct {
	PlanRepo cropplan.Repository
	TaskRepo task.Repository
}

func NewGetPlanStatisticsHandler(PlanRepo cropplan.Repository, TaskRepo task.Repository) query.Handler {
	return &getPlanStatisticsHandler{
		PlanRepo: PlanRepo,
		TaskRepo: TaskRepo,
	}
}

// GetPlanStatisticsQuery параметры запроса
type GetPlanStatisticsQuery struct {
	PlanID string `json:"plan_id"`
}

func (h *getPlanStatisticsHandler) Name() string {
	return "GetPlanStatistics"
}

// GetPlanStatisticsResponse ответ со статистикой
type GetPlanStatisticsResponse struct {
	PlanID      string `json:"plan_id"`
	PlanName    string `json:"plan_name"`
	VarietyName string `json:"variety_name"`
	Status      string `json:"status"`

	// Прогресс
	Progress         float64 `json:"progress"` // % выполнения
	CompletedStages  int     `json:"completed_stages"`
	TotalStages      int     `json:"total_stages"`
	PendingStages    int     `json:"pending_stages"`
	InProgressStages int     `json:"in_progress_stages"`
	SkippedStages    int     `json:"skipped_stages"`

	// Временные показатели
	DaysSincePlanting int `json:"days_since_planting"`
	DaysRemaining     int `json:"days_remaining"`  // до конца сезона
	DaysToHarvest     int `json:"days_to_harvest"` // прогноз

	// Урожайность
	ExpectedYield   float64 `json:"expected_yield"`   // кг
	ActualYield     float64 `json:"actual_yield"`     // кг
	YieldEfficiency float64 `json:"yield_efficiency"` // % от ожидаемого

	// Задания
	TotalTasks     int `json:"total_tasks"`
	CompletedTasks int `json:"completed_tasks"`
	PendingTasks   int `json:"pending_tasks"`
	OverdueTasks   int `json:"overdue_tasks"`

	// Эффективность работ
	AvgTaskCompletionDays float64 `json:"avg_task_completion_days"`

	// Отклонения
	DeviationDays   int    `json:"deviation_days"` // отклонение от графика
	DeviationReason string `json:"deviation_reason,omitempty"`

	// Рекомендации
	Recommendations []string `json:"recommendations,omitempty"`
}

// Handle выполняет запрос
func (h *getPlanStatisticsHandler) Handle(ctx context.Context, query any) (any, error) {
	q, ok := query.(*GetPlanStatisticsQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	// Получаем план
	plan, err := h.PlanRepo.FindByID(ctx, q.PlanID)
	if err != nil {
		return nil, fmt.Errorf("plan not found: %w", err)
	}

	// Получаем задания для плана
	tasks, err := h.TaskRepo.FindByPlan(ctx, q.PlanID)
	if err != nil {
		return nil, err
	}

	// Анализируем этапы
	stages := plan.Stages()
	totalStages := len(stages)
	completedStages := 0
	pendingStages := 0
	inProgressStages := 0
	skippedStages := 0

	for _, s := range stages {
		switch s.Status {
		case cropplan.StageStatusCompleted:
			completedStages++
		case cropplan.StageStatusPending:
			pendingStages++
		case cropplan.StageStatusInProgress:
			inProgressStages++
		case cropplan.StageStatusSkipped:
			skippedStages++
		}
	}

	progress := 0.0
	if totalStages > 0 {
		progress = float64(completedStages) / float64(totalStages) * 100
	}

	// Анализируем задания
	totalTasks := len(tasks)
	completedTasks := 0
	pendingTasks := 0
	overdueTasks := 0
	var totalCompletionDays float64
	var completionCount int

	for _, t := range tasks {
		switch t.Status {
		case task.TaskStatusCompleted:
			completedTasks++
			if t.CompletedAt != nil {
				days := t.CompletedAt.Sub(t.ScheduledDate).Hours() / 24
				if days > 0 {
					totalCompletionDays += days
					completionCount++
				}
			}
		case task.TaskStatusPending, task.TaskStatusInProgress:
			pendingTasks++
			if t.IsOverdue() {
				overdueTasks++
			}
		}
	}

	avgCompletionDays := 0.0
	if completionCount > 0 {
		avgCompletionDays = totalCompletionDays / float64(completionCount)
	}

	// Временные показатели
	daysSincePlanting := int(time.Since(plan.PlantingDate()).Hours() / 24)
	if daysSincePlanting < 0 {
		daysSincePlanting = 0
	}

	daysRemaining := 0
	if time.Now().Before(plan.SeasonEnd()) {
		daysRemaining = int(time.Until(plan.SeasonEnd()).Hours() / 24)
	}

	// Эффективность урожая
	yieldEfficiency := 0.0
	if plan.ExpectedYield() > 0 {
		yieldEfficiency = (plan.HarvestKg() / plan.ExpectedYield()) * 100
	}

	// Рекомендации
	recommendations := generateRecommendations(progress, overdueTasks, daysRemaining, yieldEfficiency)

	return &GetPlanStatisticsResponse{
		PlanID:                plan.ID(),
		PlanName:              plan.Name(),
		VarietyName:           plan.VarietyName(),
		Status:                string(plan.Status()),
		Progress:              progress,
		CompletedStages:       completedStages,
		TotalStages:           totalStages,
		PendingStages:         pendingStages,
		InProgressStages:      inProgressStages,
		SkippedStages:         skippedStages,
		DaysSincePlanting:     daysSincePlanting,
		DaysRemaining:         daysRemaining,
		DaysToHarvest:         daysRemaining, // упрощенно
		ExpectedYield:         plan.ExpectedYield(),
		ActualYield:           plan.HarvestKg(),
		YieldEfficiency:       yieldEfficiency,
		TotalTasks:            totalTasks,
		CompletedTasks:        completedTasks,
		PendingTasks:          pendingTasks,
		OverdueTasks:          overdueTasks,
		AvgTaskCompletionDays: avgCompletionDays,
		DeviationDays:         0,
		Recommendations:       recommendations,
	}, nil
}

// generateRecommendations генерирует рекомендации на основе статистики
func generateRecommendations(progress float64, overdueTasks, daysRemaining int, yieldEfficiency float64) []string {
	var recommendations []string

	if progress < 30 && daysRemaining < 14 {
		recommendations = append(recommendations, "⚠️ План сильно отстает от графика")
	}

	if overdueTasks > 0 {
		recommendations = append(recommendations, fmt.Sprintf("📋 Есть %d просроченных заданий", overdueTasks))
	}

	if daysRemaining < 7 && progress < 90 {
		recommendations = append(recommendations, "⏰ До конца сезона осталось меньше недели, ускорьте работы")
	}

	if yieldEfficiency > 0 && yieldEfficiency < 50 {
		recommendations = append(recommendations, "📉 Урожайность ниже ожидаемой, проанализируйте причины")
	}

	if len(recommendations) == 0 {
		recommendations = append(recommendations, "✅ План выполняется по графику")
	}

	return recommendations
}
