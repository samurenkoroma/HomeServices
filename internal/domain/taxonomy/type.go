package taxonomy

type TypeTaxonomy uint

var values = map[TypeTaxonomy]string{
	Animals: "Animals",
	Plants:  "Plants",
	Tools:   "Tools",
	None:    "None",
}

const (
	None    = TypeTaxonomy(iota)
	Animals = TypeTaxonomy(1)
	Plants  = TypeTaxonomy(2)
	Tools   = TypeTaxonomy(3)
)

func (t TypeTaxonomy) String() string {
	return values[t]
}
