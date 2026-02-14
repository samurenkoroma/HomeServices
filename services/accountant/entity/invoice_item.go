package entity

import (
	"errors"
)

type InvoiceItem struct {
	ID       uint
	Name     string
	Quantity int
	Price    Price
}

type Price struct {
	Value float64
	Vat   float64
	Total float64
}

func (p *Price) isValid() error {
	if p.Value < 0 {
		return errors.New("price value is negative")
	}
	if p.Vat < 0 {
		return errors.New("vat value is negative")
	}
	if p.Total < 0 || p.Total != (p.Value+p.Vat) {
		return errors.New("total value is fail")
	}
	return nil
}
func (i InvoiceItem) isValid() error {
	if i.Name == "" {
		return errors.New("name is required")
	}
	if i.Quantity <= 0 {
		return errors.New("quantity is required")
	}
	if err := i.Price.isValid(); err != nil {
		return err
	}
	return nil
}
