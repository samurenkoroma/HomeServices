package spatial

import "errors"

var (
	ErrInvalidGeometry     = errors.New("invalid geometry")
	ErrUnsupportedGeometry = errors.New("unsupported geometry")
)
