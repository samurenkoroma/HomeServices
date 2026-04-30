package season

import "errors"

var (
	ErrInvalidStatusTransition = errors.New("invalid status transition")
	ErrSeasonNotStarted        = errors.New("season not started")

	ErrInvalidName           = errors.New("invalid season name")
	ErrInvalidPeriod         = errors.New("start date must be before end date")
	ErrInvalidPlanningInPast = errors.New("invalid planning in past")
	ErrInvalidCreatedBy      = errors.New("created_by is required")
	ErrSeasonNotFound        = errors.New("season not found")
	ErrSeasonAlreadyExists   = errors.New("season with this name already exists")
	ErrSeasonOverlap         = errors.New("season overlaps with existing season")
	ErrAlreadyArchived       = errors.New("season already archived")

	// Ошибки планирования
	ErrAreaAlreadyAllocated = errors.New("area already allocated for this season")
	ErrAllocationNotFound   = errors.New("allocation not found")
	ErrInvalidAllocation    = errors.New("invalid allocation")
)
