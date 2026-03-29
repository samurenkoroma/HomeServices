package types

type Dimension struct {
	Length *float64 `json:"length"`
	Width  *float64 `json:"width"`
	Height *float64 `json:"height"`
}

func NewDimension(length, width, height float64) *Dimension {
	return &Dimension{
		Length: &length,
		Width:  &width,
		Height: &height,
	}
}

func (d *Dimension) Area() float64 {
	return *d.Length * *d.Width
}
