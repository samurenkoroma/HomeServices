package eventhandlers

import (
	"context"
	"samurenkoroma/services/internal/common/application/uow"
	"samurenkoroma/services/internal/growing/domain/croptemplate"
	"samurenkoroma/services/internal/shared/events"
)

func OnCropPlanPublished(ctx context.Context, e events.CropPlanPublished) error {

	UOW, ok := uow.FromContext(ctx)
	if !ok {
		return nil
	}

	template := croptemplate.New(
		e.PlanID,
		e.Version,
	)

	return UOW.CropTemplates().Save(ctx, template)
}
func OnCropPlanDeprecated(ctx context.Context, e events.CropPlanDeprecated) error {

	UOW, ok := uow.FromContext(ctx)
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
