package valueobject

import (
	"errors"
)

type Dimension struct {
	Length float64
	Width  float64
}

func NewDimension(length, width float64) (Dimension, error) {
	if length <= 0 || width <= 0 {
		return Dimension{}, errors.New("invalid dimension")
	}
	return Dimension{Length: length, Width: width}, nil
}

func (d Dimension) Area() float64 {
	return d.Length * d.Width
}
