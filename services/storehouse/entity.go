package storehouse

import "samurenkoroma/services/services/accountant/entity"

type SeedVariant struct {
	Price  float32
	Weight float32
}

type Plant struct {
	ID   uint
	Name string
}

type Seed struct {
	ID       uint
	Plant    *Plant
	Supplier *entity.Supplier
	Variants []*SeedVariant
}
