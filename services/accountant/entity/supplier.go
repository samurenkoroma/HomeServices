package entity

import (
	"errors"
)

type Supplier struct {
	ID     uint
	Name   string
	Site   string
	Rating float64
}

func NewSupplier(name string, site string) (*Supplier, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	return &Supplier{
		Name:   name,
		Rating: .0,
		Site:   site,
	}, nil
}
