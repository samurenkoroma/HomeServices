package cropplan

import (
	"context"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/task"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/inmemory"
	"testing"
	"time"
)

func TestCropPlanRepository(t *testing.T) {
	repo := inmemory.NewCropPlanRepo()
	ctx := context.Background()

	// Создаем план
	plan, _ := NewCropPlan(
		"plan-1", "bed-1", "Тестовый план",
		"bull_heart", "Бычье сердце", "Томат",
		time.Now(), time.Now().AddDate(0, 3, 0),
		time.Now().AddDate(0, 0, 1),
		55.7558, 37.6173,
		"user-1", "Иван Иванов",
	)

	// Сохраняем
	err := repo.Save(ctx, plan)
	if err != nil {
		t.Errorf("Save failed: %v", err)
	}

	// Находим
	found, err := repo.FindByID(ctx, "plan-1")
	if err != nil {
		t.Errorf("FindByID failed: %v", err)
	}
	if found.ID() != "plan-1" {
		t.Errorf("Expected plan-1, got %s", found.ID())
	}
}

func TestTaskRepository(t *testing.T) {
	repo := inmemory.NewTaskRepo()
	ctx := context.Background()

	// Создаем задание
	task, _ := task.NewTask(
		"task-1", "bed-1", "user-1", "Иван Иванов",
		task.TaskTypeInspection,
		"Осмотр грядки",
		time.Now(),
	)

	// Сохраняем
	err := repo.Save(ctx, task)
	if err != nil {
		t.Errorf("Save failed: %v", err)
	}

	// Находим
	found, err := repo.FindByID(ctx, "task-1")
	if err != nil {
		t.Errorf("FindByID failed: %v", err)
	}
	if found.ID != "task-1" {
		t.Errorf("Expected task-1, got %s", found.ID)
	}
}
