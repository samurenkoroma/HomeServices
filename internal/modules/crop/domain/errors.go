package domain

import "errors"

var (
	ErrStageNotFound        = errors.New("stage not found")
	ErrStageAlreadyStarted  = errors.New("stage already started")
	ErrStageNotInProgress   = errors.New("stage is not in progress")
	ErrStageAlreadyFinished = errors.New("stage already finished")
	ErrInvalidStageOrder    = errors.New("invalid stage order")
	ErrStageCannotBeSkipped = errors.New("stage cannot be skipped")

	ErrInvalidDuration       = errors.New("invalid duration")
	ErrStageOrderDuplicate   = errors.New("duplicate stage order")
	ErrStageDurationMismatch = errors.New("stages duration mismatch")
	ErrCannotModifyPublished = errors.New("cannot modify published plan")
	ErrRotationDuplicate     = errors.New("rotation rule duplicate")
)
