package command

import (
	"context"
	"samurenkoroma/services/internal/growing/domain"
	"samurenkoroma/services/internal/growing/domain/cropcycle"
	"samurenkoroma/services/internal/growing/domain/croptemplate"
)

type StartCropCycle struct {
	CycleID    string
	PlanID     string
	FacilityID string
	BedID      string
}
type StartCropCycleHandler struct {
	cycleRepo    domain.CropCycleRepository
	templateRepo domain.CropTemplateRepository
}

func NewStartCropCycleHandler(
	cycleRepo domain.CropCycleRepository,
	templateRepo domain.CropTemplateRepository,
) *StartCropCycleHandler {
	return &StartCropCycleHandler{
		cycleRepo:    cycleRepo,
		templateRepo: templateRepo,
	}
}

func (h *StartCropCycleHandler) Handle(ctx context.Context, cmd StartCropCycle) error {

	template, err := h.templateRepo.ByPlanID(ctx, cmd.PlanID)
	if err != nil {
		return err
	}

	if template == nil || !template.Active() {
		return croptemplate.ErrTemplateNotActive
	}

	cycle := cropcycle.New(
		cmd.CycleID,
		cmd.PlanID,
		template.Version(),
		cmd.FacilityID,
		cmd.BedID,
	)

	if err := cycle.Start(); err != nil {
		return err
	}

	return h.cycleRepo.Save(ctx, cycle)
}
