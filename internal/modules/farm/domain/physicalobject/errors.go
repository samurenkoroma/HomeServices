package physicalobject

import "errors"

var (
	// ... существующие ошибки ...

	ErrAlreadyActive          = errors.New("object is already active")
	ErrAlreadyInactive        = errors.New("object is already inactive")
	ErrInvalidGeometry        = errors.New("invalid geometry")
	ErrEmptyGeometry          = errors.New("geometry cannot be empty")
	ErrPhysicalObjectNotFound = errors.New("physical object not found")
)
