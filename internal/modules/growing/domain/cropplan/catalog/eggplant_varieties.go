package catalog

func init() {
	// Регистрируем вид
	RegisterSpecies(Species{
		Key:         "eggplant",
		Name:        "Баклажан",
		Family:      "nightshade",
		Description: "Теплолюбивая овощная культура",
	})

	// Сорт "Алмаз"
	diamond := Variety{
		ID:                 "diamond",
		Name:               "Алмаз",
		SpeciesKey:         "eggplant",
		SpeciesName:        "Баклажан",
		DaysToMaturity:     120,
		YieldPotential:     6.0,
		PlantHeight:        0.6,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground", "greenhouse"},
		Characteristics: map[string]string{
			"fruit_weight": "100-150g",
			"fruit_color":  "темно-фиолетовый",
		},
		Description: "Среднеспелый сорт",

		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 130},
			{Code: "BBCH-19", GDDRequired: 380},
			{Code: "BBCH-51", GDDRequired: 650},
			{Code: "BBCH-61", GDDRequired: 850, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 1000, IsCritical: true},
			{Code: "BBCH-81", GDDRequired: 1300},
			{Code: "BBCH-89", GDDRequired: 1500},
		},

		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.7,
				PlantSpacing:    0.5,
				SowingDepth:     2.0,
				GerminationRate: 80,
				SafetyFactor:    1.2,
			},
			"greenhouse": {
				GrowingType:     "greenhouse",
				RowSpacing:      0.8,
				PlantSpacing:    0.6,
				SowingDepth:     1.5,
				GerminationRate: 85,
				SafetyFactor:    1.15,
			},
		},
	}

	RegisterVariety("eggplant", diamond)
}
