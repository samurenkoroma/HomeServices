package valueobject

type Dimension struct {
	Length *float64
	Width  *float64
	Height *float64
}

func (d *Dimension) Area() float64 {
	return *d.Length * *d.Width
}
