package task

// TaskPriority приоритет задания
type TaskPriority string

const (
	PriorityLow    TaskPriority = "low"    // низкий
	PriorityMedium TaskPriority = "medium" // средний
	PriorityHigh   TaskPriority = "high"   // высокий
	PriorityUrgent TaskPriority = "urgent" // срочный
)

// String возвращает русское название приоритета
func (p TaskPriority) String() string {
	switch p {
	case PriorityLow:
		return "Низкий"
	case PriorityMedium:
		return "Средний"
	case PriorityHigh:
		return "Высокий"
	case PriorityUrgent:
		return "Срочный"
	default:
		return "Средний"
	}
}

// Color возвращает цвет для UI
func (p TaskPriority) Color() string {
	switch p {
	case PriorityLow:
		return "#4CAF50" // зеленый
	case PriorityMedium:
		return "#FFC107" // желтый
	case PriorityHigh:
		return "#FF9800" // оранжевый
	case PriorityUrgent:
		return "#F44336" // красный
	default:
		return "#9E9E9E"
	}
}
