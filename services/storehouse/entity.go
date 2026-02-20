package storehouse

type SeedVariant struct {
	Price  float32
	Weight float32
}

type Plant struct {
	ID   uint
	Name string
}

type Vendor struct {
	ID   uint
	Name string
	URL  string
}
type Seed struct {
	ID       uint
	Plant    *Plant
	Vendor   *Vendor
	Variants []*SeedVariant
}
