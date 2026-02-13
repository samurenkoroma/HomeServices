package main

import (
	"os"
	"samurenkoroma/services/internal/domain"
	"samurenkoroma/services/internal/infrastructure/payload"
	"samurenkoroma/services/internal_old/link"
	"samurenkoroma/services/internal_old/stat"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&link.Link{},
		&domain.User{},
		&stat.Stat{},
		&domain.Book{},
		&domain.Resource{},
		&domain.Author{},
		&payload.SupplierGorm{},
	)
}
