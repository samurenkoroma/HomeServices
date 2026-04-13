package catalog

func init() {
	// ========== РЕГИСТРИРУЕМ ВИД ==========
	RegisterSpecies(Species{
		Key:         "cucumber",
		Name:        "Огурец",
		Family:      "cucurbits",
		Category:    "Овощные",
		ImageUrl:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRKzl5NoyBOWI5L_YwWBK1aBHm6MZfXwUsnDw&s",
		Description: "Овощная культура семейства тыквенных, требовательна к теплу и влаге",
	})

	// ========== 1. РОДНИЧОК F1 ==========
	rodnichok := Variety{
		ID:                 "rodnichok_f1",
		Name:               "Родничок F1",
		SpeciesKey:         "cucumber",
		SpeciesName:        "Огурец",
		BaseTemperature:    12.0,
		MaxTemperature:     35.0,
		DaysToMaturity:     48,
		YieldPotential:     10.0,
		PlantHeight:        1.8,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"fruit_length": "8-10cm",
			"fruit_color":  "зеленый с полосками",
			"fruit_weight": "80-100g",
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

	// ========== 2. БОЧКОВОЙ F1 ==========
	bochkovoi := Variety{
		ID:                 "bochkovoi_f1",
		Name:               "Бочковой F1",
		SpeciesKey:         "cucumber",
		SpeciesName:        "Огурец",
		BaseTemperature:    12.0,
		MaxTemperature:     35.0,
		DaysToMaturity:     50,
		YieldPotential:     12.0,
		PlantHeight:        2.0,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"fruit_length": "10-12cm",
			"fruit_color":  "темно-зеленый",
			"fruit_weight": "100-120g",
			"type":         "партенокарпический",
			"use":          "засолочный",
		},
		Description: "Бочкового типа для засолки и маринования",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 92},
			{Code: "BBCH-19", GDDRequired: 205},
			{Code: "BBCH-51", GDDRequired: 310},
			{Code: "BBCH-61", GDDRequired: 390, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 490, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 660},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.7,
				PlantSpacing:    0.35,
				SowingDepth:     2.0,
				GerminationRate: 86,
				SafetyFactor:    1.14,
			},
		},
	}

	// ========== 3. НЕЖИНСКИЙ ==========
	nezhinsky := Variety{
		ID:                 "nezhinsky",
		Name:               "Нежинский",
		SpeciesKey:         "cucumber",
		SpeciesName:        "Огурец",
		BaseTemperature:    12.0,
		MaxTemperature:     35.0,
		DaysToMaturity:     55,
		YieldPotential:     9.0,
		PlantHeight:        1.5,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"fruit_length": "10-14cm",
			"fruit_color":  "зеленый",
			"fruit_weight": "90-110g",
			"type":         "пчелоопыляемый",
			"use":          "универсальный",
			"origin":       "Украина",
		},
		Description: "Классический сорт для открытого грунта",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 95},
			{Code: "BBCH-19", GDDRequired: 210},
			{Code: "BBCH-51", GDDRequired: 320},
			{Code: "BBCH-61", GDDRequired: 400, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 500, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 680},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.7,
				PlantSpacing:    0.3,
				SowingDepth:     2.5,
				GerminationRate: 84,
				SafetyFactor:    1.16,
			},
		},
	}

	// ========== 4. КРИСПИНА F1 ==========
	crispina := Variety{
		ID:                 "crispina_f1",
		Name:               "Криспина F1",
		SpeciesKey:         "cucumber",
		SpeciesName:        "Огурец",
		BaseTemperature:    12.0,
		MaxTemperature:     35.0,
		DaysToMaturity:     42,
		YieldPotential:     14.0,
		PlantHeight:        2.2,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"greenhouse"},
		Characteristics: map[string]string{
			"fruit_length": "12-14cm",
			"fruit_color":  "темно-зеленый",
			"fruit_weight": "80-100g",
			"type":         "партенокарпический",
			"use":          "салатный",
		},
		Description: "Ранний гибрид для теплиц",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 85},
			{Code: "BBCH-19", GDDRequired: 190},
			{Code: "BBCH-51", GDDRequired: 280},
			{Code: "BBCH-61", GDDRequired: 360, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 450, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 600},
		},
		SeedingRates: map[string]SeedingRate{
			"greenhouse": {
				GrowingType:     "greenhouse",
				RowSpacing:      0.9,
				PlantSpacing:    0.4,
				SowingDepth:     1.5,
				GerminationRate: 92,
				SafetyFactor:    1.08,
			},
		},
	}

	// ========== 5. ЭКОЛЬ F1 ==========
	ecole := Variety{
		ID:                 "ecole_f1",
		Name:               "Эколь F1",
		SpeciesKey:         "cucumber",
		SpeciesName:        "Огурец",
		BaseTemperature:    12.0,
		MaxTemperature:     35.0,
		DaysToMaturity:     45,
		YieldPotential:     11.0,
		PlantHeight:        1.9,
		RecommendedSeasons: []string{"spring", "summer", "autumn"},
		GrowingTypes:       []string{"greenhouse"},
		Characteristics: map[string]string{
			"fruit_length": "14-16cm",
			"fruit_color":  "зеленый",
			"fruit_weight": "100-120g",
			"type":         "партенокарпический",
			"use":          "салатный",
		},
		Description: "Урожайный гибрид для продленного оборота",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 88},
			{Code: "BBCH-19", GDDRequired: 195},
			{Code: "BBCH-51", GDDRequired: 290},
			{Code: "BBCH-61", GDDRequired: 370, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 465, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 630},
		},
		SeedingRates: map[string]SeedingRate{
			"greenhouse": {
				GrowingType:     "greenhouse",
				RowSpacing:      0.8,
				PlantSpacing:    0.45,
				SowingDepth:     1.5,
				GerminationRate: 91,
				SafetyFactor:    1.09,
			},
		},
	}

	// ========== 6. МЕРТУС F1 ==========
	mertus := Variety{
		ID:                 "mertus_f1",
		Name:               "Мертус F1",
		SpeciesKey:         "cucumber",
		SpeciesName:        "Огурец",
		BaseTemperature:    12.0,
		MaxTemperature:     35.0,
		DaysToMaturity:     47,
		YieldPotential:     13.0,
		PlantHeight:        2.0,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"greenhouse"},
		Characteristics: map[string]string{
			"fruit_length": "10-12cm",
			"fruit_color":  "ярко-зеленый",
			"fruit_weight": "80-90g",
			"type":         "партенокарпический",
			"use":          "универсальный",
		},
		Description: "Гладкоплодный гибрид для теплиц",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 90},
			{Code: "BBCH-19", GDDRequired: 200},
			{Code: "BBCH-51", GDDRequired: 300},
			{Code: "BBCH-61", GDDRequired: 380, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 475, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 645},
		},
		SeedingRates: map[string]SeedingRate{
			"greenhouse": {
				GrowingType:     "greenhouse",
				RowSpacing:      0.85,
				PlantSpacing:    0.4,
				SowingDepth:     1.5,
				GerminationRate: 90,
				SafetyFactor:    1.1,
			},
		},
	}

	// ========== 7. МАДРИЛЕНЕ F1 ==========
	madrilene := Variety{
		ID:                 "madrilene_f1",
		Name:               "Мадрилене F1",
		SpeciesKey:         "cucumber",
		SpeciesName:        "Огурец",
		BaseTemperature:    12.0,
		MaxTemperature:     35.0,
		DaysToMaturity:     44,
		YieldPotential:     12.5,
		PlantHeight:        2.1,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"greenhouse"},
		Characteristics: map[string]string{
			"fruit_length": "16-18cm",
			"fruit_color":  "темно-зеленый",
			"fruit_weight": "120-140g",
			"type":         "партенокарпический",
			"use":          "салатный",
		},
		Description: "Длинноплодный гибрид для салатов",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 87},
			{Code: "BBCH-19", GDDRequired: 193},
			{Code: "BBCH-51", GDDRequired: 285},
			{Code: "BBCH-61", GDDRequired: 365, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 460, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 620},
		},
		SeedingRates: map[string]SeedingRate{
			"greenhouse": {
				GrowingType:     "greenhouse",
				RowSpacing:      0.9,
				PlantSpacing:    0.5,
				SowingDepth:     1.5,
				GerminationRate: 89,
				SafetyFactor:    1.11,
			},
		},
	}

	// ========== 8. ГЕКТОР F1 ==========
	hector := Variety{
		ID:                 "hector_f1",
		Name:               "Гектор F1",
		SpeciesKey:         "cucumber",
		SpeciesName:        "Огурец",
		BaseTemperature:    12.0,
		MaxTemperature:     35.0,
		DaysToMaturity:     40,
		YieldPotential:     8.0,
		PlantHeight:        1.2,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"fruit_length": "6-8cm",
			"fruit_color":  "зеленый",
			"fruit_weight": "60-80g",
			"type":         "партенокарпический",
			"use":          "засолочный",
		},
		Description: "Ультраранний гибрид для открытого грунта",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 80},
			{Code: "BBCH-19", GDDRequired: 180},
			{Code: "BBCH-51", GDDRequired: 270},
			{Code: "BBCH-61", GDDRequired: 340, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 430, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 580},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.6,
				PlantSpacing:    0.25,
				SowingDepth:     2.0,
				GerminationRate: 87,
				SafetyFactor:    1.17,
			},
		},
	}

	// ========== 9. АЯКС F1 ==========
	ajax := Variety{
		ID:                 "ajax_f1",
		Name:               "Аякс F1",
		SpeciesKey:         "cucumber",
		SpeciesName:        "Огурец",
		BaseTemperature:    12.0,
		MaxTemperature:     35.0,
		DaysToMaturity:     46,
		YieldPotential:     10.5,
		PlantHeight:        1.7,
		RecommendedSeasons: []string{"spring", "summer"},
		GrowingTypes:       []string{"open_ground"},
		Characteristics: map[string]string{
			"fruit_length": "10-12cm",
			"fruit_color":  "зеленый",
			"fruit_weight": "90-110g",
			"type":         "пчелоопыляемый",
			"use":          "универсальный",
		},
		Description: "Устойчивый к стрессам гибрид для открытого грунта",
		PhenophaseGDD: []PhenophaseGDD{
			{Code: "BBCH-10", GDDRequired: 91},
			{Code: "BBCH-19", GDDRequired: 202},
			{Code: "BBCH-51", GDDRequired: 305},
			{Code: "BBCH-61", GDDRequired: 385, IsCritical: true},
			{Code: "BBCH-71", GDDRequired: 485, IsCritical: true},
			{Code: "BBCH-89", GDDRequired: 655},
		},
		SeedingRates: map[string]SeedingRate{
			"open_ground": {
				GrowingType:     "open_ground",
				RowSpacing:      0.7,
				PlantSpacing:    0.35,
				SowingDepth:     2.0,
				GerminationRate: 85,
				SafetyFactor:    1.15,
			},
		},
	}

	// Регистрируем все сорта огурца
	RegisterVariety("cucumber", rodnichok)
	RegisterVariety("cucumber", bochkovoi)
	RegisterVariety("cucumber", nezhinsky)
	RegisterVariety("cucumber", crispina)
	RegisterVariety("cucumber", ecole)
	RegisterVariety("cucumber", mertus)
	RegisterVariety("cucumber", madrilene)
	RegisterVariety("cucumber", hector)
	RegisterVariety("cucumber", ajax)
}
