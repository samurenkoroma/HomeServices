package cropcycle

import "errors"

var (
	// Общие ошибки
	ErrCycleNotFound          = errors.New("crop cycle not found")
	ErrCycleAlreadyExists     = errors.New("crop cycle already exists")
	ErrInvalidState           = errors.New("invalid cycle state for this operation")
	ErrInvalidOperation       = errors.New("invalid operation")
	ErrInvalidYield           = errors.New("invalid yield data")
	ErrTemplateNotFound       = errors.New("template not found")
	ErrAreaNotAvailable       = errors.New("area not available for planting")
	ErrSeasonNotActive        = errors.New("season is not active")
	ErrConcurrentModification = errors.New("concurrent modification detected")
)
