package inmemory

import (
	"context"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/task"
	"sync"
	"time"
)

// TaskRepo in-memory реализация репозитория заданий
type TaskRepo struct {
	mu    sync.RWMutex
	tasks map[string]*task.Task
}

// NewTaskRepo создает новый in-memory репозиторий заданий
func NewTaskRepo() *TaskRepo {
	return &TaskRepo{
		tasks: make(map[string]*task.Task),
	}
}

// Save сохраняет новое задание
func (r *TaskRepo) Save(ctx context.Context, t *task.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[t.ID]; exists {
		return ErrDuplicateTask
	}

	r.tasks[t.ID] = t
	return nil
}

// Update обновляет существующее задание
func (r *TaskRepo) Update(ctx context.Context, t *task.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[t.ID]; !exists {
		return task.ErrTaskNotFound
	}

	r.tasks[t.ID] = t
	return nil
}

// FindByID находит задание по ID
func (r *TaskRepo) FindByID(ctx context.Context, id string) (*task.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	t, exists := r.tasks[id]
	if !exists {
		return nil, task.ErrTaskNotFound
	}

	return t, nil
}

// Delete удаляет задание
func (r *TaskRepo) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return task.ErrTaskNotFound
	}

	delete(r.tasks, id)
	return nil
}

// FindByAssignedTo находит задания, назначенные агроному
func (r *TaskRepo) FindByAssignedTo(ctx context.Context, userID string, from, to time.Time) ([]task.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []task.Task
	for _, t := range r.tasks {
		if t.AssignedTo != userID {
			continue
		}
		// Фильтр по дате
		if !from.IsZero() && t.ScheduledDate.Before(from) {
			continue
		}
		if !to.IsZero() && t.ScheduledDate.After(to) {
			continue
		}
		result = append(result, *t)
	}
	return result, nil
}

// FindByPlan находит все задания для плана
func (r *TaskRepo) FindByPlan(ctx context.Context, planID string) ([]task.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []task.Task
	for _, t := range r.tasks {
		if t.PlanID != nil && *t.PlanID == planID {
			result = append(result, *t)
		}
	}
	return result, nil
}

// FindByBed находит задания для грядки за период
func (r *TaskRepo) FindByBed(ctx context.Context, bedID string, from, to time.Time) ([]task.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []task.Task
	for _, t := range r.tasks {
		if t.BedID != bedID {
			continue
		}
		// Фильтр по дате
		if !from.IsZero() && t.ScheduledDate.Before(from) {
			continue
		}
		if !to.IsZero() && t.ScheduledDate.After(to) {
			continue
		}
		result = append(result, *t)
	}
	return result, nil
}

// FindPending находит ожидающие задания для агронома
func (r *TaskRepo) FindPending(ctx context.Context, userID string) ([]task.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []task.Task
	for _, t := range r.tasks {
		if t.AssignedTo == userID && t.Status == task.TaskStatusPending {
			result = append(result, *t)
		}
	}
	return result, nil
}

// FindOverdue находит просроченные задания для агронома
func (r *TaskRepo) FindOverdue(ctx context.Context, userID string) ([]task.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	now := time.Now()
	var result []task.Task
	for _, t := range r.tasks {
		if t.AssignedTo == userID &&
			t.Status == task.TaskStatusPending &&
			t.DueDate.Before(now) {
			result = append(result, *t)
		}
	}
	return result, nil
}

// CompleteTasks массово завершает задания
func (r *TaskRepo) CompleteTasks(ctx context.Context, taskIDs []string, completedBy string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, id := range taskIDs {
		t, exists := r.tasks[id]
		if !exists {
			continue
		}
		if t.Status == task.TaskStatusPending || t.Status == task.TaskStatusInProgress {
			now := time.Now()
			t.Status = task.TaskStatusCompleted
			t.CompletedAt = &now
			t.CompletedBy = completedBy
			t.UpdatedAt = now
		}
	}
	return nil
}

// ReassignTasks переназначает задания другому агроному
func (r *TaskRepo) ReassignTasks(ctx context.Context, taskIDs []string, newUserID, newUserName string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, id := range taskIDs {
		t, exists := r.tasks[id]
		if !exists {
			continue
		}
		if !t.Status.IsFinished() {
			t.AssignedTo = newUserID
			t.AssignedName = newUserName
			t.UpdatedAt = time.Now()
		}
	}
	return nil
}

// GetTodayTasks возвращает задания на сегодня для агронома
func (r *TaskRepo) GetTodayTasks(ctx context.Context, userID string) ([]task.Task, error) {
	today := time.Now()
	startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1)

	return r.FindByAssignedTo(ctx, userID, startOfDay, endOfDay)
}

// Clear очищает репозиторий (для тестов)
func (r *TaskRepo) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks = make(map[string]*task.Task)
}

// Count возвращает количество заданий (для тестов)
func (r *TaskRepo) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.tasks)
}
