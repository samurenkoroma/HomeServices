package domain

import (
	"context"
	"samurenkoroma/services/internal/modules/growing/domain/cropcycle"
	"samurenkoroma/services/internal/modules/growing/domain/croptemplate"
)

type CropCycleRepository interface {
	Save(ctx context.Context, cycle *cropcycle.CropCycle) error
	ByID(ctx context.Context, id string) (*cropcycle.CropCycle, error)
}

type CropTemplateRepository interface {
	Save(ctx context.Context, template *croptemplate.CropTemplate) error
	ByPlanID(ctx context.Context, planID string) (*croptemplate.CropTemplate, error)
}
