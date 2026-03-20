package cropplan

import "errors"

var (
	// Общие ошибки
	ErrInvalidName             = errors.New("invalid name")
	ErrInvalidDuration         = errors.New("invalid duration")
	ErrInvalidStageName        = errors.New("invalid stage name")
	ErrInvalidTemperatureRange = errors.New("invalid temperature range")
	ErrInvalidPHRange          = errors.New("invalid pH range")
	ErrInvalidHumidityRange    = errors.New("invalid humidity range")
	ErrInvalidMinYears         = errors.New("invalid minimum years")
	ErrInvalidPredecessor      = errors.New("invalid predecessor")
	ErrInvalidCropType         = errors.New("invalid crop type")
	ErrInvalidVarietyName      = errors.New("invalid variety name")
	ErrInvalidVegetationDays   = errors.New("invalid vegetation days")

	// Ошибки плана
	ErrCannotModifyPublished        = errors.New("cannot modify published plan")
	ErrAlreadyPublished             = errors.New("plan already published")
	ErrOnlyPublishedCanBeDeprecated = errors.New("only published plan can be deprecated")
	ErrOnlyPublishedCanBeVersioned  = errors.New("only published plan can be versioned")
	ErrNoStages                     = errors.New("plan has no stages")
	ErrStageNotFound                = errors.New("stage not found")
	ErrStageOrderDuplicate          = errors.New("duplicate stage order")
	ErrStageDurationMismatch        = errors.New("stages duration does not match plan duration")
	ErrRotationRuleDuplicate        = errors.New("rotation rule already exists")
)
