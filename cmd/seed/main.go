package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/infrastructure/configs"
	inmemory "samurenkoroma/services/internal/infrastructure/messaging/rabbitmq"
	"samurenkoroma/services/internal/modules/growing/application/command/cultivation"
	"samurenkoroma/services/pkg/db"

	_ "github.com/lib/pq"
)

var (
	dataDir = flag.String("data", "./data", "Path to seed data directory")
	dryRun  = flag.Bool("dry-run", false, "Dry run mode - only validate, don't insert")
	module  = flag.String("module", "all", "Module to seed (crop, farm, growing, all)")
)

type seedData struct {
	CropTypes []struct {
		Name        string `json:"name"`
		Category    string `json:"category"`
		Description string `json:"description"`
		IsPerennial bool   `json:"is_perennial"`
		Family      string `json:"family"`
		Icon        string `json:"lucideIcon"`
		ImageUrl    string `json:"imageUrl"`
	} `json:"crop_types"`
	Varieties []struct {
		Name               string   `json:"name"`
		Croptype           string   `json:"crop"`
		Description        string   `json:"description"`
		VegetationDays     string   `json:"vegetation_days"`
		YieldPotential     string   `json:"yield_potential"`
		RecommendedRegions []string `json:"recommended_regions"`
	}
	CultivationPlan []struct {
		Version   int     `json:"version"`
		Name      string  `json:"name"`
		CropKey   string  `json:"crop_key"`
		VarietyId *string `json:"varietyId"`
		Steps     []struct {
			Type    string `json:"type" validate:"required"`
			Title   string `json:"title" validate:"required"`
			Trigger struct {
				Type  string         `json:"type" validate:"required"`
				Value map[string]any `json:"value" validate:"required,min=1,dive,keys,min=3,endkeys,required"`
			} `json:"trigger" validate:"required"`
		} `json:"steps" validate:"required"`
	} `json:"cultivation_plans"`
}

