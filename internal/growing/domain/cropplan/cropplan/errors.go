package cropplan

import "errors"

var (
	// ... существующие ошибки ...

	// Stage errors
	ErrStageNotFound        = errors.New("stage not found")
	ErrStageAlreadyStarted  = errors.New("stage already started")
	ErrStageNotInProgress   = errors.New("stage is not in progress")
	ErrStageAlreadyFinished = errors.New("stage already finished")
	ErrStageOrderDuplicate  = errors.New("stage with this order already exists")
	ErrInvalidStageOrder    = errors.New("invalid stage order")
	ErrStageCannotBeSkipped = errors.New("stage cannot be skipped")
)
