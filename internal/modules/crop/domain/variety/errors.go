package variety

import "errors"

var (
	ErrNotFound           = errors.New("variety not found")
	ErrEmptyCropType      = errors.New("empty crop type")
	ErrInvalidVarietyName = errors.New("invalid variety name")
)
