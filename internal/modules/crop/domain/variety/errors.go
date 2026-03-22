package variety

import "errors"

var (
	ErrEmptyCropType      = errors.New("empty crop type")
	ErrInvalidVarietyName = errors.New("invalid variety name")
)
