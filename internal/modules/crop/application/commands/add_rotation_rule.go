package commands

import (
	"context"
	"fmt"
	"log"

	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/crop/domain/cropplan"
	"samurenkoroma/services/internal/modules/crop/domain/croptype"
	"samurenkoroma/services/internal/modules/crop/infrastructure/persistence/postgres"
)

// AddRotationRuleCommand — команда добавления правила севооборота
type AddRotationRuleCommand struct {
	PlanID                string `json:"plan_id" validate:"required"`
	PredecessorCropTypeID string `json:"predecessor_crop_type_id" validate:"required"`
	MinYears              int    `json:"min_years" validate:"required,min=1"`
	Recommended           bool   `json:"recommended"`
	Notes                 string `json:"notes"`
}

// AddRotationRuleHandler — обработчик добавления правила севооборота
type AddRotationRuleHandler struct {
	uowFactory   repository.Factory
	cropTypeRepo croptype.Repository // для валидации предшественника
}

// NewAddRotationRuleHandler создает новый обработчик
func NewAddRotationRuleHandler(
	uowFactory repository.Factory,
	cropTypeRepo croptype.Repository,
) *AddRotationRuleHandler {
	return &AddRotationRuleHandler{
		uowFactory:   uowFactory,
		cropTypeRepo: cropTypeRepo,
	}
}

// Handle обрабатывает команду
func (h *AddRotationRuleHandler) Handle(ctx context.Context, cmd AddRotationRuleCommand) error {
	// Валидируем существование предшественника
	if err := h.validatePredecessor(ctx, cmd.PredecessorCropTypeID); err != nil {
		return fmt.Errorf("invalid predecessor: %w", err)
	}

	uow, err := h.uowFactory.Begin(ctx, "crop")
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	var plan *cropplan.CropPlan

	err = uow.Execute(ctx, postgres.NewCropProvider, func(provider repository.RepositoryProvider) error {
		cropProvider, ok := provider.(*postgres.CropProvider)
		if !ok {
			return fmt.Errorf("invalid provider type: expected CropProvider, got %T", provider)
		}

		// Получаем план
		plan, err = cropProvider.CropPlans().GetByID(ctx, cropplan.PlanID(cmd.PlanID))
		if err != nil {
			return fmt.Errorf("failed to get crop plan: %w", err)
		}

		if plan == nil {
			return cropplan.ErrPlanNotFound
		}

		// Создаем правило севооборота
		rule, err := cropplan.NewRotationRule(
			cmd.PredecessorCropTypeID,
			cmd.MinYears,
			cmd.Recommended,
		)
		if err != nil {
			return fmt.Errorf("failed to create rotation rule: %w", err)
		}

		rule.Notes = cmd.Notes

		// Добавляем правило в план
		if err := plan.AddRotationRule(rule); err != nil {
			return fmt.Errorf("failed to add rotation rule: %w", err)
		}

		// Сохраняем обновленный план
		if err := cropProvider.CropPlans().Save(ctx, plan); err != nil {
			return fmt.Errorf("failed to save crop plan: %w", err)
		}

		uow.RegisterAggregate(plan)
		return nil
	})

	if err != nil {
		return err
	}

	// Логируем успешное добавление
	log.Printf("Rotation rule added: plan_id=%s, predecessor=%s, min_years=%d",
		cmd.PlanID, cmd.PredecessorCropTypeID, cmd.MinYears)

	return nil
}

// validatePredecessor проверяет существование предшественника
func (h *AddRotationRuleHandler) validatePredecessor(ctx context.Context, predecessorID string) error {
	// Для временного UOW используем отдельную транзакцию только для чтения
	// Это позволяет не блокировать основную транзакцию
	tempUOW, err := h.uowFactory.Begin(ctx, "crop")
	if err != nil {
		return err
	}

	var exists bool

	err = tempUOW.Execute(ctx, postgres.NewCropProvider, func(provider repository.RepositoryProvider) error {
		cropProvider, ok := provider.(*postgres.CropProvider)
		if !ok {
			return fmt.Errorf("invalid provider type")
		}

		// Проверяем существование типа культуры
		ct, err := cropProvider.CropTypes().FindByID(ctx, croptype.CropTypeID(predecessorID))
		if err != nil {
			if err == croptype.ErrCropTypeNotFound {
				return cropplan.ErrInvalidPredecessor
			}
			return err
		}

		exists = ct != nil && ct.IsActive()
		return nil
	})

	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("%w: crop type %s not found or inactive",
			cropplan.ErrInvalidPredecessor, predecessorID)
	}

	return nil
}
