package inmemory

import "errors"

var (
	ErrDuplicatePlan = errors.New("plan with this ID already exists")
	ErrDuplicateTask = errors.New("task with this ID already exists")
)
