package landunit

import "errors"

var (
	ErrAreaExceeded          = errors.New("area exceeded")
	ErrInvalidUnitType       = errors.New("invalid land unit type")
	ErrSectionNotFound       = errors.New("section not found")
	ErrLandUnitAlreadyExists = errors.New("land unit already exists")
)
