package task

import (
	"context"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/catalog"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/phenology"
	"time"

	"github.com/google/uuid"
)

// TaskGenerator генерирует задания на основе фенологии
type TaskGenerator struct {
	phenologyService *phenology.PhenologyService
	catalogRepo      catalog.Repository
}

// NewTaskGenerator создает новый генератор заданий
func NewTaskGenerator(
	phenologyService *phenology.PhenologyService,
	catalogRepo catalog.Repository,
) *TaskGenerator {
	return &TaskGenerator{
		phenologyService: phenologyService,
		catalogRepo:      catalogRepo,
	}
}

// GenerateContext контекст для генерации заданий
type GenerateContext struct {
	Date         time.Time // дата, на которую генерируем
	BedID        string    // ID грядки
	PlanID       string    // ID плана (если есть)
	VarietyID    string    // ID сорта
	PlantingDate time.Time // дата посадки
	Latitude     float64   // широта
	Longitude    float64   // долгота
	AssignedTo   string    // ID агронома
	AssignedName string    // имя агронома
}

// GenerateForPlan генерирует задания для конкретного плана на указанную дату
func (g *TaskGenerator) GenerateForPlan(
	ctx context.Context,
	ctxGen GenerateContext,
) ([]Task, error) {

	var tasks []Task

	// 1. Получаем текущую фенологию
	phenology, err := g.phenologyService.GetCurrentPhenology(
		ctx,
		ctxGen.PlanID,
		ctxGen.VarietyID,
		ctxGen.PlantingDate,
		ctxGen.Latitude,
		ctxGen.Longitude,
	)
	if err != nil {
		return nil, err
	}

	// 2. Генерируем задания из рекомендаций
	for _, action := range phenology.RecommendedActions {
		// Проверяем, не создавали ли уже такое задание
		// (в реальности здесь будет проверка в БД)

		task, err := g.createTaskFromAction(
			ctxGen,
			action,
			phenology.CurrentPhaseCode,
		)
		if err != nil {
			continue
		}
		tasks = append(tasks, *task)
	}

	// 3. Добавляем регулярные задания в зависимости от фазы
	regularTasks := g.generateRegularTasks(ctxGen, phenology.CurrentPhaseCode)
	tasks = append(tasks, regularTasks...)

	return tasks, nil
}

// createTaskFromAction создает задание из рекомендации
func (g *TaskGenerator) createTaskFromAction(
	ctxGen GenerateContext,
	action phenology.RecommendedAction,
	phaseCode string,
) (*Task, error) {

	// Определяем тип задания по названию
	taskType := g.mapActionToTaskType(action.Title)

	scheduledDate := ctxGen.Date
	if action.DueDays > 0 {
		scheduledDate = ctxGen.Date.AddDate(0, 0, action.DueDays)
	}

	title := action.Title
	if phaseCode != "" {
		title = title + " (" + phaseCode + ")"
	}

	task, err := NewTaskWithPlan(
		uuid.New().String(),
		ctxGen.PlanID,
		ctxGen.BedID,
		ctxGen.AssignedTo,
		ctxGen.AssignedName,
		taskType,
		title,
		scheduledDate,
	)
	if err != nil {
		return nil, err
	}

	task.Description = action.Description
	task.Instructions = g.getDetailedInstructions(action.Title, phaseCode)
	task.Priority = g.mapPriority(action.Priority)
	task.Latitude = ctxGen.Latitude
	task.Longitude = ctxGen.Longitude

	return task, nil
}

// generateRegularTasks генерирует регулярные задания
func (g *TaskGenerator) generateRegularTasks(
	ctxGen GenerateContext,
	phaseCode string,
) []Task {

	var tasks []Task

	// Осмотр грядки (каждые 3 дня)
	if g.shouldGenerateInspection(ctxGen.Date, phaseCode) {
		inspection, _ := NewTaskWithPlan(
			uuid.New().String(),
			ctxGen.PlanID,
			ctxGen.BedID,
			ctxGen.AssignedTo,
			ctxGen.AssignedName,
			TaskTypeInspection,
			"Осмотр грядки ("+phaseCode+")",
			ctxGen.Date,
		)
		inspection.Description = "Проверить состояние растений, наличие вредителей и болезней"
		inspection.Priority = PriorityMedium
		tasks = append(tasks, *inspection)
	}

	// Полив (в зависимости от фазы и погоды)
	if g.shouldGenerateWatering(phaseCode) {
		watering, _ := NewTaskWithPlan(
			uuid.New().String(),
			ctxGen.PlanID,
			ctxGen.BedID,
			ctxGen.AssignedTo,
			ctxGen.AssignedName,
			TaskTypeWatering,
			"Полив",
			ctxGen.Date,
		)
		watering.Description = g.getWateringDescription(phaseCode)
		watering.Priority = g.getWateringPriority(phaseCode)
		tasks = append(tasks, *watering)
	}

	return tasks
}

