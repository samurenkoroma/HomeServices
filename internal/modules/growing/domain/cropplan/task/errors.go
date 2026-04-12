package task

import "errors"

var (
	ErrTaskNotFound      = errors.New("task not found")
	ErrTaskAlreadyDone   = errors.New("task already completed")
	ErrInvalidStatus     = errors.New("invalid task status")
	ErrUnauthorized      = errors.New("not authorized to modify this task")
	ErrInvalidAssignment = errors.New("invalid assignment")
)
