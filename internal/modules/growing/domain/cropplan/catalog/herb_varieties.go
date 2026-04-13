package catalog

func init() {
	// ========== РЕГИСТРИРУЕМ ВИДЫ ЗЕЛЕНИ ==========

	// УКРОП
	RegisterSpecies(Species{
		Key:         "dill",
		Name:        "Укроп",
		Family:      "apiaceae",
		Category:    "Овощные",
		Description: "Пряная зелень семейства зонтичных",
	})

	// ПЕТРУШКА
	RegisterSpecies(Species{
		Key:         "parsley",
		Name:        "Петрушка",
		Family:      "apiaceae",
		Category:    "Овощные",
		Description: "Двулетнее растение семейства зонтичных",
	})

	// САЛАТ
	RegisterSpecies(Species{
		Key:         "lettuce",
		Name:        "Салат",
		Family:      "asteraceae",
		Category:    "Овощные",
		ImageUrl:    "https://images.unsplash.com/photo-1597848212624-a19c35a2651d?w=400&h=300&fit=crop",
		Description: "Листовая овощная культура",
	})

	// ========== 1. УКРОП "ГРИБОВСКИЙ" ==========
	dillGribovsky := Variety{
		ID:          "dill_gribovsky",
		Name:        "Укроп Грибовский",
		SpeciesKey:  "dill",
		SpeciesName: "Укроп",

		// Температуры (холодостойкий)
		BaseTemperature: 3.0,
		MaxTemperature:  25.0,

		DaysToMaturity: 30,
		YieldPotential: 3.5,
		PlantHeight:    0.4,

		RecommendedSeasons: []string{"spring", "summer", "autumn"},
		GrowingTypes:       []string{"open_ground", "greenhouse"},

		Characteristics: map[string]string{
			"aroma":           "сильный",
			"leaf_color":      "зеленый",
			"use":             "универсальный",
			"frost_resistant": "высокая",
		},
		Description: "Раннеспелый сорт для открытого грунта",

		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", Name: "Всходы", GDDRequired: 50},
			{Code: "BBCH-19", Name: "Розетка листьев", GDDRequired: 120},
			{Code: "BBCH-51", Name: "Стеблевание", GDDRequired: 200},
			{Code: "BBCH-61", Name: "Цветение", GDDRequired: 300},
			{Code: "BBCH-89", Name: "Созревание семян", GDDRequired: 450},
		},

		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.25,
				PlantSpacing:    0.05,
				SowingDepth:     1.5,
				GerminationRate: 75,
				SafetyFactor:    1.3,
			},
		},
	}

	// ========== 2. УКРОП "КИБРАЙ" ==========
	dillKibray := Variety{
		ID:                 "dill_kibray",
		Name:               "Укроп Кибрей",
		SpeciesKey:         "dill",
		SpeciesName:        "Укроп",
		BaseTemperature:    3.0,
		MaxTemperature:     25.0,
		DaysToMaturity:     40,
		YieldPotential:     4.5,
		PlantHeight:        0.6,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"aroma":      "средний",
			"leaf_color": "сизо-зеленый",
			"use":        "свежий, заморозка",
		},
		Description: "Позднеспелый урожайный сорт",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 60},
			{Code: "BBCH-19", GDDRequired: 150},
			{Code: "BBCH-51", GDDRequired: 250},
			{Code: "BBCH-61", GDDRequired: 350},
			{Code: "BBCH-89", GDDRequired: 500},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.3,
				PlantSpacing:    0.05,
				SowingDepth:     1.5,
				GerminationRate: 75,
				SafetyFactor:    1.3,
			},
		},
	}

	// ========== 3. ПЕТРУШКА "ОБЫКНОВЕННАЯ" ==========
	parsleyCommon := Variety{
		ID:                 "parsley_common",
		Name:               "Петрушка обыкновенная",
		SpeciesKey:         "parsley",
		SpeciesName:        "Петрушка",
		BaseTemperature:    3.0,
		MaxTemperature:     25.0,
		DaysToMaturity:     35,
		YieldPotential:     4.0,
		PlantHeight:        0.3,
		RecommendedSeasons: []string{"spring", "summer", "autumn"},
		GrowingTypes:       []string{"open_ground", "greenhouse"},
		Characteristics: map[string]string{
			"flavor":          "пряный",
			"leaf_type":       "гладкий",
			"use":             "универсальный",
			"frost_resistant": "высокая",
		},
		Description: "Классический сорт листовой петрушки",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 70},
			{Code: "BBCH-19", GDDRequired: 180},
			{Code: "BBCH-51", GDDRequired: 300},
			{Code: "BBCH-61", GDDRequired: 420},
			{Code: "BBCH-89", GDDRequired: 550},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.25,
				PlantSpacing:    0.05,
				SowingDepth:     1.5,
				GerminationRate: 70,
				SafetyFactor:    1.35,
			},
		},
	}

	// ========== 4. ПЕТРУШКА "КУДРЯВАЯ" ==========
	parsleyCurly := Variety{
		ID:                 "parsley_curly",
		Name:               "Петрушка кудрявая",
		SpeciesKey:         "parsley",
		SpeciesName:        "Петрушка",
		BaseTemperature:    3.0,
		MaxTemperature:     25.0,
		DaysToMaturity:     40,
		YieldPotential:     3.5,
		PlantHeight:        0.35,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground", "greenhouse"},
		Characteristics: map[string]string{
			"flavor":    "нежный",
			"leaf_type": "кудрявый",
			"use":       "украшение, салаты",
		},
		Description: "Декоративная кудрявая петрушка",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 75},
			{Code: "BBCH-19", GDDRequired: 190},
			{Code: "BBCH-51", GDDRequired: 320},
			{Code: "BBCH-61", GDDRequired: 440},
			{Code: "BBCH-89", GDDRequired: 570},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.3,
				PlantSpacing:    0.05,
				SowingDepth:     1.5,
				GerminationRate: 70,
				SafetyFactor:    1.35,
			},
		},
	}

	// ========== 5. САЛАТ "МОСКОВСКИЙ" ==========
	lettuceMoskovsky := Variety{
		ID:                 "lettuce_moskovsky",
		Name:               "Салат Московский",
		SpeciesKey:         "lettuce",
		SpeciesName:        "Салат",
		BaseTemperature:    4.0,
		MaxTemperature:     25.0,
		DaysToMaturity:     35,
		YieldPotential:     4.5,
		PlantHeight:        0.25,
		RecommendedSeasons: []string{"spring", "summer", "autumn"},
		GrowingTypes:       []string{"open_ground", "greenhouse"},
		Characteristics: map[string]string{
			"leaf_color": "светло-зеленый",
			"leaf_shape": "лопастный",
			"use":        "салаты",
		},
		Description: "Скороспелый листовой салат",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 60},
			{Code: "BBCH-19", GDDRequired: 160},
			{Code: "BBCH-51", GDDRequired: 280},
			{Code: "BBCH-61", GDDRequired: 400},
			{Code: "BBCH-89", GDDRequired: 520},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.25,
				PlantSpacing:    0.15,
				SowingDepth:     1.0,
				GerminationRate: 80,
				SafetyFactor:    1.25,
			},
		},
	}

	// ========== 6. САЛАТ "ЛОЛЛО РОССА" ==========
	lettuceLolloRossa := Variety{
		ID:                 "lettuce_lollo_rossa",
		Name:               "Салат Лолло Росса",
		SpeciesKey:         "lettuce",
		SpeciesName:        "Салат",
		BaseTemperature:    4.0,
		MaxTemperature:     25.0,
		DaysToMaturity:     45,
		YieldPotential:     4.0,
		PlantHeight:        0.2,
		RecommendedSeasons: []string{"spring", "summer", "autumn"},
		GrowingTypes:       []string{"open_ground", "greenhouse"},
		Characteristics: map[string]string{
			"leaf_color": "красно-бордовый",
			"leaf_shape": "кудрявый",
			"use":        "украшение, салаты",
		},
		Description: "Декоративный салат с красными листьями",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 65},
			{Code: "BBCH-19", GDDRequired: 170},
			{Code: "BBCH-51", GDDRequired: 290},
			{Code: "BBCH-61", GDDRequired: 410},
			{Code: "BBCH-89", GDDRequired: 530},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.3,
				PlantSpacing:    0.2,
				SowingDepth:     1.0,
				GerminationRate: 80,
				SafetyFactor:    1.25,
			},
		},
	}

	// ========== 7. БАЗИЛИК "ФИОЛЕТОВЫЙ" ==========
	basilPurple := Variety{
		ID:                 "basil_purple",
		Name:               "Базилик фиолетовый",
		SpeciesKey:         "basil",
		SpeciesName:        "Базилик",
		BaseTemperature:    8.0, // теплолюбивый
		MaxTemperature:     30.0,
		DaysToMaturity:     35,
		YieldPotential:     3.0,
		PlantHeight:        0.3,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"greenhouse"},
		Characteristics: map[string]string{
			"aroma":      "пряный",
			"leaf_color": "фиолетовый",
			"use":        "специи, салаты",
		},
		Description: "Ароматный базилик для теплиц",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 100},
			{Code: "BBCH-19", GDDRequired: 220},
			{Code: "BBCH-51", GDDRequired: 350},
			{Code: "BBCH-61", GDDRequired: 480},
			{Code: "BBCH-89", GDDRequired: 600},
		},
		SeedingRates: map[string]SeedingRate{
			"greenhouse": {
				GrowingType:     "greenhouse",
				RowSpacing:      0.3,
				PlantSpacing:    0.15,
				SowingDepth:     1.0,
				GerminationRate: 85,
				SafetyFactor:    1.2,
			},
		},
	}

	// Регистрируем все сорта
	RegisterVariety("dill", dillGribovsky)
	RegisterVariety("dill", dillKibray)
	RegisterVariety("parsley", parsleyCommon)
	RegisterVariety("parsley", parsleyCurly)
	RegisterVariety("lettuce", lettuceMoskovsky)
	RegisterVariety("lettuce", lettuceLolloRossa)
	RegisterVariety("basil", basilPurple)
}
