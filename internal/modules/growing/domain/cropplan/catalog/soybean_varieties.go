package catalog

func init() {
	RegisterSpecies(Species{
		Key:         "soybean",
		Name:        "Соя",
		Family:      "fabaceae",
		Category:    "Бобовые",
		ImageUrl:    "https://images.example.com/crops/soybean.jpg",
		Description: "Бобовая культура семейства бобовых. Ценный источник растительного белка и масла.",
	})

	// Вилана
	vilana := Variety{
		ID:                 "vilana",
		Name:               "Вилана",
		SpeciesKey:         "soybean",
		SpeciesName:        "Соя",
		BaseTemperature:    10.0,
		MaxTemperature:     35.0,
		DaysToMaturity:     110,
		YieldPotential:     0.3,
		PlantHeight:        0.9,
		RecommendedSeasons: []string{"spring"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"protein":     "38-40%",
			"oil_content": "20-22%",
			"use":         "продовольственный",
		},
		Description: "Раннеспелый сорт сои",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 120},
			{Code: "BBCH-19", GDDRequired: 280},
			{Code: "BBCH-51", GDDRequired: 550},
			{Code: "BBCH-61", GDDRequired: 700},
			{Code: "BBCH-71", GDDRequired: 900},
			{Code: "BBCH-89", GDDRequired: 1200},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.45,
				PlantSpacing:    0.05,
				SowingDepth:     4.0,
				GerminationRate: 85,
				SafetyFactor:    1.2,
			},
		},
	}

	RegisterVariety("soybean", vilana)
}
