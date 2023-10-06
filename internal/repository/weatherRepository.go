package repository

import (
	"assignment3-012/internal/database"
	"assignment3-012/internal/models"
)

func UpdateWeather(weather *models.Weather) error {
	db := database.GetDB()
	
	return db.Model(&weather).Where("id = ?", 1).Updates(&weather).Error
}

func CreateWeather(weather *models.Weather) error {
	db := database.GetDB()
	
	return db.Create(&weather).Error
}

func GetTotalWeatherData() int64 {
	db := database.GetDB()
	
	var total int64
	
	db.Model(&models.Weather{}).Count(&total)

	return total
}