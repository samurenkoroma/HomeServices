package cropplan

import "errors"

var (
	// Общие ошибки
	ErrPlanNotFound  = errors.New("crop plan not found")
	ErrInvalidPlanID = errors.New("invalid plan id")
	ErrInvalidBedID  = errors.New("invalid bed id")

	// Ошибки статуса
	ErrPlanNotDraft         = errors.New("plan is not in draft status")
	ErrPlanNotActive        = errors.New("plan is not active")
	ErrPlanAlreadyActive    = errors.New("plan already active")
	ErrPlanAlreadyCompleted = errors.New("plan already completed")
	ErrPlanAlreadyCancelled = errors.New("plan already cancelled")

	// Ошибки этапов
	ErrNoStages            = errors.New("plan has no stages")
	ErrStageNotFound       = errors.New("stage not found")
	ErrStageAlreadyStarted = errors.New("stage already started")
	ErrStageNotInProgress  = errors.New("stage is not in progress")
	ErrStageOrderDuplicate = errors.New("stage with this order already exists")
	ErrInvalidStageOrder   = errors.New("invalid stage order")

	// Ошибки сезона и посадки
	ErrInvalidSeason        = errors.New("invalid season dates")
	ErrInvalidPlantingDate  = errors.New("planting date must be within season")
	ErrPlantingDateRequired = errors.New("planting date is required")

	// Ошибки сорта
	ErrVarietyRequired = errors.New("variety is required")
	ErrVarietyNotFound = errors.New("variety not found in catalog")

	// Ошибки ротации
	ErrRotationCheckFailed = errors.New("rotation validation failed")
)
