package catalog

func init() {
	// ========== РЕГИСТРИРУЕМ ВИД ==========
	RegisterSpecies(Species{
		Key:         "tomato",
		Name:        "Томат",
		Family:      "nightshade",
		Category:    "Овощные",
		ImageUrl:    "https://images.pexels.com/photos/9889794/pexels-photo-9889794.jpeg",
		Description: "Овощная культура семейства пасленовых",
	})

	// ========== 1. ИНКАС F1 ==========
	// Гибрид, раннеспелый, для открытого грунта и пленочных теплиц
	incasF1 := Variety{
		ID:                 "incas_f1",
		Name:               "Инкас F1",
		SpeciesKey:         "tomato",
		SpeciesName:        "Томат",
		BaseTemperature:    10.0,
		MaxTemperature:     30.0,
		DaysToMaturity:     95, // ранний
		YieldPotential:     12.0,
		PlantHeight:        1.2,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground", "greenhouse"},
		Characteristics: map[string]string{
			"fruit_weight": "180-200g",
			"fruit_color":  "красный",
			"fruit_shape":  "округлый",
			"type":         "индетерминантный",
			"use":          "салатный",
			"resistance":   "VTM, F1, Cladosporium",
		},
		Description: "Раннеспелый гибрид для пленочных теплиц и открытого грунта",
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
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.7,
				PlantSpacing:    0.4,
				SowingDepth:     1.5,
				GerminationRate: 90,
				SafetyFactor:    1.15,
			},
			"greenhouse": {
				GrowingType:     "greenhouse",
				RowSpacing:      0.8,
				PlantSpacing:    0.5,
				SowingDepth:     1.0,
				GerminationRate: 92,
				SafetyFactor:    1.1,
			},
		},
	}

	// ========== 2. РИО-ГРАНДЕ ==========
	// Среднеспелый, засухоустойчивый, для открытого грунта
	rioGrande := Variety{
		ID:                 "rio_grande",
		Name:               "Рио-Гранде",
		SpeciesKey:         "tomato",
		SpeciesName:        "Томат",
		BaseTemperature:    10.0,
		MaxTemperature:     32.0, // более устойчив к жаре
		DaysToMaturity:     110,  // среднеспелый
		YieldPotential:     10.0,
		PlantHeight:        0.8,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"fruit_weight": "100-120g",
			"fruit_color":  "красный",
			"fruit_shape":  "сливовидный",
			"type":         "детерминантный",
			"use":          "для переработки",
			"resistance":   "вертициллез, фузариоз",
		},
		Description: "Засухоустойчивый сорт для открытого грунта и фермерских хозяйств",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 110},
			{Code: "BBCH-19", GDDRequired: 320},
			{Code: "BBCH-51", GDDRequired: 550},
			{Code: "BBCH-61", GDDRequired: 730, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 880, IsCritical: true},
			{Code: "BBCH-81", GDDRequired: 1100},
			{Code: "BBCH-89", GDDRequired: 1300},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.7,
				PlantSpacing:    0.35,
				SowingDepth:     2.0,
				GerminationRate: 85,
				SafetyFactor:    1.2,
			},
		},
	}

	// ========== 3. БЕЛЛА РОСА F1 ==========
	// Ранний гибрид, холодостойкий
	bellaRosa := Variety{
		ID:                 "bella_rosa_f1",
		Name:               "Белла Роса F1",
		SpeciesKey:         "tomato",
		SpeciesName:        "Томат",
		BaseTemperature:    8.0, // холодостойкий
		MaxTemperature:     28.0,
		DaysToMaturity:     85, // ранний
		YieldPotential:     9.0,
		PlantHeight:        1.0,
		RecommendedSeasons: []string{"spring"},
		GrowingTypes:       []string{"open_ground", "greenhouse"},
		Characteristics: map[string]string{
			"fruit_weight":   "80-100g",
			"fruit_color":    "розовый",
			"fruit_shape":    "округлый",
			"type":           "детерминантный",
			"use":            "салатный",
			"cold_resistant": "да",
		},
		Description: "Холодостойкий гибрид для раннего выращивания в открытом грунте",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 90},
			{Code: "BBCH-19", GDDRequired: 250},
			{Code: "BBCH-51", GDDRequired: 430},
			{Code: "BBCH-61", GDDRequired: 580, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 700, IsCritical: true},
			{Code: "BBCH-81", GDDRequired: 880},
			{Code: "BBCH-89", GDDRequired: 1050},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.6,
				PlantSpacing:    0.4,
				SowingDepth:     1.5,
				GerminationRate: 88,
				SafetyFactor:    1.15,
			},
		},
	}

	// ========== 4. СОЛЕРОССО F1 ==========
	// Солеустойчивый гибрид для засоленных почв
	solerosso := Variety{
		ID:                 "solerossa_f1",
		Name:               "Солероссо F1",
		SpeciesKey:         "tomato",
		SpeciesName:        "Томат",
		BaseTemperature:    10.0,
		MaxTemperature:     32.0,
		DaysToMaturity:     100,
		YieldPotential:     11.0,
		PlantHeight:        1.1,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"fruit_weight":   "150-170g",
			"fruit_color":    "красный",
			"fruit_shape":    "округлый",
			"type":           "индетерминантный",
			"use":            "универсальный",
			"salt_resistant": "да",
		},
		Description: "Солеустойчивый гибрид для выращивания на засоленных почвах",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 105},
			{Code: "BBCH-19", GDDRequired: 300},
			{Code: "BBCH-51", GDDRequired: 520},
			{Code: "BBCH-61", GDDRequired: 690, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 830, IsCritical: true},
			{Code: "BBCH-81", GDDRequired: 1050},
			{Code: "BBCH-89", GDDRequired: 1250},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.7,
				PlantSpacing:    0.45,
				SowingDepth:     1.8,
				GerminationRate: 87,
				SafetyFactor:    1.18,
			},
		},
	}

	// ========== 5. МАКАН F1 ==========
	// Крупноплодный гибрид для теплиц
	macan := Variety{
		ID:                 "macan_f1",
		Name:               "Макан F1",
		SpeciesKey:         "tomato",
		SpeciesName:        "Томат",
		BaseTemperature:    10.0,
		MaxTemperature:     30.0,
		DaysToMaturity:     115,
		YieldPotential:     14.0,
		PlantHeight:        1.8,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"greenhouse"},
		Characteristics: map[string]string{
			"fruit_weight": "250-300g",
			"fruit_color":  "красный",
			"fruit_shape":  "плоскоокруглый",
			"type":         "индетерминантный",
			"use":          "салатный",
		},
		Description: "Крупноплодный гибрид для защищенного грунта",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 115},
			{Code: "BBCH-19", GDDRequired: 330},
			{Code: "BBCH-51", GDDRequired: 570},
			{Code: "BBCH-61", GDDRequired: 760, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 910, IsCritical: true},
			{Code: "BBCH-81", GDDRequired: 1150},
			{Code: "BBCH-89", GDDRequired: 1380},
		},
		SeedingRates: map[string]SeedingRate{
			"greenhouse": {
				GrowingType:     "greenhouse",
				RowSpacing:      0.9,
				PlantSpacing:    0.6,
				SowingDepth:     1.2,
				GerminationRate: 91,
				SafetyFactor:    1.1,
			},
		},
	}

	// ========== 6. ЯМАМОТО ==========
	// Японский сорт, для теплиц
	yamamoto := Variety{
		ID:                 "yamamoto",
		Name:               "Ямамото",
		SpeciesKey:         "tomato",
		SpeciesName:        "Томат",
		BaseTemperature:    10.0,
		MaxTemperature:     28.0,
		DaysToMaturity:     108,
		YieldPotential:     13.0,
		PlantHeight:        1.6,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"greenhouse"},
		Characteristics: map[string]string{
			"fruit_weight": "200-250g",
			"fruit_color":  "темно-красный",
			"fruit_shape":  "округлый",
			"type":         "индетерминантный",
			"use":          "салатный",
			"origin":       "Япония",
		},
		Description: "Японский сорт с высокими вкусовыми качествами",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 108},
			{Code: "BBCH-19", GDDRequired: 310},
			{Code: "BBCH-51", GDDRequired: 540},
			{Code: "BBCH-61", GDDRequired: 720, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 860, IsCritical: true},
			{Code: "BBCH-81", GDDRequired: 1080},
			{Code: "BBCH-89", GDDRequired: 1300},
		},
		SeedingRates: map[string]SeedingRate{
			"greenhouse": {
				GrowingType:     "greenhouse",
				RowSpacing:      0.8,
				PlantSpacing:    0.55,
				SowingDepth:     1.2,
				GerminationRate: 89,
				SafetyFactor:    1.12,
			},
		},
	}

	// ========== 7. БЫЧЬЕ СЕРДЦЕ (уже есть, обновим) ==========
	bullHeart := Variety{
		ID:                 "bull_heart",
		Name:               "Бычье сердце",
		SpeciesKey:         "tomato",
		SpeciesName:        "Томат",
		BaseTemperature:    10.0,
		MaxTemperature:     30.0,
		DaysToMaturity:     120,
		YieldPotential:     8.5,
		PlantHeight:        1.5,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground", "greenhouse"},
		Characteristics: map[string]string{
			"fruit_weight": "300-500g",
			"fruit_color":  "малиновый",
			"fruit_shape":  "сердцевидный",
			"type":         "индетерминантный",
			"use":          "салатный",
		},
		Description: "Крупноплодный салатный сорт",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 120},
			{Code: "BBCH-19", GDDRequired: 350},
			{Code: "BBCH-51", GDDRequired: 600},
			{Code: "BBCH-61", GDDRequired: 800, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 950, IsCritical: true},
			{Code: "BBCH-81", GDDRequired: 1200},
			{Code: "BBCH-89", GDDRequired: 1400},
		},
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

	// ========== 8. ХУРМА ==========
	persimmon := Variety{
		ID:                 "persimmon",
		Name:               "Хурма",
		SpeciesKey:         "tomato",
		SpeciesName:        "Томат",
		BaseTemperature:    10.0,
		MaxTemperature:     30.0,
		DaysToMaturity:     115,
		YieldPotential:     7.0,
		PlantHeight:        1.2,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground", "greenhouse"},
		Characteristics: map[string]string{
			"fruit_weight": "150-200g",
			"fruit_color":  "оранжевый",
			"fruit_shape":  "плоскоокруглый",
			"type":         "детерминантный",
			"use":          "салатный",
		},
		Description: "Сорт с оранжевыми плодами, напоминающими хурму",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 112},
			{Code: "BBCH-19", GDDRequired: 325},
			{Code: "BBCH-51", GDDRequired: 560},
			{Code: "BBCH-61", GDDRequired: 745, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 890, IsCritical: true},
			{Code: "BBCH-81", GDDRequired: 1120},
			{Code: "BBCH-89", GDDRequired: 1330},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.7,
				PlantSpacing:    0.5,
				SowingDepth:     2.0,
				GerminationRate: 84,
				SafetyFactor:    1.22,
			},
		},
	}

	// ========== 9. ЧЕРНАЯ ГРОЗДЬ ==========
	blackBunch := Variety{
		ID:                 "black_bunch",
		Name:               "Черная гроздь",
		SpeciesKey:         "tomato",
		SpeciesName:        "Томат",
		BaseTemperature:    10.0,
		MaxTemperature:     30.0,
		DaysToMaturity:     95,
		YieldPotential:     6.0,
		PlantHeight:        1.5,
		RecommendedSeasons: []string{"spring", "summer", "autumn"},
		GrowingTypes:       []string{"greenhouse"},
		Characteristics: map[string]string{
			"fruit_weight": "30-40g",
			"fruit_color":  "черно-коричневый",
			"fruit_shape":  "сливовидный",
			"type":         "индетерминантный",
			"use":          "черри",
		},
		Description: "Черри с необычным цветом и вкусом",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 95},
			{Code: "BBCH-19", GDDRequired: 270},
			{Code: "BBCH-51", GDDRequired: 460},
			{Code: "BBCH-61", GDDRequired: 620, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 750, IsCritical: true},
			{Code: "BBCH-81", GDDRequired: 950},
			{Code: "BBCH-89", GDDRequired: 1120},
		},
		SeedingRates: map[string]SeedingRate{
			"greenhouse": {
				GrowingType:     "greenhouse",
				RowSpacing:      0.7,
				PlantSpacing:    0.4,
				SowingDepth:     1.2,
				GerminationRate: 88,
				SafetyFactor:    1.13,
			},
		},
	}

	// Регистрируем все сорта томата
	RegisterVariety("tomato", incasF1)
	RegisterVariety("tomato", rioGrande)
	RegisterVariety("tomato", bellaRosa)
	RegisterVariety("tomato", solerosso)
	RegisterVariety("tomato", macan)
	RegisterVariety("tomato", yamamoto)
	RegisterVariety("tomato", bullHeart)
	RegisterVariety("tomato", persimmon)
	RegisterVariety("tomato", blackBunch)
}
