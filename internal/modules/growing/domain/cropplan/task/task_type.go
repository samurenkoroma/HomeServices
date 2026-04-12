package task

// TaskType тип агрономического задания
type TaskType string

const (
	TaskTypeInspection  TaskType = "inspection"  // осмотр
	TaskTypeWatering    TaskType = "watering"    // полив
	TaskTypeFertilizing TaskType = "fertilizing" // подкормка
	TaskTypeTreatment   TaskType = "treatment"   // обработка от вредителей
	TaskTypePruning     TaskType = "pruning"     // обрезка/формирование
	TaskTypeHarvest     TaskType = "harvest"     // сбор урожая
	TaskTypeSoilWork    TaskType = "soil_work"   // работы с почвой
	TaskTypeEquipment   TaskType = "equipment"   // работа с оборудованием
	TaskTypeReport      TaskType = "report"      // отчет
)

// String возвращает русское название типа задания
func (t TaskType) String() string {
	switch t {
	case TaskTypeInspection:
		return "Осмотр"
	case TaskTypeWatering:
		return "Полив"
	case TaskTypeFertilizing:
		return "Подкормка"
	case TaskTypeTreatment:
		return "Обработка"
	case TaskTypePruning:
		return "Обрезка"
	case TaskTypeHarvest:
		return "Сбор урожая"
	case TaskTypeSoilWork:
		return "Работа с почвой"
	case TaskTypeEquipment:
		return "Работа с оборудованием"
	case TaskTypeReport:
		return "Отчет"
	default:
		return "Неизвестно"
	}
}

// Icon возвращает иконку для UI
func (t TaskType) Icon() string {
	switch t {
	case TaskTypeInspection:
		return "🔍"
	case TaskTypeWatering:
		return "💧"
	case TaskTypeFertilizing:
		return "🧪"
	case TaskTypeTreatment:
		return "🪲"
	case TaskTypePruning:
		return "✂️"
	case TaskTypeHarvest:
		return "🌾"
	default:
		return "📋"
	}
}
