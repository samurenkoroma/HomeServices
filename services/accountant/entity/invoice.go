package entity

type Invoice struct {
	ID       uint
	Items    []*InvoiceItem
	File     string
	Supplier *Supplier
}

func (i *Invoice) AttachItem(item InvoiceItem) bool {
	if err := item.isValid(); err != nil {
		return false
	}

	i.Items = append(i.Items, &item)
	return true
}

func NewInvoice(file string, supplier *Supplier) (*Invoice, error) {
	return &Invoice{
		File:     file,
		Supplier: supplier,
	}, nil
}
