package finance

type Invoice struct {
	Items       []InvoiceItem
	File        string
	Description string
}

type Supplier struct {
}
type Price struct {
	Value float64
	Vat   float64
	Total float64
}

type Measurement int

const (
	Kg Measurement = iota
	Unit
)

type Amount struct {
	Unit  Measurement
	Count float64
}

type InvoiceItem struct {
	Name   string
	Amount Amount
	Price  Price
}
