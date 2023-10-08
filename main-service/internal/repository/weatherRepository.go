package repository

import (
	"assignment3-012/main-service/internal/database"
	"assignment3-012/main-service/internal/models"
)

func UpdateWeather(weather *models.Weather) error {
	db := database.GetDB()
	
	return db.Model(&weather).Where("id = ?", 1).Updates(&weather).Error
}

func CreateWeather(weather *models.Weather) error {
	db := database.GetDB()
	
	return db.Create(&weather).Error
}

func IsWeatherDataExist() bool {
	db := database.GetDB()
	
	var total int64
	
	db.Model(&models.Weather{}).Count(&total)

	return total != 0
}