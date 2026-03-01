package valueobject

import (
	"errors"
)

var ErrInvalidDimension = errors.New("invalid dimension")

type Dimension struct {
	length float64
	width  float64
}

func NewDimension(length, width float64) (Dimension, error) {
	if length <= 0 || width <= 0 {
		return Dimension{}, ErrInvalidDimension
	}
	return Dimension{length: length, width: width}, nil
}
func (d Dimension) Length() float64 {
	return d.length
}

func (d Dimension) Width() float64 {
	return d.width
}
func (d Dimension) Area() float64 {
	return d.length * d.width
}
