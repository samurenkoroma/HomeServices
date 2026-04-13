package catalog

func init() {
	RegisterSpecies(Species{
		Key:         "rapeseed",
		Name:        "Рапс",
		Family:      "brassicaceae",
		Category:    "Масличные",
		ImageUrl:    "https://images.example.com/crops/rapeseed.jpg",
		Description: "Масличная культура семейства капустных. Используется для производства масла и биодизеля.",
	})

	// Лидер (яровой)
	lider := Variety{
		ID:                 "lider",
		Name:               "Лидер",
		SpeciesKey:         "rapeseed",
		SpeciesName:        "Рапс",
		BaseTemperature:    3.0,
		MaxTemperature:     30.0,
		DaysToMaturity:     95,
		YieldPotential:     0.25,
		PlantHeight:        1.2,
		RecommendedSeasons: []string{"spring"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"oil_content":    "44-46%",
			"use":            "продовольственный",
			"erucic_acid":    "<0.5%",
			"glucosinolates": "<20",
		},
		Description: "Яровой рапс с низким содержанием эруковой кислоты",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 80},
			{Code: "BBCH-19", GDDRequired: 200},
			{Code: "BBCH-51", GDDRequired: 400},
			{Code: "BBCH-61", GDDRequired: 550},
			{Code: "BBCH-71", GDDRequired: 700},
			{Code: "BBCH-89", GDDRequired: 950},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.15,
				PlantSpacing:    0.05,
				SowingDepth:     2.5,
				GerminationRate: 85,
				SafetyFactor:    1.2,
			},
		},
	}

	RegisterVariety("rapeseed", lider)
}
