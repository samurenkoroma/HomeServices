package domain

type CropPlanStatus string

const (
	StatusDraft      CropPlanStatus = "draft"
	StatusPublished  CropPlanStatus = "published"
	StatusDeprecated CropPlanStatus = "deprecated"

	StatusActive    CropPlanStatus = "active"
	StatusCompleted CropPlanStatus = "completed"
	StatusCancelled CropPlanStatus = "cancelled"
)
