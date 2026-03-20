package postgres

import (
	"errors"

	"github.com/lib/pq"
)

var ErrConcurrentModification = errors.New("concurrent modification detected")

func isUniqueViolation(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code == "23505"
	}
	return false
}