// shouldGenerateInspection проверяет, нужно ли создать задание на осмотр
func (g *TaskGenerator) shouldGenerateInspection(date time.Time, phaseCode string) bool {
	// В критических фазах осмотр каждый день
	criticalPhases := []string{"BBCH-61", "BBCH-71"}
	for _, cp := range criticalPhases {
		if cp == phaseCode {
			return true
		}
	}
	// Иначе раз в 3 дня (по дню месяца)
	return date.Day()%3 == 1
}

// shouldGenerateWatering проверяет, нужен ли полив
func (g *TaskGenerator) shouldGenerateWatering(phaseCode string) bool {
	// В фазе цветения и плодоношения полив нужен чаще
	wateringPhases := []string{"BBCH-61", "BBCH-71", "BBCH-81"}
	for _, wp := range wateringPhases {
		if wp == phaseCode {
			return true
		}
	}
	return false
}

// getWateringDescription возвращает описание полива для фазы
func (g *TaskGenerator) getWateringDescription(phaseCode string) string {
	switch phaseCode {
	case "BBCH-61":
		return "Обильный полив во время цветения для хорошего опыления"
	case "BBCH-71":
		return "Регулярный полив для формирования плодов"
	case "BBCH-81":
		return "Умеренный полив, избегать переувлажнения"
	default:
		return "Полив по необходимости"
	}
}

// getWateringPriority возвращает приоритет полива
func (g *TaskGenerator) getWateringPriority(phaseCode string) TaskPriority {
	switch phaseCode {
	case "BBCH-61", "BBCH-71":
		return PriorityHigh
	default:
		return PriorityMedium
	}
}

// mapActionToTaskType маппит название действия на тип задания
func (g *TaskGenerator) mapActionToTaskType(actionTitle string) TaskType {
	actionMap := map[string]TaskType{
		"Осмотр":       TaskTypeInspection,
		"Полив":        TaskTypeWatering,
		"Подкормка":    TaskTypeFertilizing,
		"Обработка":    TaskTypeTreatment,
		"Обрезка":      TaskTypePruning,
		"Сбор урожая":  TaskTypeHarvest,
		"Удаление":     TaskTypePruning,
		"Нормирование": TaskTypePruning,
		"Стимулятор":   TaskTypeFertilizing,
	}

	for key, tt := range actionMap {
		if contains(actionTitle, key) {
			return tt
		}
	}
	return TaskTypeInspection
}

// mapPriority маппит строковый приоритет в тип
func (g *TaskGenerator) mapPriority(priority string) TaskPriority {
	switch priority {
	case "low":
		return PriorityLow
	case "medium":
		return PriorityMedium
	case "high":
		return PriorityHigh
	case "urgent":
		return PriorityUrgent
	default:
		return PriorityMedium
	}
}

// getDetailedInstructions возвращает подробные инструкции для задания
func (g *TaskGenerator) getDetailedInstructions(actionTitle, phaseCode string) string {
	instructions := map[string]string{
		"Обработка от вредителей": "1. Осмотреть растения на наличие вредителей\n2. Приготовить раствор по инструкции\n3. Опрыскать растения, особенно нижнюю сторону листьев\n4. Сделать фото до и после обработки",
		"Подкормка":               "1. Приготовить раствор удобрения\n2. Внести под корень\n3. Избегать попадания на листья\n4. После подкормки полить чистой водой",
		"Осмотр всходов":          "1. Проверить густоту всходов\n2. При загущении - проредить\n3. Отметить проблемные места на схеме",
		"Нормирование завязей":    "1. Удалить мелкие и деформированные завязи\n2. Оставить 3-4 самых крупных на кисть\n3. Сделать фото результата",
	}

	if inst, ok := instructions[actionTitle]; ok {
		return inst
	}
	return "Выполнить согласно регламенту работ"
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0)
}
