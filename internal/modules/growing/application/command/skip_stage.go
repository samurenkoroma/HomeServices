package command

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"samurenkoroma/services/internal/common/application/uow"
	"samurenkoroma/services/internal/growing/cropplan/cropplan"
)

// SkipStageHandler команда пропуска этапа
type SkipStageHandler struct {
	UowFactory uow.Factory
}

// SkipStageCmd структура команды
type SkipStageCmd struct {
	PlanID  string `json:"plan_id"`
	StageID string `json:"stage_id"`
	Reason  string `json:"reason"`
}

// DecodeSkipStage декодирует JSON в команду
func DecodeSkipStage(data []byte) (any, error) {
	var cmd SkipStageCmd
	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, fmt.Errorf("failed to decode SkipStage command: %w", err)
	}

	if cmd.PlanID == "" {
		return nil, errors.New("plan_id is required")
	}
	if cmd.StageID == "" {
		return nil, errors.New("stage_id is required")
	}

	return cmd, nil
}

// Handle выполняет команду
func (h *SkipStageHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(SkipStageCmd)
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

	// Пропускаем этап
	if err := plan.SkipStage(c.StageID); err != nil {
		return fmt.Errorf("failed to skip stage: %w", err)
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
