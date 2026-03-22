package commands

import (
	"context"
	"errors"
	"fmt"
	"log"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/modules/growing/infrastructure/persistence/postgres"

	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/cropcycle"
)

// RecordOperationCmd — команда записи операции
type RecordOperationCmd struct {
	CycleID     string                  `json:"cycle_id" validate:"required"`
	Type        cropcycle.OperationType `json:"type" validate:"required"`
	Description string                  `json:"description"`
	Amount      float64                 `json:"amount"`
	Unit        string                  `json:"unit"`
	PerformedBy string                  `json:"performed_by"`
	Notes       string                  `json:"notes"`
}

// RecordOperationHandler — обработчик записи операции
type recordOperationHandler struct {
	uowFactory repository.Factory
}

func (h *recordOperationHandler) Name() string {
	return "RecordOperation"
}

func NewRecordOperationHandler(uowFactory repository.Factory) command.Handler {
	return &recordOperationHandler{
		uowFactory: uowFactory,
	}
}

func (h *recordOperationHandler) Handle(ctx context.Context, command any) error {
	cmd, ok := command.(RecordOperationCmd)
	if !ok {
		return errors.New("invalid command type")
	}
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	var cycle *cropcycle.CropCycle

	err = uow.Execute(ctx, postgres.NewGrowingProvider, func(provider repository.RepositoryProvider) error {
		growingProvider, ok := provider.(*postgres.GrowingProvider)
		if !ok {
			return fmt.Errorf("invalid provider type")
		}

		// Получаем цикл
		cycle, err = growingProvider.CropCycles().FindByID(ctx, cropcycle.CycleID(cmd.CycleID))
		if err != nil {
			return fmt.Errorf("failed to find cycle: %w", err)
		}
		if cycle == nil {
			return cropcycle.ErrCycleNotFound
		}

		// Создаём операцию
		op := cropcycle.Operation{
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
		if err := growingProvider.CropCycles().Save(ctx, cycle); err != nil {
			return fmt.Errorf("failed to save cycle: %w", err)
		}

		uow.RegisterAggregate(cycle)
		return nil
	})

	if err != nil {
		return err
	}

	log.Printf("Operation recorded: cycle=%s, type=%s", cmd.CycleID, cmd.Type)
	return nil
}