func main() {
	flag.Parse()

	// Загружаем конфигурацию
	cfg := configs.LoadConfig()

	// Подключаемся к БД
	db, err := db.NewDB(context.Background(), cfg.Db.Dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	bus := inmemory.NewInMemoryEventBus()

	uowFactory := repository.NewUnitOfWorkFactory(db, bus)
	data := readFile()

	var parseData seedData

	if err := json.Unmarshal(data, &parseData); err != nil {
		fmt.Errorf("failed to parse crop_types.json: %w", err)
	}
	log.Printf("Found %d crop types to seed", len(parseData.CropTypes))
	switch *module {
	case "all":
		if err := seedCultivationPlans(uowFactory, parseData, *dryRun); err != nil {
			log.Fatalf("Failed to seed cultivation plan: %v", err)
		}
		//if err := seedCropTypes(uowFactory, parseData, *dryRun); err != nil {
		//	log.Fatalf("Failed to seed crop types: %v", err)
		//}
		//if err := seedVarieties(uowFactory, parseData, *dryRun); err != nil {
		//	log.Fatalf("Failed to seed varieties: %v", err)
		//}
	case "crop":
		//if err := seedCropTypes(uowFactory, parseData, *dryRun); err != nil {
		//	log.Fatalf("Failed to seed crop types: %v", err)
		//}
		//if err := seedVarieties(uowFactory, parseData, *dryRun); err != nil {
		//	log.Fatalf("Failed to seed varieties: %v", err)
		//}
	case "farm":
		fmt.Printf("farm seed")
	case "growing":
		if err := seedCultivationPlans(uowFactory, parseData, *dryRun); err != nil {
			log.Fatalf("Failed to seed cultivation plan: %v", err)
		}
	}

	log.Println("Seeding completed successfully!")
}

func seedCultivationPlans(uowFactory repository.Factory, data seedData, dryRun bool) error {
	for _, c := range data.CultivationPlan {
		if dryRun {
			log.Printf("[DRY RUN] Would create: %s (category: %s)", c.Name, c.CropKey)
			continue
		}

		// Создаём обработчик
		handler := cultivation.NewCultivationPlanHandler(uowFactory)
		cmd := cultivation.CreateCultivationPlanCmd{
			Version: c.Version,
			Name:    c.Name,
			CropKey: c.CropKey,
			Steps:   c.Steps,
		}

		// Выполняем команду
		if _, err := handler.Create(context.Background(), &cmd); err != nil {
			return fmt.Errorf("failed to create cultivation plan %s: %w", c.Name, err)
		}

		log.Printf("Created crop type: %s", c.Name)
	}
	return nil
}

//
//func seedVarieties(uowFactory repository.Factory, data seedData, dryRun bool) error {
//	for _, v := range data.Varieties {
//		if dryRun {
//			log.Printf("[DRY RUN] Would create: %s (category: %s)", v.Name, v.Croptype)
//			continue
//		}
//
//		cmd := commands.CreateVarietyCmd{
//			Name:               v.Name,
//			Crop:               v.Croptype,
//			Description:        v.Description,
//			VegetationDays:     v.VegetationDays,
//			YieldPotential:     v.YieldPotential,
//			RecommendedRegions: v.RecommendedRegions,
//		}
//
//		// Создаём обработчик
//		handler := commands.NewCreateVarietyHandler(uowFactory)
//
//		// Выполняем команду
//		if err := handler.Handle(context.Background(), cmd); err != nil {
//			// Если уже существует — пропускаем
//			if err == croptype.ErrVarietyAlreadyExists {
//				log.Printf("Skipping. Variety '%s' already exists", v.Name)
//				continue
//			}
//			return fmt.Errorf("failed to create variety %s: %w", v.Name, err)
//		}
//
//		log.Printf("Created crop type: %s", v.Name)
//	}
//	return nil
//}
//
//// seedCropTypes заливает типы культур из JSON
//func seedCropTypes(uowFactory repository.Factory, data seedData, dryRun bool) error {
//	for _, ct := range data.CropTypes {
//		if dryRun {
//			log.Printf("[DRY RUN] Would create: %s (category: %s, perennial: %v)",
//				ct.Name, ct.Category, ct.IsPerennial)
//			continue
//		}
//
//		re := regexp.MustCompile(`([а-яА-Я]*)\s\((\w+)\)`)
//		family := re.FindStringSubmatch(ct.Family)
//		category := re.FindStringSubmatch(ct.Category)
//		// Создаём команду
//		ruCat := category[1]
//		ruFam := family[1]
//		cmd := commands.CreateCropTypeCmd{
//			Name:        ct.Name,
//			Category:    strings.ToLower(category[2]),
//			CategoryRu:  &ruCat,
//			Family:      strings.ToLower(family[2]),
//			FamilyRu:    &ruFam,
//			Icon:        ct.Icon,
//			ImageURL:    ct.ImageUrl,
//			Description: ct.Description,
//			IsPerennial: ct.IsPerennial,
//		}
//
//		// Создаём обработчик
//		handler := commands.NewCreateCropTypeHandler(uowFactory)
//
//		// Выполняем команду
//		if err := handler.Handle(context.Background(), cmd); err != nil {
//			// Если уже существует — пропускаем
//			if err == croptype.ErrCropTypeAlreadyExists {
//				log.Printf("Skipping. CropType '%s' already exists", ct.Name)
//				continue
//			}
//			return fmt.Errorf("failed to create crop type %s: %w", ct.Name, err)
//		}
//
//		log.Printf("Created crop type: %s", ct.Name)
//	}
//
//	return nil
//}

func readFile() []byte {
	dataPath := filepath.Join(*dataDir, "seeds.json")
	data, err := os.ReadFile(dataPath)
	if err != nil {
		fmt.Errorf("failed to read crop_types.json: %w", err)
		return nil
	}
	return data
}
