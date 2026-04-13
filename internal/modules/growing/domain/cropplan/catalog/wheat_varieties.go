package catalog

func init() {
	RegisterSpecies(Species{
		Key:         "wheat",
		Name:        "Пшеница",
		Family:      "poaceae",
		Category:    "Зерновые",
		ImageUrl:    "https://images.example.com/crops/wheat.jpg",
		Description: "Зерновая культура семейства злаковых. Основной хлебный злак.",
	})

	// Безостая 1 (озимая)
	bezostaya1 := Variety{
		ID:                 "bezostaya_1",
		Name:               "Безостая 1",
		SpeciesKey:         "wheat",
		SpeciesName:        "Пшеница",
		BaseTemperature:    3.0,
		MaxTemperature:     28.0,
		DaysToMaturity:     280,
		YieldPotential:     0.6,
		PlantHeight:        1.0,
		RecommendedSeasons: []string{"autumn"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"grain_type":       "мягкая",
			"use":              "хлебопекарная",
			"winter_hardiness": "средняя",
			"protein":          "12-14%",
		},
		Description: "Классический сорт озимой пшеницы",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 50},
			{Code: "BBCH-19", GDDRequired: 150},
			{Code: "BBCH-30", GDDRequired: 300},
			{Code: "BBCH-51", GDDRequired: 500},
			{Code: "BBCH-61", GDDRequired: 650},
			{Code: "BBCH-71", GDDRequired: 800},
			{Code: "BBCH-89", GDDRequired: 1000},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.15,
				PlantSpacing:    0.03,
				SowingDepth:     5.0,
				GerminationRate: 90,
				SafetyFactor:    1.15,
			},
		},
	}

	// Саратовская 29 (яровая)
	saratovskaya29 := Variety{
		ID:                 "saratovskaya_29",
		Name:               "Саратовская 29",
		SpeciesKey:         "wheat",
		SpeciesName:        "Пшеница",
		BaseTemperature:    3.0,
		MaxTemperature:     30.0,
		DaysToMaturity:     100,
		YieldPotential:     0.5,
		PlantHeight:        0.9,
		RecommendedSeasons: []string{"spring"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"grain_type":         "твердая",
			"use":                "макаронная",
			"drought_resistance": "высокая",
			"protein":            "14-16%",
		},
		Description: "Яровая твердая пшеница для макарон",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 60},
			{Code: "BBCH-19", GDDRequired: 160},
			{Code: "BBCH-30", GDDRequired: 320},
			{Code: "BBCH-51", GDDRequired: 520},
			{Code: "BBCH-61", GDDRequired: 670},
			{Code: "BBCH-71", GDDRequired: 820},
			{Code: "BBCH-89", GDDRequired: 1050},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.15,
				PlantSpacing:    0.03,
				SowingDepth:     5.0,
				GerminationRate: 92,
				SafetyFactor:    1.12,
			},
		},
	}

	RegisterVariety("wheat", bezostaya1)
	RegisterVariety("wheat", saratovskaya29)
}
