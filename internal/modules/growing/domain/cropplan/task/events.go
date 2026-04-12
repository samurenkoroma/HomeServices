package task

import (
	"samurenkoroma/services/internal/core/domain/event"
)

// TaskCreatedEvent событие создания задания
type TaskCreatedEvent struct {
	event.BaseEvent
	TaskID     string
	BedID      string
	AssignedTo string
	Title      string
}

func (e TaskCreatedEvent) EventName() string { return "task.created" }

// TaskStartedEvent событие начала выполнения
type TaskStartedEvent struct {
	event.BaseEvent
	TaskID string
}

func (e TaskStartedEvent) EventName() string { return "task.started" }

// TaskCompletedEvent событие завершения
type TaskCompletedEvent struct {
	event.BaseEvent
	TaskID      string
	CompletedBy string
	Duration    int
}

func (e TaskCompletedEvent) EventName() string { return "task.completed" }

// TaskSkippedEvent событие пропуска
type TaskSkippedEvent struct {
	event.BaseEvent
	TaskID string
	Reason string
}

func (e TaskSkippedEvent) EventName() string { return "task.skipped" }
