package worktask

import (
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/worktask"
	"time"
)

type GenerateTasksHandler struct {
	Repo     cropplan.Repository
	TaskRepo worktask.Repository
}

func (h *GenerateTasksHandler) Handle(event cropplan.CropPlanActivated) error {
	plan, err := h.Repo.GetByID(event.PlanID)
	if err != nil {
		return err
	}

	var tasks []worktask.Task

	for _, step := range plan.Snapshot.Steps {
		var date time.Time

		switch {
		case step.DayOffset != 0:
			date = plan.StartDate.AddDate(0, 0, step.DayOffset)

		case step.BBCHStage != nil:
			date, err = h.Phenology.PredictDateForStage(
				plan.CropKey,
				plan.VarietyID,
				plan.LocationID,
				plan.StartDate,
				*step.BBCHStage,
			)
			if err != nil {
				return err
			}
		}

		tasks = append(tasks, worktask.Task{
			ID:            generateID(),
			CropPlanID:    plan.ID,
			Type:          step.Type,
			ScheduledDate: date,
			Status:        "pending",
			Params:        step.Params,
		})
	}

	return h.TaskRepo.SaveMany(tasks)
}
