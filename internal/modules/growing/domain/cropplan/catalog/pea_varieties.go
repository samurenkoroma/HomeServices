package catalog

func init() {
	RegisterSpecies(Species{
		Key:         "pea",
		Name:        "Горох",
		Family:      "fabaceae",
		Category:    "Бобовые",
		ImageUrl:    "https://images.example.com/crops/pea.jpg",
		Description: "Бобовая культура семейства бобовых. Ценный источник белка.",
	})

	// Атлант
	atlant := Variety{
		ID:                 "atlant",
		Name:               "Атлант",
		SpeciesKey:         "pea",
		SpeciesName:        "Горох",
		BaseTemperature:    4.0,
		MaxTemperature:     28.0,
		DaysToMaturity:     85,
		YieldPotential:     0.35,
		PlantHeight:        0.8,
		RecommendedSeasons: []string{"spring"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"protein":            "22-24%",
			"use":                "продовольственный",
			"lodging_resistance": "высокая",
			"seed_size":          "крупный",
		},
		Description: "Лущильный сорт с высоким содержанием белка",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 70},
			{Code: "BBCH-19", GDDRequired: 180},
			{Code: "BBCH-51", GDDRequired: 350},
			{Code: "BBCH-61", GDDRequired: 480},
			{Code: "BBCH-71", GDDRequired: 650},
			{Code: "BBCH-89", GDDRequired: 850},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.15,
				PlantSpacing:    0.05,
				SowingDepth:     5.0,
				GerminationRate: 88,
				SafetyFactor:    1.18,
			},
		},
	}

	RegisterVariety("pea", atlant)
}
