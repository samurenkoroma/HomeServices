package db

import (
	"database/sql"
	"log"
	"os"
	"samurenkoroma/services/internal/infrastructure/configs"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Db struct {
	*gorm.DB
}

func NewDb(config *configs.Config) *Db {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 200 * time.Millisecond,            // Slow SQL threshold
			LogLevel:      logger.LogLevel(config.Db.Logger), // Log level: Info logs everything
			Colorful:      true,                              // Enable colorful output
		},
	)
	db, err := gorm.Open(
		postgres.Open(config.Db.Dsn),
		&gorm.Config{Logger: newLogger})

	if err != nil {
		panic(err)
	}
	return &Db{db}
}

func NewDatabase(dsn string) (*sql.DB, error) {
	return sql.Open("postgres", dsn)
}
