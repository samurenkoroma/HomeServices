package command

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"samurenkoroma/services/internal/common/application/uow"
	"samurenkoroma/services/internal/growing/cropplan/cropplan"
)

// StartStageHandler команда начала выполнения этапа
type StartStageHandler struct {
	UowFactory uow.Factory
}

// StartStageCmd структура команды
type StartStageCmd struct {
	PlanID      string `json:"plan_id"`
	StageID     string `json:"stage_id"`
	CurrentBBCH int    `json:"current_bbch"` // текущая BBCH фаза (из phenology)
}

// DecodeStartStage декодирует JSON в команду
func DecodeStartStage(data []byte) (any, error) {
	var cmd StartStageCmd
	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, fmt.Errorf("failed to decode StartStage command: %w", err)
	}

	if cmd.PlanID == "" {
		return nil, errors.New("plan_id is required")
	}
	if cmd.StageID == "" {
		return nil, errors.New("stage_id is required")
	}
	if cmd.CurrentBBCH < 0 {
		return nil, errors.New("current_bbch must be non-negative")
	}

	return cmd, nil
}

// Handle выполняет команду
func (h *StartStageHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(StartStageCmd)
	if !ok {
		return errors.New("invalid command type")
	}

	// Начинаем транзакцию
	uowObj, err := h.UowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin unit of work: %w", err)
	}
	defer uowObj.Rollback()

	// Получаем план
	planRepo := uowObj.CropPlans()
	plan, err := planRepo.FindByID(ctx, c.PlanID)
	if err != nil {
		return fmt.Errorf("failed to find plan: %w", err)
	}

	// Проверяем, что план активен
	if plan.Status() != cropplan.StatusActive {
		return errors.New("can only start stages in active plan")
	}

	// Начинаем этап
	if err := plan.StartStage(c.StageID, c.CurrentBBCH); err != nil {
		return fmt.Errorf("failed to start stage: %w", err)
	}

	// Сохраняем изменения
	if err := planRepo.Update(ctx, plan); err != nil {
		return fmt.Errorf("failed to update plan: %w", err)
	}

	uowObj.RegisterAggregate(plan)

	if err := uowObj.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
