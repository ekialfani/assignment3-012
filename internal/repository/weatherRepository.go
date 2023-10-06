package repository

import (
	"assignment3-012/internal/database"
	"assignment3-012/internal/models"
)

func UpdateWeather(weather *models.Weather) error {
	db := database.GetDB()
	
	return db.Model(&weather).Where("id = ?", 1).Updates(&weather).Error
}

func GetWeather() (*models.Weather, error) {
	db := database.GetDB()
	
	weather := models.Weather{}
	err := db.First(&weather).Error

	if err != nil {
		return nil, err
	}

	return &weather, nil
}