package postgres

import (
	"errors"
)

var (
	ErrConcurrentModification = errors.New("concurrent modification detected")
	ErrSeasonAlreadyExists    = errors.New("season already exists")
	ErrSeasonNotFound         = errors.New("season not found")
)
