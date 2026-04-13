package task

import (
	"errors"
	"samurenkoroma/services/internal/core/domain/aggregate"
	"time"

	"github.com/google/uuid"
)

// Task задание для агронома
type Task struct {
	aggregate.BaseAggregate
	// Идентификация
	ID     string  `json:"id"`
	PlanID *string `json:"plan_id,omitempty"` // может быть nil (общее задание)
	BedID  string  `json:"bed_id"`

	// Кому назначено
	AssignedTo   string `json:"assigned_to"`   // ID агронома
	AssignedName string `json:"assigned_name"` // имя (денормализовано)

	// Тип и статус
	Type     TaskType     `json:"type"`
	Status   TaskStatus   `json:"status"`
	Priority TaskPriority `json:"priority"`

	// Содержание
	Title        string `json:"title"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`

	// Даты
	ScheduledDate time.Time  `json:"scheduled_date"`
	DueDate       time.Time  `json:"due_date"`
	CompletedAt   *time.Time `json:"completed_at,omitempty"`
	CompletedBy   string     `json:"completed_by,omitempty"`

	// Длительность
	EstimatedDuration int `json:"estimated_duration"` // минут
	ActualDuration    int `json:"actual_duration"`    // минут

	// Геолокация
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	// Медиа и комментарии
	Photos   []Photo   `json:"photos"`
	Comments []Comment `json:"comments"`

	// Системные поля
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewTask создает новое задание
func NewTask(
	id, bedID, assignedTo, assignedName string,
	taskType TaskType,
	title string,
	scheduledDate time.Time,
) (*Task, error) {

	if id == "" {
		return nil, errors.New("task id is required")
	}
	if bedID == "" {
		return nil, errors.New("bed id is required")
	}
	if assignedTo == "" {
		return nil, errors.New("assigned to is required")
	}
	if title == "" {
		return nil, errors.New("title is required")
	}

	now := time.Now()

	task := &Task{
		ID:            id,
		BedID:         bedID,
		AssignedTo:    assignedTo,
		AssignedName:  assignedName,
		Type:          taskType,
		Status:        TaskStatusPending,
		Priority:      PriorityMedium,
		Title:         title,
		ScheduledDate: scheduledDate,
		DueDate:       scheduledDate.AddDate(0, 0, 1),
		CreatedAt:     now,
		UpdatedAt:     now,
		Photos:        []Photo{},
		Comments:      []Comment{},
	}

	task.AddEvent(TaskCreatedEvent{
		TaskID:     id,
		BedID:      bedID,
		AssignedTo: assignedTo,
		Title:      title,
		//CreatedAt:  now,
	})

	return task, nil
}

// NewTaskWithPlan создает задание, привязанное к плану
func NewTaskWithPlan(
	id, planID, bedID, assignedTo, assignedName string,
	taskType TaskType,
	title string,
	scheduledDate time.Time,
) (*Task, error) {

	task, err := NewTask(id, bedID, assignedTo, assignedName, taskType, title, scheduledDate)
	if err != nil {
		return nil, err
	}
	task.PlanID = &planID
	return task, nil
}

// Start начинает выполнение задания
func (t *Task) Start() error {
	if t.Status != TaskStatusPending {
		return errors.New("only pending tasks can be started")
	}
	t.Status = TaskStatusInProgress
	t.UpdatedAt = time.Now()

	t.AddEvent(TaskStartedEvent{
		TaskID: t.ID,
		//StartedAt: time.Now(),
	})

	return nil
}

// Complete завершает задание
func (t *Task) Complete(duration int, completedBy string) error {
	if t.Status != TaskStatusInProgress && t.Status != TaskStatusPending {
		return errors.New("task cannot be completed")
	}

	now := time.Now()
	t.Status = TaskStatusCompleted
	t.CompletedAt = &now
	t.CompletedBy = completedBy
	t.ActualDuration = duration
	t.UpdatedAt = now

	t.AddEvent(TaskCompletedEvent{
		TaskID:      t.ID,
		CompletedBy: completedBy,
		//CompletedAt: now,
		Duration: duration,
	})

	return nil
}

// Skip пропускает задание
func (t *Task) Skip(reason string) error {
	if t.Status != TaskStatusPending {
		return errors.New("only pending tasks can be skipped")
	}

	t.Status = TaskStatusSkipped
	t.UpdatedAt = time.Now()

	t.AddComment("system", reason)

	t.AddEvent(TaskSkippedEvent{
		TaskID: t.ID,
		Reason: reason,
		//SkippedAt: time.Now(),
	})

	return nil
}

// AddPhoto добавляет фотографию
func (t *Task) AddPhoto(photo Photo) {
	t.Photos = append(t.Photos, photo)
	t.UpdatedAt = time.Now()
}

// AddComment добавляет комментарий
func (t *Task) AddComment(userID, text string) {
	t.Comments = append(t.Comments, Comment{
		ID:        uuid.New().String(),
		UserID:    userID,
		Text:      text,
		CreatedAt: time.Now(),
	})
	t.UpdatedAt = time.Now()
}

// UpdatePriority обновляет приоритет
func (t *Task) UpdatePriority(priority TaskPriority) {
	t.Priority = priority
	t.UpdatedAt = time.Now()
}

// IsOverdue проверяет, просрочено ли задание
func (t *Task) IsOverdue() bool {
	if t.Status.IsFinished() {
		return false
	}
	return time.Now().After(t.DueDate)
}
