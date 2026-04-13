package catalog

func init() {
	RegisterSpecies(Species{
		Key:         "sunflower",
		Name:        "Подсолнечник",
		Family:      "asteraceae",
		Category:    "Масличные",
		ImageUrl:    "https://images.example.com/crops/sunflower.jpg",
		Description: "Масличная культура семейства астровых. Семена используются для производства масла.",
	})

	// Лакомка
	lakomka := Variety{
		ID:                 "lakomka",
		Name:               "Лакомка",
		SpeciesKey:         "sunflower",
		SpeciesName:        "Подсолнечник",
		BaseTemperature:    8.0,
		MaxTemperature:     35.0,
		DaysToMaturity:     105,
		YieldPotential:     0.25, // 2.5 т/га семян
		PlantHeight:        1.6,
		RecommendedSeasons: []string{"spring"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"oil_content":        "48-50%",
			"use":                "кондитерский",
			"drought_resistance": "высокая",
			"seed_size":          "крупный",
		},
		Description: "Кондитерский сорт с крупными семенами",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 100},
			{Code: "BBCH-19", GDDRequired: 250},
			{Code: "BBCH-51", GDDRequired: 500},
			{Code: "BBCH-61", GDDRequired: 700},
			{Code: "BBCH-71", GDDRequired: 900},
			{Code: "BBCH-89", GDDRequired: 1150},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.7,
				PlantSpacing:    0.25,
				SowingDepth:     6.0,
				GerminationRate: 90,
				SafetyFactor:    1.15,
			},
		},
	}

	// Оливер (масличный)
	oliver := Variety{
		ID:                 "oliver",
		Name:               "Оливер",
		SpeciesKey:         "sunflower",
		SpeciesName:        "Подсолнечник",
		BaseTemperature:    8.0,
		MaxTemperature:     35.0,
		DaysToMaturity:     110,
		YieldPotential:     0.3,
		PlantHeight:        1.7,
		RecommendedSeasons: []string{"spring"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"oil_content":        "52-54%",
			"use":                "масличный",
			"drought_resistance": "высокая",
		},
		Description: "Высокомасличный гибрид для производства масла",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 105},
			{Code: "BBCH-19", GDDRequired: 260},
			{Code: "BBCH-51", GDDRequired: 520},
			{Code: "BBCH-61", GDDRequired: 720},
			{Code: "BBCH-71", GDDRequired: 930},
			{Code: "BBCH-89", GDDRequired: 1200},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.7,
				PlantSpacing:    0.25,
				SowingDepth:     6.0,
				GerminationRate: 92,
				SafetyFactor:    1.12,
			},
		},
	}

	RegisterVariety("sunflower", lakomka)
	RegisterVariety("sunflower", oliver)
}
