package cultivationarea

import "errors"

var (
	// Общие ошибки
	ErrAreaNotFound               = errors.New("cultivation area not found")
	ErrSeasonAlreadyConfigured    = errors.New("season already configured for this area")
	ErrSeasonConfigNotFound       = errors.New("season configuration not found")
	ErrAreaNotConfiguredForSeason = errors.New("area not configured for this season")
	ErrCropPlanMismatch           = errors.New("crop plan does not match configured plan")
	ErrNoCropPlanConfigured       = errors.New("no crop plan configured for this area")
	ErrFieldHasBlocks             = errors.New("field already has blocks")
	ErrFieldNotPolyculture        = errors.New("field is not configured as polyculture")
	ErrNotMonocultureField        = errors.New("field is not monoculture")
	ErrBlockRequiresParent        = errors.New("block requires parent field ID")
	ErrBedRequiresParent          = errors.New("bed requires parent block or greenhouse ID")
	ErrCropPlanRequiredForBlock   = errors.New("crop plan required for block")
	ErrCropPlanRequiredForBed     = errors.New("crop plan required for bed")
	ErrGreenhouseHasMultipleCrops = errors.New("greenhouse can have multiple crops via beds")
	ErrUnknownAreaType            = errors.New("unknown area type")
	ErrInvalidGeometry            = errors.New("invalid geometry")
	ErrInvalidArea                = errors.New("invalid area")
)
