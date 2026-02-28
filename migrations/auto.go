package main

import (
	"os"

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
	//&link.Link{},
	//&account.User{},
	//&stat.Stat{},
	//&homelib.Book{},
	//&homelib.Resource{},
	//&homelib.Author{},
	//
	//&db_table.Supplier{},
	//&db_table.Order{},
	//&db_table.OrderItem{},
	//&db_table.Product{},
	//&db_table.ProductVariant{},
	//&db_table.TaxonomyNode{},
	)
}
