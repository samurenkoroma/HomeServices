package catalog

func init() {
	// ========== РЕГИСТРИРУЕМ ВИД ==========
	RegisterSpecies(Species{
		Key:         "cucumber",
		Name:        "Огурец",
		Family:      "cucurbits",
		Description: "Овощная культура семейства тыквенных, требовательна к теплу и влаге",
	})

	// ========== СОРТ "ИЗЯЩНЫЙ" (партенокарпический) ==========
	elegant := Variety{
		ID:                 "elegant",
		Name:               "Изящный",
		SpeciesKey:         "cucumber",
		SpeciesName:        "Огурец",
		DaysToMaturity:     45,
		YieldPotential:     12.0,
		PlantHeight:        2.0,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"greenhouse"},
		Characteristics: map[string]string{
			"fruit_length": "10-12cm",
			"fruit_color":  "темно-зеленый",
			"type":         "партенокарпический",
			"use":          "салатный",
		},
		Description: "Раннеспелый партенокарпический гибрид для защищенного грунта",

		// GDD шкала для огурца (базовая температура 12°C)
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", Name: "Всходы", GDDRequired: 80, Description: "Появление всходов", IsCritical: true},
			{Code: "BBCH-19", Name: "3-4 настоящих листа", GDDRequired: 180, Description: "Активный рост", IsCritical: false},
			{Code: "BBCH-51", Name: "Бутонизация", GDDRequired: 280, Description: "Появление женских цветков", IsCritical: false},
			{Code: "BBCH-61", Name: "Цветение", GDDRequired: 350, Description: "Массовое цветение", IsCritical: true},
			{Code: "BBCH-71", Name: "Плодоношение", GDDRequired: 450, Description: "Образование зеленцов", IsCritical: true},
			{Code: "BBCH-89", Name: "Техническая спелость", GDDRequired: 600, Description: "Зеленцы готовы к сбору", IsCritical: false},
		},

		SeedingRates: map[string]SeedingRate{
			"greenhouse": {
				GrowingType:     "greenhouse",
				RowSpacing:      1.0,
				PlantSpacing:    0.4,
				SowingDepth:     1.5,
				GerminationRate: 90,
				SafetyFactor:    1.1,
			},
		},
	}

	// ========== СОРТ "РОДНИЧОК" (пчелоопыляемый) ==========
	rodnichok := Variety{
		ID:                 "rodnichok",
		Name:               "Родничок",
		SpeciesKey:         "cucumber",
		SpeciesName:        "Огурец",
		DaysToMaturity:     50,
		YieldPotential:     10.0,
		PlantHeight:        1.8,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"fruit_length": "8-10cm",
			"fruit_color":  "зеленый с полосками",
			"type":         "пчелоопыляемый",
			"use":          "засолочный",
		},
		Description: "Пчелоопыляемый сорт для открытого грунта, устойчив к болезням",

		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 90},
			{Code: "BBCH-19", GDDRequired: 200},
			{Code: "BBCH-51", GDDRequired: 300},
			{Code: "BBCH-61", GDDRequired: 380, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 480, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 650},
		},

		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.7,
				PlantSpacing:    0.3,
				SowingDepth:     2.0,
				GerminationRate: 85,
				SafetyFactor:    1.15,
			},
		},
	}

	// ========== СОРТ "ЗОЗУЛЯ" (ранний) ==========
	zozulya := Variety{
		ID:                 "zozulya",
		Name:               "Зозуля",
		SpeciesKey:         "cucumber",
		SpeciesName:        "Огурец",
		DaysToMaturity:     42,
		YieldPotential:     15.0,
		PlantHeight:        2.2,
		RecommendedSeasons: []string{"spring"},
		GrowingTypes:       []string{"greenhouse"},
		Characteristics: map[string]string{
			"fruit_length": "14-16cm",
			"fruit_color":  "ярко-зеленый",
			"type":         "партенокарпический",
			"use":          "салатный",
		},
		Description: "Суперранний гибрид для весенних теплиц",

		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 70},
			{Code: "BBCH-19", GDDRequired: 160},
			{Code: "BBCH-51", GDDRequired: 250},
			{Code: "BBCH-61", GDDRequired: 320, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 400, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 550},
		},

		SeedingRates: map[string]SeedingRate{
			"greenhouse": {
				GrowingType:     "greenhouse",
				RowSpacing:      0.9,
				PlantSpacing:    0.35,
				SowingDepth:     1.5,
				GerminationRate: 92,
				SafetyFactor:    1.08,
			},
		},
	}

	// Регистрируем сорта
	RegisterVariety("cucumber", elegant)
	RegisterVariety("cucumber", rodnichok)
	RegisterVariety("cucumber", zozulya)
}
