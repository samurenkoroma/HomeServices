package eventhandlers

import (
	"context"
	"samurenkoroma/services/internal/core/port/repository"
	"samurenkoroma/services/internal/modules/crop/domain"
	"samurenkoroma/services/internal/modules/growing/domain/croptemplate"
)

func OnCropPlanPublished(ctx context.Context, e domain.CropPlanPublished) error {

	UOW, ok := repository.FromContext(ctx)
	if !ok {
		return nil
	}

	template := croptemplate.New(
		e.PlanID,
		e.Version,
	)

	return UOW.CropTemplates().Save(ctx, template)
}
func OnCropPlanDeprecated(ctx context.Context, e domain.CropPlanDeprecated) error {

	UOW, ok := repository.FromContext(ctx)
	if !ok {
		return nil
	}
	template, err := UOW.CropTemplates().ByPlanID(ctx, e.PlanID)
	if err != nil || template == nil {
		return err
	}

	template.Disable()

	return UOW.CropTemplates().Save(ctx, template)
}
