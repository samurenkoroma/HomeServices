package command

import (
	"samurenkoroma/services/internal/field/application"
	"samurenkoroma/services/internal/field/domain/cropplan"
)

type CreateCropPlan struct {
	Uow application.UnitOfWork
}

type CreateCropPlanCmd struct {
	ID    string
	BedID string
	Crop  string
}

func (h *CreateCropPlan) Handle(cmd CreateCropPlanCmd) error {
	plan := cropplan.New(
		cropplan.CropPlanID(cmd.ID),
		cropplan.BedID(cmd.BedID),
		cmd.Crop,
	)

	err := h.Uow.CropPlans().Save(plan)
	if err != nil {
		return err
	}

	return h.Uow.Commit()
}
