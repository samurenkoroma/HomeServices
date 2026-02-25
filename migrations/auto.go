package main

import (
	"os"
	"samurenkoroma/services/internal/infrastructure/db_table"
	"samurenkoroma/services/internal_old/link"
	"samurenkoroma/services/internal_old/stat"
	"samurenkoroma/services/services/account"
	"samurenkoroma/services/services/homelib"

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
		&account.User{},
		&stat.Stat{},
		&homelib.Book{},
		&homelib.Resource{},
		&homelib.Author{},

		&db_table.Supplier{},
		&db_table.Order{},
		&db_table.OrderItem{},
		&db_table.Product{},
		&db_table.ProductVariant{},
		&db_table.TaxonomyNode{},
	)
}
