package task

import (
	"context"
	"time"
)

// Repository интерфейс для работы с заданиями
type Repository interface {
	// Базовые операции
	Save(ctx context.Context, task *Task) error
	Update(ctx context.Context, task *Task) error
	FindByID(ctx context.Context, id string) (*Task, error)
	Delete(ctx context.Context, id string) error

	// Запросы
	FindByAssignedTo(ctx context.Context, userID string, from, to time.Time) ([]Task, error)
	FindByPlan(ctx context.Context, planID string) ([]Task, error)
	FindByBed(ctx context.Context, bedID string, from, to time.Time) ([]Task, error)
	FindPending(ctx context.Context, userID string) ([]Task, error)
	FindOverdue(ctx context.Context, userID string) ([]Task, error)

	// Массовые операции
	CompleteTasks(ctx context.Context, taskIDs []string, completedBy string) error
	ReassignTasks(ctx context.Context, taskIDs []string, newUserID, newUserName string) error
}
