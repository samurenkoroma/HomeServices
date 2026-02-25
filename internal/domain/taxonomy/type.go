package taxonomy

type TypeTaxonomy uint8

var values = map[TypeTaxonomy]string{
	Animals:    "animals",
	Plants:     "plants",
	Tools:      "tools",
	Equipments: "equipments",
	None:       "none",
}

const (
	None TypeTaxonomy = iota
	Animals
	Plants
	Tools
	Equipments
)

func (t TypeTaxonomy) String() string {
	return values[t]
}

func TypeFromString(name string) uint8 {
	for k, v := range values {
		if v == name {
			return uint8(k)
		}
	}
	return 0
}
