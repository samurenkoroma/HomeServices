package utils

import (
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cultivation"
)

func MapTemplateToSnapshot(t cultivation.Step) cropplan.CultivationStep {
	return cropplan.CultivationStep{
		ID:   t.ID,
		Type: t.Type,
		Trigger: cropplan.StepTrigger{
			Type:  string(t.Trigger.Type),
			Value: t.Trigger.Value,
		},
	}
}

func BuildSnapshot(plan *cultivation.CultivationPlan) cropplan.CultivationSnapshot {
	steps := make([]cropplan.CultivationStep, 0, len(plan.Steps))

	for _, s := range plan.Steps {
		steps = append(steps, MapTemplateToSnapshot(s))
	}

	return cropplan.CultivationSnapshot{
		ID:      plan.ID,
		Version: plan.Version,
		Steps:   steps,
	}
}
