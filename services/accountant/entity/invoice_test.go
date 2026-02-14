package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestCreateInvoice(t *testing.T) {
	supplier, _ := NewSupplier(uuid.New().String(), "test", "http://evo.net.ua")
	testcases := []struct {
		text   string
		input  *Invoice
		output *Invoice
	}{
		{
			text: "",
			input: &Invoice{
				Items:    nil,
				File:     "test.pdf",
				Supplier: supplier,
			},
			output: &Invoice{},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.text, func(t *testing.T) {

		})
	}

}
