package catalog

func init() {
	RegisterSpecies(Species{
		Key:         "barley",
		Name:        "Ячмень",
		Family:      "poaceae",
		Category:    "Зерновые",
		ImageUrl:    "https://images.example.com/crops/barley.jpg",
		Description: "Зерновая культура семейства злаковых. Используется для производства крупы, пива и кормов.",
	})

	// Вакула
	vakula := Variety{
		ID:                 "vakula",
		Name:               "Вакула",
		SpeciesKey:         "barley",
		SpeciesName:        "Ячмень",
		BaseTemperature:    3.0,
		MaxTemperature:     30.0,
		DaysToMaturity:     85,
		YieldPotential:     0.55,
		PlantHeight:        0.8,
		RecommendedSeasons: []string{"spring"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"use":                "пивоваренный",
			"drought_resistance": "высокая",
			"protein":            "10-12%",
		},
		Description: "Пивоваренный ячмень с высоким качеством зерна",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 55},
			{Code: "BBCH-19", GDDRequired: 150},
			{Code: "BBCH-30", GDDRequired: 300},
			{Code: "BBCH-51", GDDRequired: 480},
			{Code: "BBCH-61", GDDRequired: 600},
			{Code: "BBCH-71", GDDRequired: 750},
			{Code: "BBCH-89", GDDRequired: 950},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.15,
				PlantSpacing:    0.03,
				SowingDepth:     4.0,
				GerminationRate: 88,
				SafetyFactor:    1.18,
			},
		},
	}

	RegisterVariety("barley", vakula)
}
