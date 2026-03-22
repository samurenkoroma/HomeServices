package commands

import (
	"context"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/cropcycle"
)

type RecordOperationCommand struct {
	CycleID     string                  `json:"cycle_id" validate:"required"`
	Type        cropcycle.OperationType `json:"type" validate:"required"`
	Description string                  `json:"description"`
	Amount      float64                 `json:"amount"`
	Unit        string                  `json:"unit"`
	PerformedBy string                  `json:"performed_by"`
	Notes       string                  `json:"notes"`
}

type RecordOperationHandler struct {
	uowFactory repository.Factory
}

func (h *RecordOperationHandler) Handle(ctx context.Context, cmd RecordOperationCommand) error {
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return err
	}

	return uow.Execute(ctx, func(provider repository.RepositoryProvider) error {
		growingProvider := provider.(*postgres.GrowingProvider)

		// Получаем цикл
		cycle, err := growingProvider.CropCycles().GetByID(ctx, cmd.CycleID)
		if err != nil {
			return err
		}

		// Создаем операцию
		op := cropcycle.Operation{
			ID:          generateID(),
			Type:        cmd.Type,
			Description: cmd.Description,
			Amount:      cmd.Amount,
			Unit:        cmd.Unit,
			PerformedBy: cmd.PerformedBy,
			Notes:       cmd.Notes,
		}

		// Записываем
		if err := cycle.RecordOperation(op); err != nil {
			return err
		}

		// Сохраняем
		return growingProvider.CropCycles().Save(ctx, cycle)
	})
}
