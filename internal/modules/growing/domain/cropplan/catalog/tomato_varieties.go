package catalog

func init() {
	// ========== РЕГИСТРИРУЕМ ВИД ==========
	RegisterSpecies(Species{
		Key:         "tomato",
		Name:        "Томат",
		Family:      "nightshade",
		Description: "Овощная культура семейства пасленовых",
	})

	// ========== СОРТ "БЫЧЬЕ СЕРДЦЕ" ==========
	bullHeart := Variety{
		ID:                 "bull_heart",
		Name:               "Бычье сердце",
		SpeciesKey:         "tomato",
		SpeciesName:        "Томат",
		DaysToMaturity:     120,
		YieldPotential:     8.5,
		PlantHeight:        1.8,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground", "greenhouse"},
		Characteristics: map[string]string{
			"fruit_weight": "300-500g",
			"fruit_color":  "малиновый",
			"taste":        "сладкий",
		},
		Description: "Крупноплодный сорт для салатов",

		// GDD шкала (базовая температура 10°C)
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", Name: "Первый настоящий лист", GDDRequired: 120, Description: "Появление первого листа", IsCritical: false},
			{Code: "BBCH-19", Name: "9 и более листьев", GDDRequired: 350, Description: "Активный рост", IsCritical: false},
			{Code: "BBCH-51", Name: "Бутонизация", GDDRequired: 600, Description: "Появление бутонов", IsCritical: false},
			{Code: "BBCH-61", Name: "Цветение", GDDRequired: 800, Description: "Раскрытие цветов", IsCritical: true},
			{Code: "BBCH-71", Name: "Завязывание плодов", GDDRequired: 950, Description: "Образование завязей", IsCritical: true},
			{Code: "BBCH-81", Name: "Созревание", GDDRequired: 1200, Description: "Плоды начинают краснеть", IsCritical: false},
			{Code: "BBCH-89", Name: "Полная спелость", GDDRequired: 1400, Description: "Плоды готовы к сбору", IsCritical: false},
		},

		// Нормы высева
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.7,
				PlantSpacing:    0.5,
				SowingDepth:     2.0,
				GerminationRate: 85,
				SafetyFactor:    1.2,
			},
			"greenhouse": {
				GrowingType:     "greenhouse",
				RowSpacing:      0.8,
				PlantSpacing:    0.6,
				SowingDepth:     1.5,
				GerminationRate: 90,
				SafetyFactor:    1.1,
			},
		},
	}

	// ========== СОРТ "ЧЕРРИ" ==========
	cherry := Variety{
		ID:                 "cherry",
		Name:               "Черри",
		SpeciesKey:         "tomato",
		SpeciesName:        "Томат",
		DaysToMaturity:     90,
		YieldPotential:     5.0,
		PlantHeight:        1.5,
		RecommendedSeasons: []string{"spring", "summer", "autumn"},
		GrowingTypes:       []string{"greenhouse"},
		Characteristics: map[string]string{
			"fruit_weight": "15-20g",
			"fruit_color":  "красный",
			"taste":        "сладкий",
		},
		Description: "Мелкоплодный, очень урожайный сорт",

		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 100},
			{Code: "BBCH-19", GDDRequired: 280},
			{Code: "BBCH-51", GDDRequired: 480},
			{Code: "BBCH-61", GDDRequired: 650, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 780, IsCritical: true},
			{Code: "BBCH-81", GDDRequired: 950},
			{Code: "BBCH-89", GDDRequired: 1100},
		},

		SeedingRates: map[string]SeedingRate{
			"greenhouse": {
				GrowingType:     "greenhouse",
				RowSpacing:      0.6,
				PlantSpacing:    0.4,
				SowingDepth:     1.0,
				GerminationRate: 92,
				SafetyFactor:    1.1,
			},
		},
	}

	// ========== СОРТ "ДУБОК" (холодостойкий) ==========
	oak := Variety{
		ID:                 "oak",
		Name:               "Дубок",
		SpeciesKey:         "tomato",
		SpeciesName:        "Томат",
		DaysToMaturity:     100,
		YieldPotential:     6.0,
		PlantHeight:        0.8,
		RecommendedSeasons: []string{"spring"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"fruit_weight":   "80-100g",
			"cold_resistant": "yes",
		},
		Description: "Холодостойкий сорт для открытого грунта",

		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 110},
			{Code: "BBCH-19", GDDRequired: 300},
			{Code: "BBCH-51", GDDRequired: 520},
			{Code: "BBCH-61", GDDRequired: 700, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 850, IsCritical: true},
			{Code: "BBCH-81", GDDRequired: 1050},
			{Code: "BBCH-89", GDDRequired: 1250},
		},

		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.6,
				PlantSpacing:    0.4,
				SowingDepth:     2.5,
				GerminationRate: 88,
				SafetyFactor:    1.15,
			},
		},
	}

	// Регистрируем сорта
	RegisterVariety("tomato", bullHeart)
	RegisterVariety("tomato", cherry)
	RegisterVariety("tomato", oak)
}
