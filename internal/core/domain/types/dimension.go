package types

type Dimension struct {
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
}

func NewDimension(length, width float64) *Dimension {
	return &Dimension{
		Length: length,
		Width:  width,
	}
}

func (d *Dimension) AreaInHectares() float64 {
	return (d.Length * d.Width) / 10000
}
