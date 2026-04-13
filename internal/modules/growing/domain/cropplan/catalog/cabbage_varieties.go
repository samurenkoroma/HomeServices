package catalog

func init() {
	// ========== РЕГИСТРИРУЕМ ВИД ==========
	RegisterSpecies(Species{
		Key:         "cabbage",
		Name:        "Капуста",
		Family:      "brassicaceae",
		Description: "Овощная культура семейства капустных",
		Category:    "Овощные",
		ImageUrl:    "https://media.istockphoto.com/id/187114561/uk/%D1%84%D0%BE%D1%82%D0%BE/%D0%BF%D1%96%D0%B2%D1%82%D0%BE%D1%80%D0%B0-%D1%87%D0%B8%D1%81%D1%82%D0%B8%D1%85-%D0%BA%D0%B0%D0%BF%D1%83%D1%81%D1%82%D0%B8-%D0%BD%D0%B0-%D0%B1%D1%96%D0%BB%D0%BE%D0%BC%D1%83-%D1%82%D0%BB%D1%96.webp?s=2048x2048&w=is&k=20&c=Q2I34XFkxsovr0pM_IfyQuyeBCHhUiX1D_zsKAmFXKA=",
	})

	// ========== 1. СЛАВА 1305 ==========
	slava1305 := Variety{
		ID:                 "slava_1305",
		Name:               "Слава 1305",
		SpeciesKey:         "cabbage",
		SpeciesName:        "Капуста",
		BaseTemperature:    5.0,
		MaxTemperature:     25.0,
		DaysToMaturity:     110,
		YieldPotential:     7.0,
		PlantHeight:        0.4,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"head_weight": "2-3kg",
			"head_color":  "светло-зеленый",
			"head_shape":  "округлый",
			"use":         "свежий, квашение",
			"origin":      "СССР",
		},
		Description: "Классический среднеспелый сорт для квашения",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 150},
			{Code: "BBCH-19", GDDRequired: 400},
			{Code: "BBCH-51", GDDRequired: 700},
			{Code: "BBCH-61", GDDRequired: 900, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 1100, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 1400},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.6,
				PlantSpacing:    0.5,
				SowingDepth:     1.5,
				GerminationRate: 80,
				SafetyFactor:    1.25,
			},
		},
	}

	// ========== 2. МЕГАТОН F1 ==========
	megaton := Variety{
		ID:                 "megaton_f1",
		Name:               "Мегатон F1",
		SpeciesKey:         "cabbage",
		SpeciesName:        "Капуста",
		BaseTemperature:    5.0,
		MaxTemperature:     25.0,
		DaysToMaturity:     120,
		YieldPotential:     12.0,
		PlantHeight:        0.45,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"head_weight": "4-6kg",
			"head_color":  "зеленый",
			"head_shape":  "округлый",
			"use":         "свежий, хранение",
		},
		Description: "Крупноплодный гибрид для длительного хранения",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 160},
			{Code: "BBCH-19", GDDRequired: 430},
			{Code: "BBCH-51", GDDRequired: 750},
			{Code: "BBCH-61", GDDRequired: 950, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 1150, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 1500},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.7,
				PlantSpacing:    0.6,
				SowingDepth:     1.5,
				GerminationRate: 82,
				SafetyFactor:    1.22,
			},
		},
	}

	// ========== 3. ПОДАРОК ==========
	podarok := Variety{
		ID:                 "podarok",
		Name:               "Подарок",
		SpeciesKey:         "cabbage",
		SpeciesName:        "Капуста",
		BaseTemperature:    5.0,
		MaxTemperature:     25.0,
		DaysToMaturity:     115,
		YieldPotential:     8.5,
		PlantHeight:        0.4,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"head_weight": "3-4kg",
			"head_color":  "светло-зеленый",
			"head_shape":  "плоскоокруглый",
			"use":         "универсальный",
		},
		Description: "Хорошо хранится, устойчив к растрескиванию",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 155},
			{Code: "BBCH-19", GDDRequired: 420},
			{Code: "BBCH-51", GDDRequired: 730},
			{Code: "BBCH-61", GDDRequired: 930, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 1130, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 1450},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.6,
				PlantSpacing:    0.5,
				SowingDepth:     1.5,
				GerminationRate: 81,
				SafetyFactor:    1.23,
			},
		},
	}

	// ========== 4. МЕНЗА F1 ==========
	menza := Variety{
		ID:                 "menza_f1",
		Name:               "Менза F1",
		SpeciesKey:         "cabbage",
		SpeciesName:        "Капуста",
		BaseTemperature:    5.0,
		MaxTemperature:     25.0,
		DaysToMaturity:     125,
		YieldPotential:     11.0,
		PlantHeight:        0.5,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"head_weight": "4-5kg",
			"head_color":  "сине-зеленый",
			"head_shape":  "округлый",
			"use":         "свежий, квашение",
		},
		Description: "Голландский гибрид с высокими вкусовыми качествами",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 165},
			{Code: "BBCH-19", GDDRequired: 440},
			{Code: "BBCH-51", GDDRequired: 770},
			{Code: "BBCH-61", GDDRequired: 970, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 1170, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 1520},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.7,
				PlantSpacing:    0.55,
				SowingDepth:     1.5,
				GerminationRate: 83,
				SafetyFactor:    1.2,
			},
		},
	}

	// ========== 5. КРАУТМАН АМАГЕР ==========
	krautman := Variety{
		ID:                 "krautman_amager",
		Name:               "Краутман Амагер",
		SpeciesKey:         "cabbage",
		SpeciesName:        "Капуста",
		BaseTemperature:    5.0,
		MaxTemperature:     25.0,
		DaysToMaturity:     130,
		YieldPotential:     9.0,
		PlantHeight:        0.4,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"head_weight": "3-4kg",
			"head_color":  "зеленый",
			"head_shape":  "округлый",
			"use":         "квашение",
		},
		Description: "Позднеспелый сорт для квашения",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 170},
			{Code: "BBCH-19", GDDRequired: 450},
			{Code: "BBCH-51", GDDRequired: 800},
			{Code: "BBCH-61", GDDRequired: 1000, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 1200, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 1550},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.6,
				PlantSpacing:    0.5,
				SowingDepth:     1.5,
				GerminationRate: 80,
				SafetyFactor:    1.25,
			},
		},
	}

	// ========== 6. АГРЕССОР F1 ==========
	agressor := Variety{
		ID:                 "agressor_f1",
		Name:               "Агрессор F1",
		SpeciesKey:         "cabbage",
		SpeciesName:        "Капуста",
		BaseTemperature:    5.0,
		MaxTemperature:     27.0,
		DaysToMaturity:     105,
		YieldPotential:     10.0,
		PlantHeight:        0.4,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"head_weight": "3-4kg",
			"head_color":  "зеленый",
			"head_shape":  "округлый",
			"use":         "универсальный",
		},
		Description: "Жаростойкий гибрид для южных регионов",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 145},
			{Code: "BBCH-19", GDDRequired: 390},
			{Code: "BBCH-51", GDDRequired: 680},
			{Code: "BBCH-61", GDDRequired: 880, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 1080, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 1380},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.65,
				PlantSpacing:    0.5,
				SowingDepth:     1.5,
				GerminationRate: 84,
				SafetyFactor:    1.19,
			},
		},
	}

	// ========== 7. АТРИЯ F1 ==========
	atria := Variety{
		ID:                 "atria_f1",
		Name:               "Атрия F1",
		SpeciesKey:         "cabbage",
		SpeciesName:        "Капуста",
		BaseTemperature:    5.0,
		MaxTemperature:     25.0,
		DaysToMaturity:     108,
		YieldPotential:     11.5,
		PlantHeight:        0.45,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"head_weight": "3-5kg",
			"head_color":  "темно-зеленый",
			"head_shape":  "округлый",
			"use":         "свежий",
		},
		Description: "Гибрид для промышленного выращивания",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 148},
			{Code: "BBCH-19", GDDRequired: 395},
			{Code: "BBCH-51", GDDRequired: 690},
			{Code: "BBCH-61", GDDRequired: 890, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 1090, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 1420},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.65,
				PlantSpacing:    0.5,
				SowingDepth:     1.5,
				GerminationRate: 85,
				SafetyFactor:    1.18,
			},
		},
	}

	// ========== 8. БЕЛОСНЕЖКА ==========
	belosnezhka := Variety{
		ID:                 "belosnezhka",
		Name:               "Белоснежка",
		SpeciesKey:         "cabbage",
		SpeciesName:        "Капуста",
		BaseTemperature:    5.0,
		MaxTemperature:     25.0,
		DaysToMaturity:     100,
		YieldPotential:     6.5,
		PlantHeight:        0.35,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"head_weight": "1-2kg",
			"head_color":  "белый",
			"head_shape":  "округлый",
			"use":         "салатный",
		},
		Description: "Раннеспелая капуста для салатов",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 140},
			{Code: "BBCH-19", GDDRequired: 370},
			{Code: "BBCH-51", GDDRequired: 650},
			{Code: "BBCH-61", GDDRequired: 850, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 1050, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 1350},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.5,
				PlantSpacing:    0.4,
				SowingDepth:     1.2,
				GerminationRate: 83,
				SafetyFactor:    1.2,
			},
		},
	}

	// ========== 9. СТОРИДОР F1 ==========
	storidor := Variety{
		ID:                 "storidor_f1",
		Name:               "Сторидор F1",
		SpeciesKey:         "cabbage",
		SpeciesName:        "Капуста",
		BaseTemperature:    5.0,
		MaxTemperature:     25.0,
		DaysToMaturity:     112,
		YieldPotential:     9.5,
		PlantHeight:        0.42,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"head_weight": "3-4kg",
			"head_color":  "зеленый",
			"head_shape":  "округлый",
			"use":         "хранение",
		},
		Description: "Гибрид для длительного хранения",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 152},
			{Code: "BBCH-19", GDDRequired: 410},
			{Code: "BBCH-51", GDDRequired: 720},
			{Code: "BBCH-61", GDDRequired: 920, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 1120, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 1440},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.6,
				PlantSpacing:    0.5,
				SowingDepth:     1.5,
				GerminationRate: 82,
				SafetyFactor:    1.22,
			},
		},
	}

	// ========== 10. РИНДА F1 ==========
	rinda := Variety{
		ID:                 "rinda_f1",
		Name:               "Ринда F1",
		SpeciesKey:         "cabbage",
		SpeciesName:        "Капуста",
		BaseTemperature:    5.0,
		MaxTemperature:     25.0,
		DaysToMaturity:     100,
		YieldPotential:     8.0,
		PlantHeight:        0.4,
		RecommendedSeasons: []string{"spring", "summer", "autumn"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"head_weight": "2-3kg",
			"head_color":  "зеленый",
			"head_shape":  "округлый",
			"use":         "универсальный",
		},
		Description: "Популярный гибрид для свежего потребления и переработки",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 142},
			{Code: "BBCH-19", GDDRequired: 380},
			{Code: "BBCH-51", GDDRequired: 660},
			{Code: "BBCH-61", GDDRequired: 860, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 1060, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 1360},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.6,
				PlantSpacing:    0.45,
				SowingDepth:     1.5,
				GerminationRate: 84,
				SafetyFactor:    1.19,
			},
		},
	}

	// ========== 11. ДЖИНТАМА F1 ==========
	jintama := Variety{
		ID:                 "jintama_f1",
		Name:               "Джинтама F1",
		SpeciesKey:         "cabbage",
		SpeciesName:        "Капуста",
		BaseTemperature:    5.0,
		MaxTemperature:     25.0,
		DaysToMaturity:     118,
		YieldPotential:     10.5,
		PlantHeight:        0.48,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"head_weight": "4-5kg",
			"head_color":  "темно-зеленый",
			"head_shape":  "округлый",
			"use":         "свежий, хранение",
		},
		Description: "Крупноплодный гибрид японской селекции",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 158},
			{Code: "BBCH-19", GDDRequired: 425},
			{Code: "BBCH-51", GDDRequired: 740},
			{Code: "BBCH-61", GDDRequired: 940, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 1140, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 1480},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.7,
				PlantSpacing:    0.55,
				SowingDepth:     1.5,
				GerminationRate: 83,
				SafetyFactor:    1.2,
			},
		},
	}

	// Регистрируем все сорта капусты
	RegisterVariety("cabbage", slava1305)
	RegisterVariety("cabbage", megaton)
	RegisterVariety("cabbage", podarok)
	RegisterVariety("cabbage", menza)
	RegisterVariety("cabbage", krautman)
	RegisterVariety("cabbage", agressor)
	RegisterVariety("cabbage", atria)
	RegisterVariety("cabbage", belosnezhka)
	RegisterVariety("cabbage", storidor)
	RegisterVariety("cabbage", rinda)
	RegisterVariety("cabbage", jintama)
}
