package database

import (
	"assignment3-012/main-service/internal/config"
	"assignment3-012/main-service/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	err error
)

func StartDB() {
	config := config.DBConfig()

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error:", err)
	}

	db.Debug().AutoMigrate(models.Weather{})
}

func GetDB() *gorm.DB {
	return db
}