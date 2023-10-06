package database

import (
	"assignment3-012/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host = "localhost"
	user = "postgres"
	password = "postgres"
	port = "5432"
	dbname = "weather_db"
	sslmode = "disable"
	db *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s", host, user, password, port, dbname, sslmode)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error:", err)
	}

	db.Debug().AutoMigrate(models.Weather{})
}

func GetDB() *gorm.DB {
	return db
}