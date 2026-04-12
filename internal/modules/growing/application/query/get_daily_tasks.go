package query

import (
	"context"
	"encoding/json"
	"errors"
	"samurenkoroma/services/internal/modules/growing/application/dto"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/task"
	"time"
)

// GetDailyTasksHandler запрос ежедневных заданий
type GetDailyTasksHandler struct {
	TaskRepo task.Repository
}

// GetDailyTasksQuery параметры запроса
type GetDailyTasksQuery struct {
	UserID           string     `json:"user_id"`
	Date             *time.Time `json:"date,omitempty"`
	IncludeCompleted bool       `json:"include_completed,omitempty"`
}

// GetDailyTasksResponse ответ с заданиями
type GetDailyTasksResponse struct {
	Date    string        `json:"date"`
	Tasks   []dto.TaskDTO `json:"tasks"`
	Summary TaskSummary   `json:"summary"`
}

// TaskSummary сводка по заданиям
type TaskSummary struct {
	Total          int `json:"total"`
	Completed      int `json:"completed"`
	Pending        int `json:"pending"`
	Overdue        int `json:"overdue"`
	HighPriority   int `json:"high_priority"`
	UrgentPriority int `json:"urgent_priority"`
}

// DecodeGetDailyTasks декодирует JSON в запрос
func DecodeGetDailyTasks(data []byte) (any, error) {
	var q GetDailyTasksQuery
	if err := json.Unmarshal(data, &q); err != nil {
		return nil, err
	}
	if q.UserID == "" {
		return nil, errors.New("user_id is required")
	}
	return q, nil
}

// Handle выполняет запрос
func (h *GetDailyTasksHandler) Handle(ctx context.Context, query any) (any, error) {
	q, ok := query.(GetDailyTasksQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	// Определяем дату
	targetDate := time.Now()
	if q.Date != nil {
		targetDate = *q.Date
	}
	startOfDay := time.Date(targetDate.Year(), targetDate.Month(), targetDate.Day(), 0, 0, 0, 0, targetDate.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1)

	// Получаем задания
	tasks, err := h.TaskRepo.FindByAssignedTo(ctx, q.UserID, startOfDay, endOfDay)
	if err != nil {
		return nil, err
	}

	// Фильтрация по статусу
	var filteredTasks []task.Task
	for _, t := range tasks {
		if !q.IncludeCompleted && t.Status == task.TaskStatusCompleted {
			continue
		}
		filteredTasks = append(filteredTasks, t)
	}

	// Считаем статистику
	summary := TaskSummary{
		Total: len(filteredTasks),
	}
	for _, t := range filteredTasks {
		switch t.Status {
		case task.TaskStatusCompleted:
			summary.Completed++
		case task.TaskStatusPending, task.TaskStatusInProgress:
			summary.Pending++
		}

		if t.IsOverdue() && t.Status != task.TaskStatusCompleted {
			summary.Overdue++
		}
		if t.Priority == task.PriorityHigh {
			summary.HighPriority++
		}
		if t.Priority == task.PriorityUrgent {
			summary.UrgentPriority++
		}
	}

	// Конвертируем в DTO
	tasksDTO := make([]dto.TaskDTO, len(filteredTasks))
	for i, t := range filteredTasks {
		tasksDTO[i] = dto.ToTaskDTO(&t)
	}

	return &GetDailyTasksResponse{
		Date:    startOfDay.Format("2006-01-02"),
		Tasks:   tasksDTO,
		Summary: summary,
	}, nil
}
