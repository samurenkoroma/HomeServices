package facility

import "errors"

var (
	ErrAreaExceeded          = errors.New("area exceeded")
	ErrInvalidSpaceType      = errors.New("invalid land space type")
	ErrSectionNotFound       = errors.New("section not found")
	ErrLandUnitAlreadyExists = errors.New("land unit already exists")

	ErrBlockNotAllowed        = errors.New("blocks allowed only in field facility")
	ErrDuplicateArea          = errors.New("duplicate growing area")
	ErrAreaNotFound           = errors.New("growing area not found")
	ErrBedMustHaveParentBlock = errors.New("bed must belong to block")
	ErrInvalidFacilityType    = errors.New("invalid facility type")
	ErrBedNotAllowed          = errors.New("direct bed allowed only in greenhouse facility")
)
