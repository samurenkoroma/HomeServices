package postgres

import (
	"errors"
)

var ErrConcurrentModification = errors.New("concurrent modification detected")
