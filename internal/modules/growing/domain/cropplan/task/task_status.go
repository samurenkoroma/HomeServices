package task

// TaskStatus статус выполнения задания
type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"     // ожидает выполнения
	TaskStatusInProgress TaskStatus = "in_progress" // в процессе
	TaskStatusCompleted  TaskStatus = "completed"   // выполнено
	TaskStatusSkipped    TaskStatus = "skipped"     // пропущено
	TaskStatusCancelled  TaskStatus = "cancelled"   // отменено
)

// String возвращает русское название статуса
func (s TaskStatus) String() string {
	switch s {
	case TaskStatusPending:
		return "Ожидает"
	case TaskStatusInProgress:
		return "В процессе"
	case TaskStatusCompleted:
		return "Выполнено"
	case TaskStatusSkipped:
		return "Пропущено"
	case TaskStatusCancelled:
		return "Отменено"
	default:
		return "Неизвестно"
	}
}

// IsFinished проверяет, является ли статус конечным
func (s TaskStatus) IsFinished() bool {
	return s == TaskStatusCompleted || s == TaskStatusSkipped || s == TaskStatusCancelled
}
