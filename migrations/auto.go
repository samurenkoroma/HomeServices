package main

import (
	"os"
	"samurenkoroma/services/internal/infrastructure/db_table"
	"samurenkoroma/services/internal_old/link"
	"samurenkoroma/services/internal_old/stat"
	"samurenkoroma/services/services/account"
	"samurenkoroma/services/services/homelib"
	"samurenkoroma/services/services/storehouse"

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
		&db_table.Invoice{},
		&db_table.InvoiceItem{},
		&storehouse.Seed{},
		&storehouse.Vendor{},
		&storehouse.VendorSeeds{},
	)
}
