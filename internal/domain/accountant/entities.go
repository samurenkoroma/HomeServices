package accountant

import (
	"errors"
)

type Invoice struct {
	Items    []InvoiceRow
	File     string
	Supplier Supplier
}

type Price struct {
	Value float64
	Vat   float64
	Total float64
}

type InvoiceRow struct {
	Name   string
	Amount int
	Price  Price
}

type Supplier struct {
	ID     string
	Name   string
	Rating float64
}

func NewSupplier(id string, name string) (*Supplier, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	return &Supplier{
		ID:     id,
		Name:   name,
		Rating: .0,
	}, nil
}
