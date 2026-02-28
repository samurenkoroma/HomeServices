package landunit

import "errors"

var (
	ErrAreaExceeded          = errors.New("area exceeded")
	ErrInvalidSpaceType      = errors.New("invalid land space type")
	ErrSectionNotFound       = errors.New("section not found")
	ErrLandUnitAlreadyExists = errors.New("land unit already exists")
)
